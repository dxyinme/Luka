package FileServer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dxyinme/LukaComm/util/MD5"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	chunkSize = 1<<16
	templateUpload = "http://%s/api/uploadSlice?MD5=%s&from=%s&end=%s"
)

type FileClient struct {
	Host string
}


func (fc *FileClient) SendFile(filepath string, from string) error {
	fp, err := os.Open(filepath)
	if err != nil {
		return err
	}
	md5, err := MD5.CalcMD5File(filepath)
	if err != nil {
		return err
	}
	for ;; {
		b := make([]byte, chunkSize)
		n, err := fp.Read(b)
		if err != nil {
			return err
		}
		end := "0"
		if n < chunkSize {
			end = "1"
		}
		uploadUrl := fmt.Sprintf(templateUpload, fc.Host, md5, from, end)

		rsp, err := http.Post(uploadUrl, "multipart/form-data", bytes.NewReader(b[:n]))
		if err != nil {
			return err
		}
		rspStr,err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return err
		}
		if string(rspStr) != "OK" {
			return errors.New(string(rspStr))
		}
		if end == "1" {
			break
		}
	}
	return nil
}