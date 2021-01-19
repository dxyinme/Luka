package UserDao


// uid : [group1, group2, group3, ... ]
type UserGroupDao interface {
	JoinGroupDao(uid string, groupName string) error
	LeaveGroupDao(uid string, groupName string) error
	CreateGroupDao(uid string, groupName string) error
	DeleteGroupDao(uid string, groupName string) error
	GetGroupNameList(uid string) ([]string, error)
}