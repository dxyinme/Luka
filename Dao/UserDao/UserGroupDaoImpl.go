package UserDao


type UGDImpl struct {
	
}

func (U *UGDImpl) JoinGroupDao(uid string, groupName string) error {
	panic("implement me")
}

func (U *UGDImpl) LeaveGroupDao(uid string, groupName string) error {
	panic("implement me")
}

func (U *UGDImpl) CreateGroupDao(uid string, groupName string) error {
	panic("implement me")
}

func (U *UGDImpl) DeleteGroupDao(uid string, groupName string) error {
	panic("implement me")
}

func (U *UGDImpl) GetGroupNameList(uid string) ([]string, error) {
	panic("implement me")
}
