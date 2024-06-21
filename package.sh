export GOOS=linux
export GOARCH=amd64
go mod tidy
go build -o bootstrap ./main