package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

func TestImport(t *testing.T) {
	var newFds []*desc.FileDescriptor
	parser := protoparse.Parser{ImportPaths: []string{
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\jerry",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\common",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\wfl",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\realCommon",
		"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\partyBattle",
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
		newFds = append(newFds, tmpFds...)
	}

	newFds, err := parser.ParseFiles("realCommon.proto", "common.proto")
	if err != nil {
		panic(err)
	}

	fmt.Println(newFds)

	//parser := protoparse.Parser{LookupImport: func(imp string) (*desc.FileDescriptor, error) {
	//	if imp == "google/protobuf/descriptor.proto" {
	//		return desc.LoadFileDescriptor("D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\common.proto")
	//	}
	//
	//	return nil, errors.New("unexpected filename")
	//}}

	// 定义要解析的 proto 文件描述符列表
	//fileNames := []string{
	//	"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\common.proto",
	//	//"D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\hamster.proto",
	//}
	//
	//// 解析 proto 文件
	//protoFiles, err := parser.ParseFiles(fileNames...)
	//if err != nil {
	//	panic(err)
	//}
	//
	//newFds, err := parser.ParseFiles("D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\hamster.proto")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(newFds)
	//
	//fmt.Println(protoFiles)
}

//func TestMarshall(t *testing.T) {
//	pbBytes, err := proto.Marshal(&pb.RecruitQueryPoolResp{
//		Id:          500,
//		CardIds:     []int32{1, 2, 3},
//		LuckCardIds: []int32{4, 5, 6},
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	jsonBytes, err := json.Marshal(pbBytes)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(string(jsonBytes))
//
//	var tmp []byte
//
//	err = json.Unmarshal(jsonBytes, &tmp)
//	if err != nil {
//		panic(err)
//	}
//
//	var tt pb.RecruitQueryPoolResp
//
//	err = proto.Unmarshal(tmp, &tt)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(tt)
//
//	err = json.Unmarshal(jsonBytes, &tt)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(tt)
//}
//
//func TestMultiNestedOneof(t *testing.T) {
//
//	msg := &pb.Resp{
//		Identifier: 1001,
//		Operation:  "admin",
//		MsgId:      "200",
//		Data: &pb.Resp_PullMsgBySeqListResp{PullMsgBySeqListResp: &pb.PullMsgBySeqListResp{
//			Name: "jerry",
//			Age:  21,
//			Data: &pb.PullMsgBySeqListResp_Pec{Pec: &pb.PersonC{
//				Name:    "sam",
//				Address: "beijing",
//			}},
//		},
//		},
//	}
//
//	marshal, err := proto.Marshal(msg)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println(marshal)
//
//	dynaMsg := dynamic.NewMessage(findMsgType(fds, "Resp"))
//
//	err = dynaMsg.Unmarshal(marshal)
//	if err != nil {
//		panic(err)
//	}
//
//	marshalJSON, err := dynaMsg.MarshalJSON()
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Println("marshalJSON-->", marshalJSON)
//
//}

// 测试解析含有 import 的 proto
func TestProtoImport(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println("dir-->", dir)

	var fileDescriptors []*desc.FileDescriptor
	parser := protoparse.Parser{}

	// 1
	fileDesc, err := parser.ParseFiles("D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\common.proto")
	if err != nil {
		panic(err)
	}

	fileDescriptors = append(fileDescriptors, fileDesc...)

	// 2
	fileDesc, err = parser.ParseFiles("D:\\development\\goProject\\study\\protoStudy\\89t\\websocket\\proto\\hamster.proto")
	if err != nil {
		panic(err)
	}

	fileDescriptors = append(fileDescriptors, fileDesc...)
}

func TestPrnit(t *testing.T) {
	b := []byte{10, 5, 106, 101, 114, 114, 121, 16, 24, 24, 0}

	fmt.Println(string(b))

}
