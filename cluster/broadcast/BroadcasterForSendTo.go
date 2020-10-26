package broadcast

type BroadcasterForSendTo struct {
	finished chan bool
}

func (bFst *BroadcasterForSendTo) Initial() {
	bFst.finished = make(chan bool, 1)
}

func (bFst *BroadcasterForSendTo) Do() {

}

