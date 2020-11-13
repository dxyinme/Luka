package ListCache

import (
	"github.com/dxyinme/Luka/util/syncList"
	"sync"
)

type ListCache struct {
	mu sync.Mutex
	mp map[string]*syncList.SyncList
}

func (lc *ListCache) Get(key string) (*syncList.SyncList, bool) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	v, ok := lc.mp[key]
	return v, ok
}

func (lc *ListCache) Set(key string, v *syncList.SyncList) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.mp[key] = v
}

func (lc *ListCache) Delete(key string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	delete(lc.mp, key)
}

//constructor
func New() *ListCache {
	return &ListCache{
		mp: make(map[string]*syncList.SyncList),
	}
}
