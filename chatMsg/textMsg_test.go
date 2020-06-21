package chatMsg

import (
	"github.com/golang/glog"
	"testing"
)

func TestNewTextMsg(t *testing.T) {
	o := NewTextMsg(
		"user",
		Single,
		"lalala",
		"what's up?",
		)
	glog.Info("time : %d\n",o.MsgTime)
}