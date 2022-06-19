package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/w1kend/parcel_delivery_test/internal/app/parcel_delivery"
	"github.com/w1kend/parcel_delivery_test/internal/config"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery_test/internal/pkg/repositories"
	"github.com/w1kend/parcel_delivery_test/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.BuildConfig()

	tokenManager := auth.NewTokenManager(cfg.JWTSecret)
	authInterceptor := auth.NewAuthInterceptor(tokenManager, map[string]interface{}{
		"/parcel_delivery.ParcelDelivery/SignIn": nil,
		"/parcel_delivery.ParcelDelivery/SignUp": nil,
	})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)
	reflection.Register(grpcServer)

	db, err := sql.Open("postgres", getDbDsn(cfg))
	if err != nil {
		panic(err)
	}
	ordersRepo := repositories.NewOrdersRepo(db)
	usersRepo := repositories.NewUsersRepository(db)
	hasher := auth.NewHasher(cfg.HashCost)

	impl := parcel_delivery.NewImplementation(ordersRepo, usersRepo, hasher, tokenManager)

	parcel_delivery_grpc.RegisterParcelDeliveryServer(grpcServer, &impl)

	lis, err := net.Listen("tcp", ":"+cfg.AppPort)
	fmt.Println("listen on port " + cfg.AppPort)
	if err != nil {
		panic(err)
	}

	fmt.Println("app started")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("stopped server with: %s", err.Error())
	}

	os.Exit(1)
}

func getDbDsn(cfg config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)
}
