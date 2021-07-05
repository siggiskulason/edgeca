export GOPATH="$HOME/go/"
export PATH=$PATH:$GOPATH/bin
go get github.com/99designs/gqlgen/cmd@v0.13.0

cd ../internal/server/graphqlimpl
go run github.com/99designs/gqlgen generate