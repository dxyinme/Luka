package Master

import "fmt"

// 用于保存不同的KeeperChannel
type KeeperPool struct {
	mp map[string] *KeeperChannel
}

var KP *KeeperPool

// 初始化 keeperPool
func InitialKeeperPool() {
	KP = &KeeperPool{mp: make(map[string] *KeeperChannel)}
}

// 加入一个新的keeper
func updateKeeper(name string, kc *KeeperChannel) error {
	if _, ok := KP.mp[name]; !ok {
		KP.mp[name] = kc
		return nil
	}
	return fmt.Errorf("keeperName %s has been used", name)
}
