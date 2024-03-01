package parse

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

// 根据文件路径，解析 proto
func ParseProtoWithFiles(filePath ...string) []*desc.FileDescriptor {
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

// ParseProtoWithDir 根据目录解析 proto 文件
func ParseProtoWithDirs(filePathDir ...string) []*desc.FileDescriptor {
	var fileDescriptors []*desc.FileDescriptor
	parser := protoparse.Parser{}

	for _, protoDir := range filePathDir {
		files, err := ioutil.ReadDir(protoDir)
		if err != nil {
			fmt.Println("Failed to read directory:", err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if filepath.Ext(file.Name()) != ".proto" {
				continue
			}
			filename := filepath.Join(protoDir, file.Name())
			fileDesc, err := parser.ParseFiles(filename)
			if err != nil {
				fmt.Printf("Failed to parse %s: %v\n", filename, err)
				continue
			}
			fileDescriptors = append(fileDescriptors, fileDesc...)
		}
	}

	return fileDescriptors
}
