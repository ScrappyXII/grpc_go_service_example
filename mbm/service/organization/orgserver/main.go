// Package main implements the Organization Server hosting the Organization gRPC services
package main

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	pb "grpc_go_service_example/mbm/service/organization/organization"

	"github.com/twinj/uuid"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// UserSchema definition
type UserSchema struct {
	id    string // UUID of the user
	orgid string // UUID of the user's organization
	name  string // name of the user
}

// OrganizationSchema definition
type OrganizationSchema struct {
	id          string // UUID of the organization
	name        string // name of the organization
	description string // description of the organization
}

//TODO: in future, move from transient data to DB or other storage

var userTable []UserSchema
var orgTable []OrganizationSchema

// server is used to implement the Organization Server.
type server struct{}

var mutex = &sync.Mutex{}

// CreateOrganization implements adding an organization in the Organization Server
func (s *server) CreateOrganization(ctx context.Context, in *pb.CreateOrganizationRequest) (*pb.OrganizationResponse, error) {
	log.Println("Server: Received Organization Creation for: " + in.Name)

	if in.Name == "" {
		log.Println("Server: Organization name was not supplied")
		return nil, errors.New("Organization name is required")
	}

	mutex.Lock()
	defer mutex.Unlock()

	// TODO: in future, when moving to a real DB model, may want to prevent duplication of the organization name. For now, just allow it.

	u := uuid.NewV4().String()
	orgTable = append(orgTable, OrganizationSchema{u, in.Name, in.Description})
	log.Println("Server: Created Organization " + in.Name + " with description " + in.Description)

	return &pb.OrganizationResponse{Id: u, Name: in.Name, Description: in.Description}, nil
}

// FetchOrganizationList implements listing all organizations from the Organization Server
func (s *server) FetchOrganizationList(ctx context.Context, in *pb.Empty) (*pb.OrganizationListResponse, error) {
	log.Println("Server: Received Request to List All Organizations")

	mutex.Lock()
	defer mutex.Unlock()

	orgResponseList := new(pb.OrganizationListResponse)
	// iterate over all Orgs
	for i := 0; i < len(orgTable); i++ {
		// add organization to list
		log.Println("Server: Found Org: " + orgTable[i].name)
		orgResponseList.Organizations = append(orgResponseList.Organizations, &pb.OrganizationResponse{Id: orgTable[i].id, Name: orgTable[i].name, Description: orgTable[i].description})
	}

	return orgResponseList, nil
}

// CreateUser implements adding a user in the Organization Server for a specific organization
func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	log.Println("Server: Received User Creation for: " + in.Name + "(" + in.OrganizationId + ")")

	if in.OrganizationId == "" {
		log.Println("Server: Organization ID was not supplied")
		return nil, errors.New("Organization ID is required")
	}
	if in.Name == "" {
		log.Println("Server: User name was not supplied")
		return nil, errors.New("User name is required")
	}
	// verify that organization ID supplied is an existing organization
	orgValid := false
	for i := 0; i < len(orgTable); i++ {
		if orgTable[i].id == in.OrganizationId {
			orgValid = true
			break
		}
	}
	if orgValid == false {
		log.Println("Server: Organziation ID doesn't exist")
		return nil, errors.New("Valid Organization ID is required")
	}

	// TODO: in future, when moving to a real DB model, should check to ensure no duplication of the user name in the given organzation. For now, just allow it and assign new UUID.

	mutex.Lock()
	defer mutex.Unlock()

	u := uuid.NewV4().String()
	userTable = append(userTable, UserSchema{u, in.OrganizationId, in.Name})
	return &pb.UserResponse{Id: u, Name: in.Name, OrganizationId: in.OrganizationId}, nil
}

// FetchUserList implements listing all users from the Organization Server
func (s *server) FetchUserList(ctx context.Context, in *pb.Empty) (*pb.UserListResponse, error) {
	log.Println("Server: Received Request to List All Users")

	mutex.Lock()
	defer mutex.Unlock()

	userResponseList := new(pb.UserListResponse)
	// iterate over all Users
	for i := 0; i < len(userTable); i++ {
		// add user to list
		log.Println("Server: Found User: " + userTable[i].name)
		userResponseList.Users = append(userResponseList.Users, &pb.UserResponse{Id: userTable[i].id, Name: userTable[i].name, OrganizationId: userTable[i].orgid})
	}

	return userResponseList, nil
}

// FetchUserListByOrganization implements listing all users from the Organization Server for a specific organization
func (s *server) FetchUserListByOrganization(ctx context.Context, in *pb.ByOrganizationRequest) (*pb.UserListResponse, error) {
	log.Printf("Server: Received Request to List All Users for Organization ID: " + in.OrganizationId)

	if in.OrganizationId == "" {
		log.Println("Server: Organization ID was not supplied")
		return nil, errors.New("Organization ID is required")
	}

	mutex.Lock()
	defer mutex.Unlock()

	// TODO: in future, verify organization ID is a valid organization. For now, the result is simply empty which is also an accurate result.

	userResponseList := new(pb.UserListResponse)
	// iterate over the Users looking for ones that are part of the specified Org
	for i := 0; i < len(userTable); i++ {
		if userTable[i].orgid == in.OrganizationId {
			// add user to list
			log.Println("Server: Found User: " + userTable[i].name)
			userResponseList.Users = append(userResponseList.Users, &pb.UserResponse{Id: userTable[i].id, Name: userTable[i].name, OrganizationId: userTable[i].orgid})
		}
	}

	return userResponseList, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOrganizationServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
