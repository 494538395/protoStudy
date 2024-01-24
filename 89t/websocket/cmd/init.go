package main

import (
	"protoStudy/89t/websocket/parse"
)

func init() {
	pbfiles := []string{
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\wfl.proto",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\packet.proto",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\hamster.proto"}
	fds = parse.ParseProto(pbfiles...)
}
