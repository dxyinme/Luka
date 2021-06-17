package lived

import (
	"github.com/dxyinme/LukaComm/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

type Controller struct {
	liveInfoMp map[string]*LiveInfo
	Lc sync.Mutex
	lid int
}

func NewController() *Controller {
	ret := &Controller{}
	ret.liveInfoMp = make(map[string]*LiveInfo)
	ret.lid = 0
	return ret
}

func (l *Controller) bootALive() {
	lidStr := strconv.Itoa(l.lid)
	l.liveInfoMp[lidStr] = NewLive(lidStr)
	l.lid ++
}

func getRoom(li *LiveInfo, roomName string) (rtmpCode string, err error) {
	c := http.Client{}
	resp, err := c.Get("http://" + li.ApiAddr + "/control/get?room=" + roomName)
	if err != nil {
		return
	}
	var respJson Response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = util.IJson.Unmarshal(body, &respJson)
	if err != nil {
		return
	}
	rtmpCode = respJson.Data.(string)
	return
}

func (l *Controller) BootLive(w http.ResponseWriter, r *http.Request) {
	res := Response{
		w:      w,
		Status: 200,
		Data:   nil,
	}
	if err := r.ParseForm(); err != nil {
		res.Status = 400
		return
	}
	l.Lc.Lock()
	defer l.Lc.Unlock()
	defer res.SendJson()
	l.bootALive()
	res.Data = l.lid - 1
}

func (l *Controller) NewRoom(w http.ResponseWriter, r *http.Request) {
	res := Response{
		w:      w,
		Status: 200,
		Data:   nil,
	}
	if err := r.ParseForm(); err != nil {
		res.Status = 400
		return
	}
	l.Lc.Lock()
	defer l.Lc.Unlock()
	defer res.SendJson()

	type data struct {
		RtmpCode string	`json:"rtmpCode"`
		RtmpUrl string `json:"rtmpUrl"`
		FlvUrl string `json:"flvUrl"`
		HlsUrl string `json:"hlsUrl"`
	}

	var resData data
	for _,v := range l.liveInfoMp {
		rtmpCode, err := getRoom(v, r.Form.Get("room"))
		if err != nil {
			res.Status = 500
			res.Data = err.Error()
			return
		}
		resData.RtmpCode = rtmpCode
		resData.RtmpUrl = "rtmp://" + v.RtmpAddr + "/live"
		resData.FlvUrl = "http://" + v.FlvAddr + "/live/" + r.Form.Get("room") + ".flv"
		resData.HlsUrl = "http://" + v.FlvAddr + "/live/" + r.Form.Get("room") + ".m3u8"
		res.Data = resData
		break
	}
	return
}
