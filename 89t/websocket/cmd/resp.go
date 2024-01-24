package main

import (
	"encoding/json"
	"fmt"

	pb "protoStudy/89t/websocket/proto"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

//  Resp

// 模拟 GameServer 返回的数据
func mockServerRes() []byte {
	isWfl := true

	if isWfl {
		bytes, err := proto.Marshal(&pb.RecruitQueryPoolResp{
			Id:          10001,
			CardIds:     []int32{1, 2, 3},
			LuckCardIds: []int32{4, 5, 6},
		})
		if err != nil {
			panic(err)
		}

		respBytes, err := proto.Marshal(&pb.Resp{
			Identifier: 1,
			Operation:  "admin",
			MsgId:      "1001",
			Data: &pb.Resp_PushData{PushData: &pb.GeneralMsgData{
				Event: "recruit.query.pool",
				Topic: 1,
				SeqId: 500,
				Data:  bytes,
			}},
		})
		if err != nil {
			panic(err)
		}

		var tt pb.Resp

		err = proto.Unmarshal(respBytes, &tt)
		if err != nil {
			panic(err)
		}
		fmt.Println("tt-->", tt)

		return respBytes
	}

	bytes, err := proto.Marshal(&pb.HamsterChangeDirection{
		UserID:       "user01",
		NewDirection: 3.14,
	})
	if err != nil {
		panic(err)
	}

	respBytes, err := proto.Marshal(&pb.Packet{
		Event: "hamster.change.direction",
		Data:  bytes,
		Id:    1001,
	})
	if err != nil {
		panic(err)
	}

	return respBytes
}

// 接送 GameServer 数据，构造前端返回 JSON 字符串
func makeFrontRespByServerResp(respBytes []byte, fds []*desc.FileDescriptor) string {
	msg := dynamic.NewMessage(findMsgType(fds, "Resp"))

	err := msg.Unmarshal(respBytes)
	if err != nil {
		panic(err)
	}

	marshalJSON, err := msg.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}

	err = json.Unmarshal(marshalJSON, &m)
	if err != nil {
		panic(err)
	}

	replaceBinaryToData(m, fds)

	fmt.Println(m)

	jsonBytes, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}

// jiang map 中的二进制提更换为 data
func replaceBinaryToData(m map[string]interface{}, fds []*desc.FileDescriptor) {
	evt, found := m["event"]
	if found {
		// 说明在最外面扎到了 event
		event := evt.(string)
		// 2.根据 event 获取 data 的 protoType
		protoType := getProtoTypeByEventInResp(event)
		// 3.动态获取 proto 结构
		msgType := findMsgType(fds, protoType)
		msg := dynamic.NewMessage(msgType)
		// 4.拿到数据
		data := m["data"].(string)
		tmpDataBytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		var tmpBytes []byte
		err = json.Unmarshal(tmpDataBytes, &tmpBytes)
		if err != nil {
			panic(err)
		}
		// 5.将数据解析到动态 HamsterChangeDirection 结构中
		err = msg.Unmarshal(tmpBytes)
		if err != nil {
			panic(err)
		}
		// 6.得到 HamsterChangeDirection proto 的二进制数据
		jsonBytes, err := msg.MarshalJSON()
		// 7.将 data 的数据替换成二进制
		m["data"] = string(jsonBytes)
	}

	for _, v := range m {
		switch v.(type) {
		// 基本数据类型，不处理
		case []uint8, int, string, float64:
		// 对象数据类型
		default:
			intrMap := v.(map[string]interface{})
			_, found := intrMap["event"]
			if !found {
				replaceDataToBinary(intrMap, fds)
				return
			}

			// 1.获取 event
			event := intrMap["event"].(string)
			// 2.根据 event 获取 data 的 protoType
			protoType := getProtoTypeByEventInResp(event)
			// 3.动态获取 proto 结构
			msgType := findMsgType(fds, protoType)
			msg := dynamic.NewMessage(msgType)
			// 4.拿到数据
			data := intrMap["data"].(string)

			tmpDataBytes, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}

			var tmpBytes []byte

			err = json.Unmarshal(tmpDataBytes, &tmpBytes)
			if err != nil {
				panic(err)
			}
			// 5.将数据解析到动态 RecruitQueryPoolReq 结构中
			err = msg.Unmarshal(tmpBytes)
			if err != nil {
				panic(err)
			}
			// 6.得到 RecruitQueryPoolReq proto 的二进制数据
			jsonBytes, err := msg.MarshalJSON()
			if err != nil {
				panic(err)
			}
			// 7.将 data 的数据替换成二进制
			intrMap["data"] = string(jsonBytes)

		}
	}
}
