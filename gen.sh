#!/bin/sh

CLIENT_OUTDIR=client/lib
SERVER_OUTPUT_DIR=server/hello

mkdir -p ${CLIENT_OUTDIR} ${SERVER_OUTPUT_DIR}

# protocol/helloworld.proto
#  --js_out => helloworld_pb.js
#  --grpc-web_out => helloworld_pb.d.ts
#  --grpc-web_out => HelloworldServiceClientPb.ts
#  --go_out => helloworld.pb.go
protoc --proto_path=proto hello.proto \
    --js_out=import_style=commonjs:${CLIENT_OUTDIR} \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:${CLIENT_OUTDIR} \
    --go_out=plugins=grpc:${SERVER_OUTPUT_DIR}
