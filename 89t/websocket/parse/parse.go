package parse

import (
	"fmt"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

func ParseProto(filePath ...string) []*desc.FileDescriptor {
	var parser protoparse.Parser
	fileDescriptors, err := parser.ParseFiles(filePath...)
	if err != nil {
		panic(err)
	}

	fs := fileDescriptors[0]

	for _, msgDescriptor := range fs.GetMessageTypes() {
		fmt.Println(msgDescriptor.GetName())
		for _, fieldDesc := range msgDescriptor.GetFields() {
			fmt.Println("\t", fieldDesc.GetType().String(), fieldDesc.GetName())
		}
		fmt.Println()
	}
	for _, enumDescriptor := range fs.GetEnumTypes() {
		fmt.Println(enumDescriptor.GetName())
		for _, valueDescriptor := range enumDescriptor.GetValues() {
			fmt.Println("\t", valueDescriptor.GetName())
		}
		fmt.Println()
	}

	return fileDescriptors
}
