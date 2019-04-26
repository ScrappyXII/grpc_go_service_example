package main

import (
	"context"
	pb "mbm/service/organization/organization"
	"testing"
)

type OrganizationTestSchema struct {
	id          string   // UUID of the organization
	name        string   // name of the organization
	description string   // description of the organization
	users       []string // user list
}

// populate test data
var orgs = []OrganizationTestSchema{
	OrganizationTestSchema{
		id:          "",
		name:        "Some Org",
		description: "An org with a few users",
		users: []string{
			"John Doe",
			"Jane Doe",
			"Another Doe",
		},
	},
	OrganizationTestSchema{
		id:          "",
		name:        "MBM Software",
		description: "Consulting Company",
		users: []string{
			"Marc B Manza",
		},
	},
	OrganizationTestSchema{
		id:          "",
		name:        "Plain Old Company",
		description: "",
	},
}

var s server
var ctx context.Context
var orgscreated = false
var userscreated = false

func TestCreateOrganization(t *testing.T) {
	if orgscreated {
		return
	}

	// iterate through
	for i := 0; i < len(orgs); i++ {
		// [START populate_proto]
		p := pb.CreateOrganizationRequest{
			Name:        orgs[i].name,
			Description: orgs[i].description,
		}
		// [END populate_proto]

		resp, err := s.CreateOrganization(ctx, &p)

		if err != nil {
			t.Errorf("Could not create organization: %v", err)
		} else {
			// ensure we got back the same org data we sent as well as the assigned orgID
			if (resp.Name != orgs[i].name) || (resp.Description != orgs[i].description) || (resp.Id == "") {
				t.Errorf("Create Organization failed! Wanted: " + orgs[i].name + " " + orgs[i].description + "  Got: " + resp.Name + " " + resp.Description + " " + resp.Id)
			}
			// save returned UUID so we can use it in other tests such adding users
			orgs[i].id = resp.Id
		}
	}

	// try adding an incorrect org name (empty name)
	p := pb.CreateOrganizationRequest{
		Name: "",
	}
	resp, err := s.CreateOrganization(ctx, &p)
	if (err == nil) || (resp != nil) {
		t.Errorf("Create Organization failed! Should not accept empty organization name but it did")
	}

	orgscreated = true
}

func TestListOrganizations(t *testing.T) {
	// [START populate_proto]
	p := pb.Empty{}
	// [END populate_proto]

	if !orgscreated {
		resp, err := s.FetchOrganizationList(ctx, &p)
		// list should be empty at this point
		if err != nil {
			t.Errorf("Could not obtain organization list: %v", err)
		} else {
			orgList := resp.GetOrganizations()
			if len(orgList) != 0 {
				t.Errorf("List Orgs error! Wanted: default empty list  Got: a non-empty list")
			}
		}
	}

	// now add the test orgs and retry. we should get a list now
	TestCreateOrganization(t)
	resp, err := s.FetchOrganizationList(ctx, &p)
	if err != nil {
		t.Errorf("Could not obtain organization list: %v", err)
	} else {
		// should get back the same number of orgs we put in. TestCreateOrganization already checked that we added these correctly so not rechecking values here but we could in the future.
		orgList := resp.GetOrganizations()
		if len(orgList) != len(orgs) {
			t.Errorf("List Orgs failed! Did not get same list back as was added.")
		}
	}
}

