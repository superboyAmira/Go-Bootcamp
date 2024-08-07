TRANSMITTER_VER=transmitter_v1
RECEIVER_VER=receiver_v1
PROTOC_TRANSMITTER_DIR=./pkg/$(TRANSMITTER_VER)
PROTOC_RECEIVER_DIR=./pkg/$(RECEIVER_VER)

DC_NAME=docker-compose.yml

protobuf_install : 
	sudo apt-get install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export GO_PATH=~/go
	export PATH=$PATH:$(go env GOPATH)/bin

protoc_gen : proto_clear
	mkdir $(PROTOC_TRANSMITTER_DIR)
	# mkdir $(PROTOC_RECEIVER_DIR)


	protoc \
	--go_opt=paths=import \
	--go_out=./pkg/$(TRANSMITTER_VER) \
	--go-grpc_opt=paths=import \
	--go-grpc_out=./pkg/$(TRANSMITTER_VER) \
	./api/transmitter/transmitter-contract.proto
	mv $(PROTOC_TRANSMITTER_DIR)/team00/pkg/$(TRANSMITTER_VER)/* ./pkg/$(TRANSMITTER_VER)
	rm -rf $(PROTOC_TRANSMITTER_DIR)/team00

	# protoc \
	# --go_opt=paths=import \
	# --go_out=./pkg/$(RECEIVER_VER) \
	# --go-grpc_opt=paths=import \
	# --go-grpc_out=./pkg/$(RECEIVER_VER) \
	# ./api/receiver/receiver-contract.proto
	# mv $(PROTOC_RECEIVER_DIR)/team00/pkg/$(RECEIVER_VER)/* ./pkg/$(RECEIVER_VER)
	# rm -rf $(PROTOC_RECEIVER_DIR)/team00

proto_clear :
	rm -rf ./pkg/*

container_up :
	docker-compose -f $(DC_NAME) up

container_rm:
	docker-compose stop \
	&& docker-compose rm