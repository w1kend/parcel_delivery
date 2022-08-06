package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/w1kend/parcel_delivery/internal/app/parcel_delivery"
	"github.com/w1kend/parcel_delivery/internal/config"
	"github.com/w1kend/parcel_delivery/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery/internal/pkg/health"
	"github.com/w1kend/parcel_delivery/internal/pkg/repositories"
	"github.com/w1kend/parcel_delivery/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	tokenManager := auth.NewTokenManager(config.GetValue(config.JWTSecret).String())
	authInterceptor := auth.NewAuthInterceptor(tokenManager, map[string]interface{}{
		"/parcel_delivery.ParcelDelivery/SignIn": nil,
		"/parcel_delivery.ParcelDelivery/SignUp": nil,
	})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)
	reflection.Register(grpcServer)

	health := health.NewChecker("parcel_delivery")

	db, err := config.NewPostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	health.Track("postgres", config.PgHealthChecker(db))

	ordersRepo := repositories.NewOrdersRepo(db)
	usersRepo := repositories.NewUsersRepository(db)
	hasher := auth.NewHasher()

	impl := parcel_delivery.NewImplementation(ordersRepo, usersRepo, hasher, tokenManager)

	parcel_delivery_grpc.RegisterParcelDeliveryServer(grpcServer, &impl)

	lis, err := net.Listen("tcp", ":"+config.GetValue(config.AppPort).String())
	fmt.Println("listen on port " + config.GetValue(config.AppPort).String())
	if err != nil {
		panic(err)
	}

	healthEndpoind(ctx, health)
	fmt.Println("health endpoint started")

	fmt.Println("app started")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("stopped server with: %s", err.Error())
	}
	cancel()

	os.Exit(1)
}

func healthEndpoind(ctx context.Context, checker *health.Checker) {
	srv := http.Server{
		Addr: ":8080",
	}

	mux := http.NewServeMux()
	srv.Handler = mux

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res, err := json.Marshal(checker.Check())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
	})

	done := make(chan struct{}, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("stopped health server with: %s", err.Error())
			done <- struct{}{}
		}
	}()

	go func() {
		select {
		case <-ctx.Done():
			if err := srv.Shutdown(ctx); err != nil {
				fmt.Printf("stopped server with: %s", err.Error())
			}
		case <-done:
			return
		}
	}()
}
