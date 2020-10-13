package config

import (
	"bufio"
	"github.com/golang/glog"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	Host     string
	HostAddr string
	AllHosts []string
	KeeperID int
)

func SetHost(host string) {
	Host = strings.TrimSpace(host)
	HostAddr = ":" + strings.TrimSpace(strings.Split(host, ":")[1])
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
		if len(nowLineInfo) < 2 {
			continue
		} else {
			// [Host] length is 3 , [AllHostX] length is 2
			if nowLineInfo[0] == "[Host]" {
				SetHost(nowLineInfo[1])
				KeeperID,err = strconv.Atoi(strings.TrimSpace(nowLineInfo[2]))
				if err != nil {
					glog.Fatal(err)
				}
			} else {
				AllHosts = append(AllHosts, strings.TrimSpace(nowLineInfo[1]))
			}
		}
	}
}