package config

import (
	"bufio"
	"flag"
	"github.com/golang/glog"
	"io"
	"os"
	"strconv"
	"strings"
)

var (

	// ClusterFile : the ipports for each server in cluster
	ClusterFile = flag.String("ClusterFile", "", "the file of ClusterInfo")

	// hostAddrPtr
	hostAddrPtr = flag.String("HostAddr", "", "the addr of this keeper")
	// keeperIDPtr
	keeperIdPtr = flag.Int("KeeperID", 0, "the keeperID of this keeper")

	Host     string
	HostAddr string
	AllHosts []string
	KeeperID int
)

func GetIP() string {
	return strings.Split(Host, ":")[0]
}

func SetHost(host string) {
	Host = strings.TrimSpace(host)
	HostAddr = ":" + strings.TrimSpace(strings.Split(host, ":")[1])
	glog.Infof("Host=[%s], Port=[%s]", Host, HostAddr)
}

func SetAllHosts(hosts []string) {
	AllHosts = hosts
}

func LoadFromFile() {
	var filename = *ClusterFile
	file, err := os.Open(filename)
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
		nowLineInfo := strings.Split(line, " ")
		if len(nowLineInfo) < 2 {
			continue
		} else {
			// [Host] length is 3 , [AllHostX] length is 2
			if nowLineInfo[0] == "[Host]" {
				SetHost(nowLineInfo[1])
				KeeperID, err = strconv.Atoi(strings.TrimSpace(nowLineInfo[2]))
				if err != nil {
					glog.Fatal(err)
				}
			} else {
				AllHosts = append(AllHosts, strings.TrimSpace(nowLineInfo[1]))
			}
		}
	}
}

func LoadFromCmd() {
	SetHost(*hostAddrPtr)
	KeeperID = *keeperIdPtr
}