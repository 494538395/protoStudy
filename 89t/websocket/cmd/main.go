package main

import (
	"fmt"

	"github.com/jhump/protoreflect/desc"
)

var fds []*desc.FileDescriptor

func main() {

	//makeServerReq(model.PartyBattleJSON, fds)

	fmt.Println(makeFrontRespByServerResp(mockServerRes(), fds))
}
