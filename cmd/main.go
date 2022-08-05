package main

import (
	"fmt"
	"net"
	"os"

	_ "github.com/lib/pq"
	"github.com/w1kend/parcel_delivery/internal/app/parcel_delivery"
	"github.com/w1kend/parcel_delivery/internal/config"
	"github.com/w1kend/parcel_delivery/internal/pkg/auth"
	"github.com/w1kend/parcel_delivery/internal/pkg/repositories"
	"github.com/w1kend/parcel_delivery/pkg/parcel_delivery_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	tokenManager := auth.NewTokenManager(config.GetValue(config.JWTSecret).String())
	authInterceptor := auth.NewAuthInterceptor(tokenManager, map[string]interface{}{
		"/parcel_delivery.ParcelDelivery/SignIn": nil,
		"/parcel_delivery.ParcelDelivery/SignUp": nil,
	})

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)
	reflection.Register(grpcServer)

	db, err := config.NewPostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

	fmt.Println("app started")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("stopped server with: %s", err.Error())
	}

	os.Exit(1)
}
