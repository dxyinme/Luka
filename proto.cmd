protoc --proto_path=proto --go_out=plugins=grpc:proto proto/Keeper.proto
protoc --proto_path=proto --go_out=plugins=grpc:proto proto/RemoteCall.proto
