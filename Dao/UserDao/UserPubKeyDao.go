package UserDao

type UserPubKeyDao interface {
	SetUserPubKey(uid string, pubKey []byte) (err error)
	GetUserPubKey(uid string) (pubKey []byte, err error)
}