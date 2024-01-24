package main

import (
	"encoding/json"
	"fmt"
	"testing"

	pb "protoStudy/89t/websocket/proto"

	"github.com/jhump/protoreflect/dynamic"
	"google.golang.org/protobuf/proto"
)

func TestMarshall(t *testing.T) {
	pbBytes, err := proto.Marshal(&pb.RecruitQueryPoolResp{
		Id:          500,
		CardIds:     []int32{1, 2, 3},
		LuckCardIds: []int32{4, 5, 6},
	})
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(pbBytes)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonBytes))

	var tmp []byte

	err = json.Unmarshal(jsonBytes, &tmp)
	if err != nil {
		panic(err)
	}

	var tt pb.RecruitQueryPoolResp

	err = proto.Unmarshal(tmp, &tt)
	if err != nil {
		panic(err)
	}

	fmt.Println(tt)

	err = json.Unmarshal(jsonBytes, &tt)
	if err != nil {
		panic(err)
	}

	fmt.Println(tt)
}

func TestMultiNestedOneof(t *testing.T) {

	msg := &pb.Resp{
		Identifier: 1001,
		Operation:  "admin",
		MsgId:      "200",
		Data: &pb.Resp_PullMsgBySeqListResp{PullMsgBySeqListResp: &pb.PullMsgBySeqListResp{
			Name: "jerry",
			Age:  21,
			Data: &pb.PullMsgBySeqListResp_Pec{Pec: &pb.PersonC{
				Name:    "sam",
				Address: "beijing",
			}},
		},
		},
	}

	marshal, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(marshal)

	dynaMsg := dynamic.NewMessage(findMsgType(fds, "Resp"))

	err = dynaMsg.Unmarshal(marshal)
	if err != nil {
		panic(err)
	}

	marshalJSON, err := dynaMsg.MarshalJSON()
	if err != nil {
		panic(err)
	}

	fmt.Println("marshalJSON-->", marshalJSON)

}
