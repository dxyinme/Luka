package config

import (
	"bufio"
	"github.com/golang/glog"
	"io"
	"os"
	"strings"
)

var (
	Host     string
	HostAddr string
	AllHosts []string
)

func SetHost(host string) {
	Host = strings.TrimSpace(host)
	Host = ":" + strings.TrimSpace(strings.Split(host, ":")[1])
}

func SetAllHosts(hosts []string) {
	AllHosts = hosts
}

func LoadFromFile(filename string) {
	file,err := os.Open(filename)
	if err != nil {
		glog.Error(err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		lineByte, _, ok := buf.ReadLine()
		line := string(lineByte)
		if ok == io.EOF {
			break
		}
		nowLineInfo := strings.Split(line , " ")
		if len(nowLineInfo) != 2 {
			continue
		} else {
			if nowLineInfo[0] == "[Host]" {
				SetHost(nowLineInfo[1])
			} else {
				AllHosts = append(AllHosts, strings.TrimSpace(nowLineInfo[1]))
			}
		}
	}
}