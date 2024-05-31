package main

import (
	"context"

	"log"
	"net"
	"strconv"
	"math/rand"

	pb "grpc/protoc" // alias for the generated protobuf code
	"google.golang.org/grpc"
	
)
const (
	port = ":8000"
)

//slice to store user data as pointer

var users []*pb.UserInfo
//struct to handle the server
type UserServer struct {
	pb.UnimplementedUserServer 
}


func main() {

	initUsers()
	//create tcp listener port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create new grpc server
	s := grpc.NewServer()
   //add to the grpc server
	pb.RegisterUserServer(s, &UserServer{})
//stop untill server is blocked
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initUsers() {
	//created user information
	user1 := &pb.UserInfo{
		Id: "1",
		Fname: "Steve",
		City: "LA",
		Phone: 1234567890,
		Height: 5.8,
		Married: true,
	}
	user2 := &pb.UserInfo{
		Id: "2",
		Fname: "Alice",
		City: "NY",
		Phone: 9876543210,
		Height: 5.5,
		Married: false,
	}

	users = append(users, user1)
	users = append(users, user2)
}
//getuser method to connect with userserver struct
func (s *UserServer) GetUsers(in *pb.Empty,
	stream pb.User_GetUsersServer) error {
	log.Printf("Received: %v", in)
	for _, user := range users {
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}
//getuser method to connect with userserver struct to fetch the data by id 
func (s *UserServer) GetUser(ctx context.Context,
	in *pb.Id) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)

	res := &pb.UserInfo{}

	for _, user := range users {
		if user.GetId() == in.GetValue() {
			res = user
			break
		}
	}

	return res, nil
}
//createuser method to connect with userserver to create new user and generating new rand id
func (s *UserServer) CreateUser(ctx context.Context,
	in *pb.UserInfo) (*pb.Id, error) {
	log.Printf("Received: %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(10))
	in.Id = res.GetValue()
	users = append(users, in)
	return &res, nil
}

//updateuser method is to connect with the userserver update the data 
func (s *UserServer) UpdateUser(ctx context.Context,
	in *pb.UserInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for _, user := range users {
		if user.GetId() == in.GetId() {
			//append is used to add data into the userapp

			users = append(users[:index], users[index+1:]...)
			in.Id = user.GetId()
			users = append(users, in)
			res.Value = 1
			break
		}
	}

	return &res, nil
}
//deleteuser is connect with userserver to delete tha data using the ids
func (s *UserServer) DeleteUser(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	res := pb.Status{}
	for _, user := range users {
		if user.GetId() == in.GetValue() {
			users = append(users[:index], users[index+1:]...)
			res.Value = 1
			break
		}
	}

	return &res, nil
}
