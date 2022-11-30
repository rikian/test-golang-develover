package main

import grpcApp "go/service1/grpc-app"

func main() { grpcApp.NewListenerImpl().Run() }
