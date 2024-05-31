This is a gRPC service for managing user information. It allows clients to retrieve, create, update, and delete user data.
creating endpoint to access the app using this grpc server


Build a endpoint with gRPC and Golang
Steps ($:represents terminal commands)
$mkdir grpc-go
$cd grpc-go
$mkdir protoc
$touch api.proto
Update api.proto
$cd ..
$go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$go mod init grpc
$export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go-grpc_out=. --go_out=. api.proto
$go mod tidy
$touch server/main.go
$touch client/main.go

Update server/main.go
$gofmt -s -w server/main.go
$go run server/main.go
#go run client/main.go
