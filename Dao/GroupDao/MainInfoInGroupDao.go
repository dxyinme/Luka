package GroupDao

type MainInfoInGroupDao interface {
	CreateGroup(groupName string, uid string) error
	DeleteGroup(groupName string, uid string) error
	GetAllGroup() ([]string, []string, error)
}