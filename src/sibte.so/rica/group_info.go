package rica

import (
	"sync"
)

type GroupInfoManager interface {
	AddUser(string, string, interface{}) bool
	RemoveUser(string, string)
	GetUsers(string) []string
	GetUserInfoObject(string, string) interface{}
}

type inMemGroupInfo struct {
	sync.Mutex
	infoMap map[string]map[string]interface{}
}

func NewInMemoryGroupInfo() GroupInfoManager {
	return &inMemGroupInfo{
		infoMap: make(map[string]map[string]interface{}),
	}
}

func (me *inMemGroupInfo) AddUser(group, user string, inf interface{}) bool {
	me.Lock()
	defer me.Unlock()

	if _, ok := me.infoMap[group]; !ok {
		me.infoMap[group] = make(map[string]interface{})
	}

	me.infoMap[group][user] = inf
	return true
}

func (me *inMemGroupInfo) RemoveUser(group, user string) {
	me.Lock()
	defer me.Unlock()

	if _, ok := me.infoMap[group]; !ok {
		return
	}

	delete(me.infoMap[group], user)
}

func (me *inMemGroupInfo) GetUsers(group string) []string {
	me.Lock()
	defer me.Unlock()
	ret := make([]string, 0)
	if _, ok := me.infoMap[group]; !ok {
		return ret
	}

	ret = make([]string, len(me.infoMap[group]))
	i := 0
	for user, _ := range me.infoMap[group] {
		ret[i] = user
		i++
	}

	return ret
}

func (me *inMemGroupInfo) GetUserInfoObject(group, user string) interface{} {
	me.Lock()
	defer me.Unlock()

	if _, ok := me.infoMap[group]; !ok {
		return nil
	}

	if ret, ok := me.infoMap[group][user]; ok {
		return ret
	}

	return nil
}
