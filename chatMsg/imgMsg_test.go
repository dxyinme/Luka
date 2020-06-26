package chatMsg

import "testing"

func TestNewImgMsg(t *testing.T) {
	var o Msg
	o = NewImgMsg("1","2",Single,Img,[]byte("111111111111111111112222222"))
	if o.GetFrom() != "1" {
		t.Fatal("err [From]")
	}
}
