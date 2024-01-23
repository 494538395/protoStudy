package parse

import (
	"github.com/jhump/protoreflect/desc/protoparse"
)

func parseProto(filePath string) {
	var parser protoparse.Parser
	fileDescriptors, err := parser.ParseFiles("filePath")
	if err != nil {
		panic(err)
	}

}
