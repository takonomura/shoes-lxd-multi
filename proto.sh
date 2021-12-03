#!/bin/bash -x

rm -rf ./tmp
mkdir -p ./tmp/proto.go
protoc -I . --go_out=tmp/proto.go --go-grpc_out=tmp/proto.go proto/*.proto
mv tmp/proto.go/github.com/whywaita/shoes-lxd-multi/proto.go/shoeslxdmulti/* ./proto.go
rm -rf ./tmp
