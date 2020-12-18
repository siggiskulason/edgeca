export GOPATH="$HOME/go/"
export PATH=$PATH:$GOPATH/bin

go get google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
export GO111MODULE=on # Enable module mode
#go get -u google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
rm edgeca_*.go

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative edgeca.proto
