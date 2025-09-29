mkdir -p manager-api/proto
mkdir -p manager-api/docs
rm -rf .manager-api/proto/*
rm -rf .manager-api/docs/*
protoc \
    -I=. \
    --go_out=. \
    --go-grpc_out=. \
    --grpc-gateway_out=. \
    --openapiv2_out=manager-api/docs \
    --openapiv2_opt logtostderr=true \
    manager-api/api/*.proto \