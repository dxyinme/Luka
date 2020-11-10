package syncList

import (
	"container/list"
	"sync"
)

type SyncList struct {
	mutex sync.Mutex
	tmp *list.List
}

func (sl *SyncList) Lock() {
	sl.mutex.Lock()
}

func (sl *SyncList) Unlock() {
	sl.mutex.Unlock()
}

func (sl *SyncList) Back() *list.Element {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	return sl.tmp.Back()
}

func (sl *SyncList) Front() *list.Element {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	return sl.tmp.Front()
}

func (sl *SyncList) PushBack(v interface{}) *list.Element {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	return sl.tmp.PushBack(v)
}

func (sl *SyncList) Remove(element *list.Element) interface{} {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()
	return sl.tmp.Remove(element)
}

func (sl *SyncList) Len() int {
	return sl.tmp.Len()
}

// constructor

func New() *SyncList {
	return &SyncList{tmp: list.New()}
}

