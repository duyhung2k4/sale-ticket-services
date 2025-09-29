mkdir -p gen
rm -rf ./gen/*
protoc \
    -I=. \
    --go_out=. \
    --go-grpc_out=. \
    --grpc-gateway_out=. \
    --openapiv2_out=docs \
    --openapiv2_opt logtostderr=true \
    api/*.proto \