package main

import (
	"encoding/json"
	"fmt"

	pb "protoStudy/89t/websocket/pbfile"

	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

// Req

// 接收工具页面的请求 JSON 字符串，构建发送给 GameServe 的二进制数据
func makeServerReq(jsonData string, fds []*desc.FileDescriptor) []byte {
	// 1. 将 JSON 字符串转成 Map
	var m map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &m)
	if err != nil {
		panic(err)
	}

	// 2. 将 data 部分替换成二进制
	replaceDataToBinary(m, fds)

	fmt.Println(m)

	// 构造发送给游戏服务器的 二进制数据
	protoName := "Packet"
	bytesToGameServer, err := makeBytesTOGameServer(fds, protoName, m)
	if err != nil {
		panic(err)
	}

	// 校验
	if protoName == "Packet" {
		// 模拟 Party Battle 游戏服务，验证
		var gamePkt pb.Packet

		err = proto.Unmarshal(bytesToGameServer, &gamePkt)
		if err != nil {
			panic(err)
		}
		fmt.Println("gamePkt.Id-->", gamePkt.Id)
		fmt.Println("gamePkt.Event-->", gamePkt.Event)
		// 根据 event ，解析
		if gamePkt.Event == "hamster.change.direction" {
			var hamsterDirection pb.HamsterChangeDirection
			err := proto.Unmarshal(gamePkt.Data, &hamsterDirection)
			if err != nil {
				panic(err)
			}
			fmt.Println("hamsterDirection.UserID-->", hamsterDirection.UserID)
			fmt.Println("hamsterDirection.UserID-->", hamsterDirection.NewDirection)
		}
	}

	if protoName == "Req" {
		// 模拟 WFL，验证
		var t pb.Req
		err = proto.Unmarshal(bytesToGameServer, &t)
		if err != nil {
			panic(err)
		}

		fmt.Println(t)

		fmt.Println("Req.MsgId-->", t.MsgId)
		fmt.Println("Req.Identifier-->", t.Identifier)
		fmt.Println("Req.Operation-->", t.Operation)

		sendMsg := t.GetSendMsgReq()
		fmt.Println("SendMsgReq.Desc-->", sendMsg.Event)
		fmt.Println("SendMsgReq.Info.Topic-->", sendMsg.Topic)

		if sendMsg.Event == "recruit.query.pool" {

			var rqp pb.RecruitQueryPoolReq
			err = proto.Unmarshal(sendMsg.Data, &rqp)
			if err != nil {
				panic(err)
			}
			fmt.Println("RecruitQueryPoolReq.id-->", rqp.Id)
		}

		//fmt.Println("sendMsg.Data-->", sendMsg.Data)
		//fmt.Println("sendMsg.Event-->", sendMsg.Event)
		//fmt.Println("sendMsg.Topic-->", sendMsg.Topic)
		//
		//if sendMsg.Event == "recruit.query.pool" {
		//
		//	var rqp pb.RecruitQueryPoolReq
		//	err = proto.Unmarshal(sendMsg.Data, &rqp)
		//	if err != nil {
		//		panic(err)
		//	}
		//	fmt.Println(rqp)
		//}
	}

	return bytesToGameServer
}

func makeBytesTOGameServer(fds []*desc.FileDescriptor, protoName string, jsonMap map[string]interface{}) ([]byte, error) {
	// 1.获取最外层框架 proto
	msgType := findMsgType(fds, protoName)
	msg := dynamic.NewMessage(msgType)

	// 2.将处理过的 JSON Map 序列化
	mapBytes, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}

	// 3.赋值给动态 proto
	err = json.Unmarshal(mapBytes, &msg)
	if err != nil {
		panic(err)
	}

	// 4.构造发送给 游戏服务器的 proto 二进制
	return msg.Marshal()
}

func replaceDataToBinary(m map[string]interface{}, fds []*desc.FileDescriptor) {
	// 解析 JSON 字符串，对二进制数据处理
	evt, found := m["event"]
	if found {
		// 说明在最外面扎到了 event
		event := evt.(string)
		// 2.根据 event 获取 data 的 protoType
		protoType := getProtoTypeByEventInSend(event)
		// 3.动态获取 proto 结构
		msgType := findMsgType(fds, protoType)
		msg := dynamic.NewMessage(msgType)
		// 4.拿到数据
		data := m["data"]
		dataBytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		// 5.将数据解析到动态 HamsterChangeDirection 结构中
		err = json.Unmarshal(dataBytes, &msg)
		if err != nil {
			panic(err)
		}
		// 6.得到 HamsterChangeDirection proto 的二进制数据
		msgBytes, err := msg.Marshal()
		// 7.将 data 的数据替换成二进制
		m["data"] = msgBytes
	}

	for _, v := range m {
		switch v.(type) {
		// 基本数据类型，不处理
		case []uint8, int, string, float64:
			// 第一种情况
			// message Packet {
			//   string event = 1; // 事件名
			//   bytes data = 2; // 消息体
			//   uint64 id = 3; // 消息id
			//}
		// 对象数据类型
		default:
			// 第二种情况
			//message Req {
			//  int32    identifier = 1;
			//  string   operation = 2;
			//  string   msgId = 3;
			//  oneof data {
			//    SendMsgReq sendMsgReq = 4;
			//    PullMsgBySeqListReq pullMsgBySeqListReq = 5;
			//    GetMsgMaxAndMinSeqReq maxAndMinSeqReq = 6;
			//  }
			//}

			intrMap := v.(map[string]interface{})
			_, found := intrMap["event"]
			if !found {
				replaceDataToBinary(intrMap, fds)
				return
			}

			// 1.获取 event
			event := intrMap["event"].(string)
			// 2.根据 event 获取 data 的 protoType
			protoType := getProtoTypeByEventInSend(event)
			// 3.动态获取 proto 结构
			msgType := findMsgType(fds, protoType)
			msg := dynamic.NewMessage(msgType)
			// 4.拿到数据
			data := intrMap["data"]
			dataBytes, err := json.Marshal(data)
			if err != nil {
				panic(err)
			}
			// 5.将数据解析到动态 RecruitQueryPoolReq 结构中
			err = json.Unmarshal(dataBytes, &msg)
			if err != nil {
				panic(err)
			}
			// 6.得到 RecruitQueryPoolReq proto 的二进制数据
			msgBytes, err := msg.Marshal()
			if err != nil {
				panic(err)
			}
			// 7.将 data 的数据替换成二进制
			intrMap["data"] = msgBytes
		}
	}

	return
}
