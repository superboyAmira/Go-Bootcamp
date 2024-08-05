TRANSMITTER_VER=transmitter_v1
PROTOC_TRANSMITTER_DIR=./pkg/$(TRANSMITTER_VER)


protobuf_install : 
	sudo apt-get install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	export GO_PATH=~/go
	export PATH=$PATH:/$GO_PATH/bin

protoc_gen :
	mkdir $(PROTOC_TRANSMITTER_DIR)
	protoc \
	--go_opt=paths=import \
	--go_out=./pkg/$(TRANSMITTER_VER) \
	--go-grpc_opt=paths=import \
	--go-grpc_out=./pkg/$(TRANSMITTER_VER) \
	./api/transmitter/transmitter-contract.proto
	mv $(PROTOC_TRANSMITTER_DIR)/team00/pkg/$(TRANSMITTER_VER)/* ./pkg/$(TRANSMITTER_VER)
	rm -rf $(PROTOC_TRANSMITTER_DIR)/team00