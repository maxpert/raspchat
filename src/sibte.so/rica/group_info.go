package rica

import (
    "github.com/Workiva/go-datastructures/trie/ctrie"
)

type GroupInfoManager interface {
    AddUser(string, string, interface{}) bool
    RemoveUser(string, string)
    GetUsers(string) []string
    GetUserInfoObject(string, string) interface{}
}

type inMemGroupInfo struct {
    channelsCtrie *ctrie.Ctrie
}

func NewInMemoryGroupInfo() GroupInfoManager {
    return &inMemGroupInfo{
        channelsCtrie: ctrie.New(nil),
    }
}

func (i *inMemGroupInfo) AddUser(group, user string, inf interface{}) bool {
    usersCtrie, ok := i.createOrGetGroupMap(group)
    if !ok {
        return false
    }

    usersCtrie.Insert([]byte(user), inf)
    return true
}

func (i *inMemGroupInfo) RemoveUser(group, user string) {
    usersCtrie, ok := i.createOrGetGroupMap(group)
    if !ok {
        return
    }

    usersCtrie.Remove([]byte(user))
}

func (i *inMemGroupInfo) GetUsers(group string) []string {
    usersCtrie, ok := i.createOrGetGroupMap(group)
    if !ok {
        return make([]string, 0)
    }

    snapShotCtrie := usersCtrie.ReadOnlySnapshot()
    ret := make([]string, snapShotCtrie.Size())
    j := 0
    for entry := range snapShotCtrie.Iterator(nil) {
        ret[j] = string(entry.Key)
        j++
    }

    return ret
}

func (i *inMemGroupInfo) GetUserInfoObject(group, user string) interface{} {
    if usersCtrie, ok := i.createOrGetGroupMap(group); ok {
        if userObj, ok := usersCtrie.Lookup([]byte(user)); ok {
            return userObj
        }
    }

    return nil
}

func (i *inMemGroupInfo) createOrGetGroupMap(group string) (*ctrie.Ctrie, bool) {
    // If found on first shot we are good to go, no race conditions
    if groupObj, ok := i.channelsCtrie.Lookup([]byte(group)); ok {
        usersCtrie, ok := groupObj.(*ctrie.Ctrie)
        return usersCtrie, ok
    }

    // Insert new entry, there might be a race condition
    // should be resolved when doing Lookup, there would be only one winner
    i.channelsCtrie.Insert([]byte(group), ctrie.New(nil))
    if groupObj, ok := i.channelsCtrie.Lookup([]byte(group)); ok {
        usersCtrie, ok := groupObj.(*ctrie.Ctrie)
        return usersCtrie, ok
    }

    return nil, false
}
