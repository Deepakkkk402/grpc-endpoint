package main

import(
	"context"
	pb "grpc/protoc"

	"log"
	"io"
	"time"
	

	

	"google.golang.org/grpc"
	
) 

const (
	address = "localhost:8081"
)

func main() {
	//establish connection with grpc
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//to close the program

	defer conn.Close()
	//newclient to interact with grpc
	client := pb.NewUserClient(conn)
//to run the user methods to interact with grpc
	RunGetUsers(client)
	RunGetUser(client, "1")
	RunCreateUser(client, "24325645", "John Doe", "New York", 1234567890, 5.8, true)
    RunUpdateUser(client, "98498081", "24325645", "Jane Smith", "Los Angeles", 9876543210, 5.5, false)
    RunDeleteUser(client, "98498081")
}
// function to reterive all the user data from grpc
func RunGetUsers(client pb.UserClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUsers(_) = _, %v", client, err)
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetUsers(_) = _, %v", client, err)
	}
	log.Printf("UserInfo: %v", user)
}
}
//function to get the user according to the id form the grpc
func RunGetUser(client pb.UserClient, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: userId}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUser(_) = _, %v", client, err)
	}
	log.Printf("UserInfo: %v", res)
}
//function to create new user on grpc
func RunCreateUser(client pb.UserClient, id, fname, city string, phone int64, height float32, married bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UserInfo{
		Id: id,
		Fname: fname,
		City: city,
		Phone: phone,
		Height: height,
		Married: married,
	}

	//send request to the sever to add the user
	res, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateUser(_) = _, %v", client, err)
	}
	//after succesfull
	log.Printf("CreateUser Id: %v", res)

}
//function to update the user on grpc
func RunUpdateUser(client pb.UserClient, userId, id, fname, city string, phone int64, height float32, married bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UserInfo{
		Id: id,
		Fname: fname,
		City: city,
		Phone: phone,
		Height: height,
		Married: married,
	}
	//send the request the server to update
	res, err := client.UpdateUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateUser(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1{
		log.Printf("Update is done")
	}else {
	log.Printf("UpdateUser not Successful")
	}
}
//function to delete the user on the grpc using the userid
func runDeleteUser(client pb.UserClient, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: userId}
	//send the request to the grpc
	res, err := client.DeleteUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteUser(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1{
		log.Printf("Deleteuser suceess")
	}else{
	log.Printf("DeleteUser not Successful")
}
}
