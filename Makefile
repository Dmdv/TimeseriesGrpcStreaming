
proto:
	protoc -I./grpc/schema --go_out=plugins=grpc:grpc/schema stream.proto
	mockgen -source=./grpc/schema/stream.pb.go -destination ./grpc/mock/mock_stream.go -package mock_stream

gen:
	swag init

lint:
	golangci-lint run

serve:
	WEB__LISTEN=0.0.0.0:8000 \
	DB__HOST=localhost \
	go run main.go serve