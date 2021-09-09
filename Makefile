# go env
# export PATH="$PATH:$(go env GOPATH)/bin"

gen:
	protoc  --proto_path=proto proto/*.proto --go-grpc_out=./pb --go_out=. proto/*.proto

clean:
	rm pb/*.go	

test:
	go test -cover -race ./...