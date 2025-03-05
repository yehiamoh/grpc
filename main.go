package main

import (
	"log"
	"net"

	"grpc/server"
	pb "grpc/server/addressbook"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("couldnot listen on port")
	}
	grpcServer := grpc.NewServer() // create instance of server from grpc package
	reflection.Register(grpcServer)
	pb.RegisterAddressBookServiceServer(grpcServer, server.NewAddressBookServer()) //then regitser the server to the services
	//the server is responsilbe for header compresion ,multiplexing,streaming
	log.Println("server started")
	grpcServer.Serve(lis)
}
