package WorkerPool

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"testing"
)

func TestCopy(t *testing.T) {
	T1 := &chatMsg.Msg{
		From:           "",
		Target:         "",
		Content:        nil,
		MsgType:        0,
		MsgContentType: 0,
		SendTime:       "",
		GroupName:      "",
		Spread:         false,
	}
	T2 := *T1
	log.Printf("%p\n", T1)
	log.Printf("%p\n", &T2)
	T2.Spread = true
	if T1.Spread {
		t.Fatal("error")
	}
}


// be used in local-test
//func TestNormalImpl_SyncGroupInfo(t *testing.T) {
//	type respType struct {
//		GroupNameList []string
//		UidList []string
//	}
//	var (
//		body []byte
//		respItem respType
//	)
//	GetUrl := "http://" + "localhost:12777" + "/group/GroupInfo/GetAllGroup"
//	c := http.Client{}
//	resp, err := c.PostForm(GetUrl, url.Values{})
//	if err != nil {
//		log.Fatal(err)
//	}
//	body, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = util.IJson.Unmarshal(body, &respItem)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(respItem)
//}