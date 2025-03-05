package main

import (
	"context"
	pb "grpc/server/addressbook"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("couldnot connect to service")
	}
	defer conn.Close()
	client := pb.NewAddressBookServiceClient(conn)
	if _, err = client.AddPerson(context.Background(), &pb.Person{
		Name: "yehia",
		Id:   5,
	}); err != nil {
		log.Fatal("couldnot add person")
	}
}
