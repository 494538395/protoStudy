package main

import (
	"fmt"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

var fds []*desc.FileDescriptor

func init() {
	parser := protoparse.Parser{ImportPaths: []string{
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\jerry",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\common",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\wfl",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\partyBattle",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\realCommon",
	}}
	pbfiles := []string{
		"wfl.proto",
		"packet.proto",
		"hamster.proto",
		"common.proto",
		"realCommon.proto",
		"hamster.proto",
	}

	for _, pbfile := range pbfiles {
		tmpFds, err := parser.ParseFiles(pbfile)
		if err != nil {
			panic(err)
		}
		fds = append(fds, tmpFds...)
	}
	fmt.Println(fds)
}
