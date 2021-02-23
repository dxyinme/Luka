package FileServer

import (
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	fileMp map[string]*fInfo
)

func Initial(router *mux.Router) {
	fileMp = make(map[string]*fInfo)
	router.HandleFunc("/uploadSlice", uploadSlice)
	router.HandleFunc("/downloadSlice", downloadSlice)
}

func returnErr(errStr string, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_,_ = w.Write([]byte(errStr))
}

// /upload?MD5=[MD5CODE]&from=[from]&end=[0 for no yet , 1 for end]
func uploadSlice(w http.ResponseWriter, r *http.Request) {
	var (
		md5 string
		from string
		end string
	)
	err := r.ParseForm()
	if err != nil {
		glog.Error(err)
		return
	}
	md5 = r.Form.Get("MD5")
	if md5 == "" {
		glog.Error("defect param `MD5`")
		returnErr("defect param `MD5`", w)
		return
	}
	from = r.Form.Get("from")
	if from == "" {
		glog.Error("defect param `from`")
		returnErr("defect param `from`", w)
		return
	}
	end = r.Form.Get("end")
	if end == "" {
		glog.Error("defect param `end`")
		returnErr("defect param `end`", w)
		return
	}

	b , err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Error(err)
		returnErr(err.Error(), w)
		return
	}

	f, ok := fileMp[from + "/" + md5]
	if !ok {
		f = NewFInfo(from , md5)
		fileMp[from + "/" + md5] = f
	}
	err = f.Append(b)
	if err != nil {
		glog.Error(err)
		returnErr(err.Error(), w)
		return
	}
	if end == "1" {
		err = f.Locked()
		delete(fileMp, from + "/" + md5)
		if err != nil {
			glog.Error(err)
			returnErr(err.Error(), w)
			return
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	_,_ = w.Write([]byte("OK"))
}

// /downloadSlice?MD5=[MD5CODE]&from=[from]
func downloadSlice(w http.ResponseWriter, r *http.Request) {
	var (
		md5 string
		from string
	)
	err := r.ParseForm()
	if err != nil {
		glog.Error(err)
		return
	}
	md5 = r.Form.Get("MD5")
	if md5 == "" {
		glog.Error("defect param `MD5`")
		returnErr("defect param `MD5`", w)
		return
	}
	from = r.Form.Get("from")
	if from == "" {
		glog.Error("defect param `from`")
		returnErr("defect param `from`", w)
		return
	}
	_, ok := fileMp[from + "/" + md5]
	if ok {
		returnErr("not yet for download", w)
		return
	}
	fileDist, err := os.Open(from + "/" + md5)
	if err != nil {
		returnErr(err.Error(), w)
		return
	}
	defer fileDist.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err = io.Copy(w, fileDist); err != nil {

	}
}