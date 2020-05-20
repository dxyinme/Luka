package Provider

import (
	"Luka/util"
	"net/http"
)
// 心跳机制
func ConfirmReply(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte(util.OK))
}