func TestCreateUser(t *testing.T) {
	// populate organizations
	TestCreateOrganization(t)

	if userscreated {
		return
	}

	// iterate through all test orgs so we can create users for each
	for orgcount := 0; orgcount < len(orgs); orgcount++ {
		// iterate through Users for the given org
		for i := 0; i < len(orgs[orgcount].users); i++ {
			// [START populate_proto]
			p := pb.CreateUserRequest{
				Name:           orgs[orgcount].users[i],
				OrganizationId: orgs[orgcount].id,
			}
			// [END populate_proto]

			resp, err := s.CreateUser(ctx, &p)

			if err != nil {
				t.Errorf("Could not create user: %v", err)
			} else {
				// ensure we got back the same user data we sent as well as the assigned userID
				if (resp.Name != orgs[orgcount].users[i]) || (resp.OrganizationId != orgs[orgcount].id) || (resp.Id == "") {
					t.Errorf("Create User failed! Wanted: " + orgs[orgcount].users[i] + " " + orgs[orgcount].id + "  Got: " + resp.Name + " " + resp.OrganizationId + " " + resp.Id)
				}
			}
		}

		// try adding an invalid empty user name for this org
		p := pb.CreateUserRequest{
			Name:           "",
			OrganizationId: orgs[orgcount].id,
		}
		resp, err := s.CreateUser(ctx, &p)
		if (err == nil) || (resp != nil) {
			t.Errorf("Create User failed! Should not accept empty user name but it did")
		}
	}

	// try adding a user with an empty orgID
	p := pb.CreateUserRequest{
		Name:           "Dummy User",
		OrganizationId: "",
	}
	resp, err := s.CreateUser(ctx, &p)
	if (err == nil) || (resp != nil) {
		t.Errorf("Create User failed! Should not accept empty organizationID but it did")
	}

	// try adding a user with a non-existant orgID
	p = pb.CreateUserRequest{
		Name:           "Dummy User",
		OrganizationId: "abc123",
	}
	resp, err = s.CreateUser(ctx, &p)
	if (err == nil) || (resp != nil) {
		t.Errorf("Create User failed! Should require an existing organizationID but it did not")
	}
	userscreated = true
}

func TestListAllUsers(t *testing.T) {
	// [START populate_proto]
	p := pb.Empty{}
	// [END populate_proto]

	if !userscreated {
		resp, err := s.FetchUserList(ctx, &p)
		// list should be empty at this point
		if err != nil {
			t.Errorf("Could not obtain user list: %v", err)
		} else {
			userList := resp.GetUsers()
			if len(userList) != 0 {
				t.Errorf("List Users error! Wanted: default empty list  Got: a non-empty list")
			}
		}
	}

	// now add the test orgs and users and retry. we should get a list now
	TestCreateUser(t)
	resp, err := s.FetchUserList(ctx, &p)
	if err != nil {
		t.Errorf("Could not obtain user list: %v", err)
	} else {
		// should get back the same number of users we put in. TestCreateUser (and TestCreateOrganzartion) already checked that we added these correctly so not rechecking values here but we could in the future.
		orgList := resp.GetUsers()

		// count how many users have across all orgs
		userCount := 0
		for i := 0; i < len(orgs); i++ {
			userCount += len(orgs[i].users)
		}

		if len(orgList) != userCount {
			t.Errorf("List Users failed! Did not get same list back as was added.")
		}
	}
}

func TestListOrganizationUsers(t *testing.T) {
	// load up our test data
	TestCreateUser(t)

	// iterate through all test orgs so we can list users for each
	for orgcount := 0; orgcount < len(orgs); orgcount++ {
		// [START populate_proto]
		p := pb.ByOrganizationRequest{
			OrganizationId: orgs[orgcount].id,
		}
		// [END populate_proto]

		resp, err := s.FetchUserListByOrganization(ctx, &p)
		if err != nil {
			t.Errorf("Could not obtain user list: %v", err)
		} else {
			// should get back the same number of users we put in. TestCreateUser (and TestCreateOrganzartion) already checked that we added these correctly so not rechecking values here but we could in the future.
			userList := resp.GetUsers()

			if len(userList) != len(orgs[orgcount].users) {
				t.Errorf("List Users failed! Did not get same list back as was added.")
			}
		}
	}

	// try listing users for an empty orgID
	p := pb.ByOrganizationRequest{
		OrganizationId: "",
	}
	resp, err := s.FetchUserListByOrganization(ctx, &p)
	if (err == nil) || (resp != nil) {
		t.Errorf("List User failed! Should require an organizationID be supplied but it did not")
	}
}
