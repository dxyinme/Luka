package chatMsg

import (
	"log"
	"testing"
)

func TestNewTextMsg(t *testing.T) {
	o := NewTextMsg(
		"user",
		Single,
		"lalala",
		"what's up?",
		)
	log.Printf("time : %d\n",o.MsgTime)
}