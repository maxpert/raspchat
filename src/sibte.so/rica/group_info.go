package rica

import (
    "fmt"

    "github.com/Workiva/go-datastructures/trie/ctrie"
)

type GroupInfoManager interface {
    AddUser(string, string, interface{}) bool
    RemoveUser(string, string)
    GetUsers(string) []string
    GetUserInfoObject(string, string) interface{}
    GetAllInfoObjects(string) map[string]interface{}
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
        panic(fmt.Sprintln("Unable to add user", user, "from", group))
    }

    usersCtrie.Insert([]byte(user), inf)
    return true
}

func (i *inMemGroupInfo) RemoveUser(group, user string) {
    usersCtrie, ok := i.createOrGetGroupMap(group)
    if !ok {
        panic(fmt.Sprintln("Unable to remove user", user, "from", group))
    }

    userKey := []byte(user)
    for {
        usersCtrie.Remove(userKey)
        if _, ok := usersCtrie.Lookup(userKey); !ok {
            break
        }
    }
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

func (i *inMemGroupInfo) GetAllInfoObjects(group string) map[string]interface{} {
    if usersCtrie, ok := i.createOrGetGroupMap(group); ok {
        snapshot := usersCtrie.Snapshot()
        ret := make(map[string]interface{})
        for u := range snapshot.Iterator(nil) {
            ret[string(u.Key)] = u.Value
        }

        return ret
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
