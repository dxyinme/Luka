package util

import (
	"github.com/dxyinme/LukaComm/util"
	"io/ioutil"
	"net/http"
)


// ParseBody
func ParseBody(r *http.Request) (mp map[string]interface{}, err error) {
	mp = make(map[string]interface{})
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = util.IJson.Unmarshal(buf,&mp)
	return
}

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