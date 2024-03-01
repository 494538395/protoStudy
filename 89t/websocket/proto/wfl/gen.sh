# 在 proto 目录下执行
protoc --go_out=plugins=grpc:./  wfl.proto
protoc --go_out=plugins=grpc:./  packet.proto
protoc --go_out=plugins=grpc:./  hamster.proto
protoc --go_out=plugins=grpc:./  common.proto
