package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/tndgoat/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer is the server that provides laptop services.
type LaptopServer struct {
	pb.UnimplementedLaptopServiceServer
	Store LaptopStore
}

// NewLaptopServer creates a new LaptopServer.
func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{
		Store: store,
	}
}

// CreateLaptop creates a new laptop and returns its ID.
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	log.Printf("Received a request to create a laptop with ID: %s", laptop.GetId())

	if len(laptop.GetId()) > 0 {
		// Validate the UUID format of the laptop ID.
		_, err := uuid.Parse(laptop.GetId())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Laptop ID is not a valid UUID: %v", err,
			)
		}
	} else {
		// Generate a new UUID for the laptop ID.
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				"Cannot generate a new laptop ID: %v", err,
			)
		}
		laptop.Id = id.String()
	}

	// Save the laptop to the store.
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(
			code,
			"Cannot save laptop to the store: %v", err,
		)
	}

	log.Printf("Laptop with ID %s has been saved successfully.", laptop.GetId())

	res := &pb.CreateLaptopResponse{
		Id: laptop.GetId(),
	}

	return res, nil
}
