protoc \
    -I=. \
   --go_out=. \
   --go-grpc_out=. \
   --grpc-gateway_out=. \
   api/*.proto