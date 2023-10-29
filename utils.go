package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UsersGrpcConnection() (*grpc.ClientConn, context.Context, context.CancelFunc, error) {
	conn, err := grpc.Dial("users:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("Error while dialing users service: %v", err)
		return nil, nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return conn, ctx, cancel, nil
}

func ProductsGrpcConnection() (*grpc.ClientConn, context.Context, context.CancelFunc, error) {
	conn, err := grpc.Dial("products:50002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("Error while dialing products service: %v", err)
		return nil, nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return conn, ctx, cancel, nil
}

func ReviewsGrpcConnection() (*grpc.ClientConn, context.Context, context.CancelFunc, error) {
	conn, err := grpc.Dial("reviews:50003", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("Error while dialing reviews service: %v", err)
		return nil, nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return conn, ctx, cancel, nil
}
