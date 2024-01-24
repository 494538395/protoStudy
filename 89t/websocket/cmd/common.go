package main

import (
	"fmt"

	"github.com/jhump/protoreflect/desc"
)

func getProtoTypeByEventInSend(event string) string {
	switch event {
	case "recruit.query.pool":
		return "RecruitQueryPoolReq"
	case "hamster.change.direction":
		return "HamsterChangeDirection"
	}

	return "unknown!!"
}

func getProtoTypeByEventInResp(event string) string {
	switch event {
	case "recruit.query.pool":
		return "RecruitQueryPoolResp"
	case "hamster.change.direction":
		return "HamsterChangeDirection"
	}

	return "unknown!!"
}

// 根据 proto message Name 找 message 定义
func findMsgType(fds []*desc.FileDescriptor, target string) *desc.MessageDescriptor {

	for _, fd := range fds {
		r := fd.FindMessage(target)
		fmt.Println("r-->", r)

		for _, msgTypeDescriptor := range fd.GetMessageTypes() {
			if msgTypeDescriptor.GetName() == target {
				return msgTypeDescriptor
			}
		}
	}

	return nil
}
