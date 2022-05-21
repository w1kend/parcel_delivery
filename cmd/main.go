package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"github.com/w1kend/guavapay_test/internal/app/parcel_delivery"
	"github.com/w1kend/guavapay_test/internal/pkg/api/parcel_delivery_grpc"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()

	db, err := sql.Open("postgres", getDbDsn())
	if err != nil {
		panic(err)
	}

	impl := parcel_delivery.NewImplementation(db)

	parcel_delivery_grpc.RegisterParcelDeliveryServer(grpcServer, &impl)

	lis, err := net.Listen("tcp", ":8000")
	fmt.Println("listen on port 8000")
	if err != nil {
		panic(err)
	}

	fmt.Println("app started")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("stopped server with: %s", err.Error())
	}

	os.Exit(1)
}

func getDbDsn() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
