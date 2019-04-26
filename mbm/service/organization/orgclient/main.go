// Package main implements an Organization command-line client for the Organization service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	pb "grpc_go_service_example/mbm/service/organization/organization"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	cmdlineOrg  = "orgs"
	cmdlineUser = "users"
)

func main() {
	// Set up flags for command line to determine what action(s) we are taking
	// Options include:
	//  orgs			: lists all organizations
	//  orgs -name=<orgname>  [-desc=<description>]		: create organization with given name
	//  users [-orgid=orgid]				: list all users or users for a specific organization
	//  users -name=username -orgid=orgid	: add user to a specific organization
	orgCommand := flag.NewFlagSet(cmdlineOrg, flag.ExitOnError)
	orgCreateFlag := orgCommand.String("name", "", "Organization Name to be Added")
	orgDescriptionFlag := orgCommand.String("desc", "", "Organization Description")

	userCommand := flag.NewFlagSet(cmdlineUser, flag.ExitOnError)
	userAddFlag := userCommand.String("name", "", "User Name to be added")
	userOrgFlag := userCommand.String("orgid", "", "Organization ID")

	if len(os.Args) == 1 {
		fmt.Println("Organization Client Command:")
		fmt.Println("orgs    :Lists all Organizations")
		fmt.Println("orgs -name=<orgname>  [-desc=<description>]   :Create Organization with given name and description")
		fmt.Println("users [-orgid=<orgid>]	     :List all Users or Users for a specific Organization")
		fmt.Println("users -name=<username> -orgid=<orgid>		:Add User to a specific Organization")
		return
	}

	// Set up a connection to the Organization server.
	// NOTE: in future, the connection should support AuthN/Z to determine if/what that client can do
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrganizationClient(conn)

	// parse command line using the flag sets
	switch os.Args[1] {
	case "orgs":
		orgCommand.Parse(os.Args[2:])
	case "users":
		userCommand.Parse(os.Args[2:])
	default:
		fmt.Println(os.Args[1] + " is not valid command.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if orgCommand.Parsed() {
		if *orgCreateFlag != "" {
			// Create Org with Org name and description
			response, err := c.CreateOrganization(ctx, &pb.CreateOrganizationRequest{Name: *orgCreateFlag, Description: *orgDescriptionFlag})
			if err != nil {
				log.Fatalf("could not create organization: %v", err)
			}
			fmt.Println("Successfully added organization - ID: " + response.Id + " Name: " + response.Name + " Description: " + response.Description)
		} else {
			// Get list of all Organizations
			response, err := c.FetchOrganizationList(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Could not obtain organization list: %v", err)
			}
			// Print out the list of all Organizations
			orgList := response.GetOrganizations()
			if len(orgList) > 0 {
				fmt.Println("Organization List: ")
				for i := 0; i < len(orgList); i++ {
					fmt.Println("  -> ID:" + orgList[i].Id + "  Name:" + orgList[i].Name + "   Description:" + orgList[i].Description)
				}
			} else {
				fmt.Println("No Organizations")
			}
		}
	}

	if userCommand.Parsed() {
		if *userAddFlag != "" {
			if *userOrgFlag != "" {
				// Add a User to a specified Organization
				response, err := c.CreateUser(ctx, &pb.CreateUserRequest{Name: *userAddFlag, OrganizationId: *userOrgFlag})
				if err != nil {
					fmt.Println("Could not create user: " + err.Error())
				} else {
					fmt.Println("Successfully added User - ID:" + response.Id + " Name:" + response.Name + " OrgID:" + response.OrganizationId)
				}
			} else {
				fmt.Println("Cannot add a User without an Organization ID")
			}
		} else {
			if *userOrgFlag == "" {
				// Get list of all Users across all Organizations
				response, err := c.FetchUserList(ctx, &pb.Empty{})
				if err != nil {
					log.Fatalf("Could not obtain user list: %v", err)
				}

				// Print out the list of all Users
				userList := response.GetUsers()
				if len(userList) > 0 {
					fmt.Println("User List for All Organizations: ")
					for i := 0; i < len(userList); i++ {
						fmt.Println("  -> ID:" + userList[i].Id + "  Name:" + userList[i].Name + "   OrgID:" + userList[i].OrganizationId)
					}
				} else {
					fmt.Println("No Users")
				}
			} else {
				// Get list of Users for specified Org
				response, err := c.FetchUserListByOrganization(ctx, &pb.ByOrganizationRequest{OrganizationId: *userOrgFlag})
				if err != nil {
					log.Fatalf("Could not obtain user list: %v", err)
				}

				// Print out the list of Users for the specific Org
				userList := response.GetUsers()
				if len(userList) > 0 {
					fmt.Println("User List for Organization: " + *userOrgFlag)
					for i := 0; i < len(userList); i++ {
						fmt.Println("  -> ID:" + userList[i].Id + "  Name:" + userList[i].Name)
					}
				} else {
					fmt.Println("No Users")
				}
			}
		}
	}

}
