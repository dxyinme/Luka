package util

import (
	"github.com/dxyinme/LukaComm/util"
)

// ReParseBody
func ReParseBody(mp map[string]interface{}) (b []byte) {
	b, _ = util.IJson.Marshal(mp)
	return
}

func SolveError(mp map[string]interface{}, err error) bool {
	if err != nil {
		mp["error"] = err.Error()
		return true
	}
	return false
}

