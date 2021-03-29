package lived

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

var (
	alivePort chan int
	defaultPortRange = "50010,50110"
	portRange = flag.String("pr", defaultPortRange , "PortRange")
	LcUtil sync.Mutex
	ip = flag.String("ip", "127.0.0.1", "ip address")
	bashTemplate = "nohup ./livego --api_addr=:%d --httpflv_addr=:%d --rtmp_addr=:%d --hls_addr=:%d --flv_dir=./%s/tmp > %s.nohup 2>&1 &"
)

func Init() {
	var (
		l uint
		r uint
		err error
		Sz int
	)
	_, err = fmt.Sscanf(*portRange, "%d,%d", &l, &r)
	if err != nil {
		_, err = fmt.Sscanf(defaultPortRange,"%d,%d", &l, &r)
	}
	if l > r {
		t := l
		l = r
		r = t
	}
	Sz = int(r - l + 1)
	alivePort = make(chan int, Sz + 1)
	for i := int(l); i <= int(r); i ++ {
		alivePort <- i
	}
}

func Bash(cmdline string) {
	cmdObj := exec.Command("bash", "-c", cmdline)
	output, err := cmdObj.CombinedOutput()
	if err != nil {
		if ins, ok := err.(*exec.ExitError); ok {
			out := string(output)
			exitcode := ins.ExitCode()
			glog.Warningf("bash cmd:[%s], out:[%s], exitcode:[%d]", cmdline, out, exitcode)
			return
		}
	}
	return
}

func NewLive(name string) *LiveInfo {
	LcUtil.Lock()
	defer LcUtil.Unlock()
	if len(alivePort) < 4 {
		return nil
	}
	apiPort := <-alivePort
	rtmpPort := <-alivePort
	hlsPort := <-alivePort
	flvPort := <-alivePort
	nowBash := fmt.Sprintf(bashTemplate, apiPort, flvPort, rtmpPort, hlsPort, name, name)
	err := os.MkdirAll(name + "/tmp", 0777)
	if err != nil {
		glog.Error(err)
		return nil
	}
	f, err := os.OpenFile("tmp.script", os.O_RDWR | os.O_CREATE , 0777)
	if err != nil {
		glog.Error(err)
		return nil
	}
	_,err = f.Write([]byte(nowBash))
	if err != nil {
		glog.Error(err)
	}
	f.Close()
	Bash("bash tmp.script")
	return &LiveInfo{
		ApiAddr: *ip + ":" + strconv.Itoa(apiPort),
		RtmpAddr: *ip + ":" + strconv.Itoa(rtmpPort),
		HlsAddr: *ip + ":" + strconv.Itoa(hlsPort),
		FlvAddr: *ip + ":" + strconv.Itoa(flvPort),
		Name: name,
	}
}