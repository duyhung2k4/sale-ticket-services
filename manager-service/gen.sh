PROTO_PATH=api
OUT_PATH=gen

protoc \
    -I=$PROTO_PATH \
   --go_out=$OUT_PATH --go_opt=paths=source_relative \
   --go-grpc_out=$OUT_PATH --go-grpc_opt=paths=source_relative \
   --grpc-gateway_out=$OUT_PATH --grpc-gateway_opt=paths=source_relative \
   $PROTO_PATH/*.proto