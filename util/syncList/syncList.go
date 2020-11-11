package syncList

import (
	"container/list"
	"sync"
)

// only be used in normal type or pointer[do not use SearchAndRemove in complex struct]
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

// CAN NOT BE USE IN List<contain> or List<map> or List<func> or List<Slice>
// search all the interface equal to v and delete them
func (sl *SyncList) SearchAndRemove(v interface{}) {
	sl.mutex.Lock()
	for item := sl.tmp.Front() ; item != nil ; item = item.Next() {
		if v == item.Value {
			sl.tmp.Remove(item)
		}
	}
	sl.mutex.Unlock()
}

func (sl *SyncList) Len() int {
	return sl.tmp.Len()
}

// constructor

func New() *SyncList {
	return &SyncList{tmp: list.New()}
}

