package GroupDao

type MainInfoInGroupDao interface {
	CreateGroup(groupName string, uid string) error
	DeleteGroup(groupName string, uid string) error
	// get all group and master-user of them.
	GetAllGroup() ([]string, []string, error)
	GroupGetAllUser(groupName string) ([]string, error)
}