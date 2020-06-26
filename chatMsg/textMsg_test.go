package chatMsg

import (
	"log"
	"testing"
	"time"
)

func TestNewTextMsg(t *testing.T) {
	o := NewTextMsg(
		"user",
		Single,
		"lalala",
		"what's up?",
	)
	log.Printf("time : %s\n", o.MsgTime)
}

func TestTime(t *testing.T) {
	log.Println(time.Now().String())
}