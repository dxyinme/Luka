package Master

// 用于传递不同的keeper之间的信息，这是一个位于master-server中的类
//
// todo
type KeeperChannel struct {
	name 	string
	url 	string
}

// constructor for keeperChannel
func NewKeeperChannel(nameNow,urlNow string) *KeeperChannel{
	return &KeeperChannel{
		name: nameNow,
		url:  urlNow,
	}
}