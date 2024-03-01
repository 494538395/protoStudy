package main

import (
	"fmt"

	"protoStudy/89t/websocket/model"
)

func main() {

	makeServerReq(model.PartyBattleJSON, fds)
	//
	fmt.Println(makeFrontRespByServerResp(mockServerRes(), fds))
}
