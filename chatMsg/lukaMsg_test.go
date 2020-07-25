package chatMsg

import (
	"github.com/dxyinme/Luka/util"
	"testing"
)

func TestLukaMsg_GetContent(t *testing.T) {
	tests := []*LukaMsg{
		NewLukaMsg("1", "2", Single, Text, []byte("haha"), false),
	}
	wants := []string{
		"haha",
	}
	for i := 0; i < len(tests); i++ {
		if string(util.B64Encode([]byte(wants[i]))) != tests[i].GetContent() {
			t.Errorf("test case %d, wants(%s) , but(%s)", i, wants[i], tests[i].GetContent())
		}
	}
}

func TestLukaMsg_GetFrom(t *testing.T) {
	tests := []*LukaMsg{
		NewLukaMsg("1", "2", Single, Text, []byte("haha"),false),
	}
	wants := []string{
		"1",
	}
	for i := 0; i < len(tests); i++ {
		if wants[i] != tests[i].GetFrom() {
			t.Errorf("test case %d, wants(%s) , but(%s)", i, wants[i], tests[i].GetFrom())
		}
	}
}

func TestLukaMsg_GetMsgContentType(t *testing.T) {
	tests := []*LukaMsg{
		NewLukaMsg("1", "2", Single, Text, []byte("haha"), false),
	}
	wants := []MsgContentTypeEnum{
		Text,
	}
	for i := 0; i < len(tests); i++ {
		if wants[i] != tests[i].GetMsgContentType() {
			t.Errorf("test case %d, wants(%d) , but(%d)", i, wants[i], tests[i].GetMsgContentType())
		}
	}
}

func TestLukaMsg_GetMsgType(t *testing.T) {
	tests := []*LukaMsg{
		NewLukaMsg("1", "2", Single, Text, []byte("haha"),false),
	}
	wants := []MsgTypeEnum{
		Single,
	}
	for i := 0; i < len(tests); i++ {
		if wants[i] != tests[i].GetMsgType() {
			t.Errorf("test case %d, wants(%d) , but(%d)", i, wants[i], tests[i].GetMsgType())
		}
	}
}

func TestLukaMsg_GetTarget(t *testing.T) {
	tests := []*LukaMsg{
		NewLukaMsg("1", "2", Single, Text, []byte("haha"),false),
	}
	wants := []string{
		"2",
	}
	for i := 0; i < len(tests); i++ {
		if wants[i] != tests[i].GetTarget() {
			t.Errorf("test case %d, wants(%s) , but(%s)", i, wants[i], tests[i].GetTarget())
		}
	}
}

func TestLukaMsg_GetTime(t *testing.T) {
	// useless
}
