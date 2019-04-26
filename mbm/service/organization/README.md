# grpc_go_service_example

## About gRPC Go Organization Service Example

This example showcases using gRPC with Go to define and implement a simple Organization server and client.

The Organization gRPC service defines methods to:
- Add an Organization
- List all Organizations
- Add a User to an Organization
- List all Users from all Organizations
- List all Users for a specific Organization

The solution is comprised of the following applications:
- Organization Server which hosts the Organization gRPC services
- Organization Client which provides a user access to exercise the Organization services


## Prerequisites

The example requires the following already be installed:
- Go 1.12.1 or higher (golang.org)
- Go support for Protocol Buffers (https://github.com/golang/protobuf)
- gRPC package (google.golang.org/grpc)
- UUID packabe (github.com/twinj/uuid)


## Source
The code for the solution can be found under the /mbm/service/organization directory. Under this directory you can find:

/orgserver		Organization Server

/organization 		Organization gRPC proto package

/orgclient		Organization Client


Build: You can rebuild the Organization Server, Organization Client and/or Organization proto code by running 'go build' in the respective directories.

Test: You can run the Organization Server unit tests by running 'go test' in the orgserver directory. 

NOTE: This example uses transient data for the Organization and User lists. 

## Usage
To start the Organization Server and run it in the background, go to the /orgserver directory, build the org server and run:

	./orgserver &

Note: This will display the pid (process id) of the orgserver process. If you need to stop the Organization Server, use the kill command with the pid of the orgserver process (eg, kill <pid>).


Using the Organization Client:
To run the Organization Client, go to the /orgclient directory, ensure the client is built, then run the client:

	./orgclient


This will display the list of commands understood by the Organization Client, including:

	List all Organizations - displays each Organization's Organization ID, Name and Description:
		./orgclient orgs
	 
	Create an Organization w/ name and optional description - displays the new Organization's ID, Name and Description:
		./orgclient orgs -name=<organizationname>  [-desc=<description>]
	
	List all Users - displays each User's User ID, Name and Organization IDs:
		./orgclient users

	List all Users for a specific Organization - displayed each User's User ID and Name:
		./orgclient users -orgid=<organizationid>
	
	Add a User to a specific Organization  - displays the new User's User ID, Name and Organization ID:
		./orgclient users -name=<username> -orgid=<organizationid>


Note: In the above use cases, the Organization ID is displayed when you successfully create an Organization as well as when you list the Organizations. The Organization ID is a unique ID that must be used when displaying Users for a specific Organization or when adding a new User to an Organization.

