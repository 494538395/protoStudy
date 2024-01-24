package model

var WflJSON = `
{
 "identifier": 1001,
 "operation": "admin",
 "msgId": "100",
 "sendMsgReq": {
   "event": "recruit.query.pool",
   "topic": "PUSH",
   "data": {
     "id": 100
   }
 }
}
	`

var PartyBattleJSON = `
{
	"event": "hamster.change.direction",
	"id": 150,
	"data": {
		"userID": "user01",
		"newDirection": 3.14
	}
}
	`

var NestedWFLJSON = `
{
  "identifier": 1001,
  "operation": "admin",
  "msgId": "100",
  "sendMsgReq": {
    "desc": "hello",
    "info": {
      "event": "recruit.query.pool",
      "topic": "PUSH",
      "data": {
        "id": 100
      }
    }
  }
}`

var TestData = `
{
  "event": "some_event",
  "data": "SGVsbG8gdGhlcmUh",
  "topic": "PUSH"
}`
