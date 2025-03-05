package server

import (
	"context"
	pb "grpc/server/addressbook"
	"log"

	"google.golang.org/protobuf/proto"
)

type addressbookServer struct {
	pb.UnimplementedAddressBookServiceServer
}

func NewAddressBookServer() *addressbookServer { // go has no constructor so we make function which return an object
	return &addressbookServer{}
}

// ListPerons function take two arg ctx and req and retrun listpersonresponse and error ..... the first part of it which is (s *addressbookServer) is like self or this
func (s *addressbookServer) ListPersons(ctx context.Context, req *pb.ListPersonRequest) (*pb.ListPersonResponse, error) {
	log.Println("listing persons.....")
	return &pb.ListPersonResponse{}, nil
}

func (s *addressbookServer) AddPerson(ctx context.Context, person *pb.Person) (*pb.AddPersonResponse, error) {
	log.Println("adding person.....")
	response := &pb.AddPersonResponse{
		Person: person,
	}
	responseSize := proto.Size(response)
	log.Printf("Response: %v, Size: %d bits\n", response, responseSize)
	return response, nil
}
