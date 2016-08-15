package rica

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
    "errors"
    "fmt"
    "math/rand"
    "regexp"
    "strings"

    "github.com/Workiva/go-datastructures/trie/ctrie"
)

var invalidAliasRegex *regexp.Regexp = nil
var cMaxNickAttempts int = 4

type NickRegistry struct {
    registryCtrie *ctrie.Ctrie
}

func NewNickRegistry() *NickRegistry {
    if invalidAliasRegex == nil {
        invalidAliasRegex, _ = regexp.Compile("[^_A-Za-z0-9]")
    }

    return &NickRegistry{
        registryCtrie: ctrie.New(nil),
    }
}

func (r *NickRegistry) GetMappingSnapshot() map[string]string {
    snapshot := r.registryCtrie.ReadOnlySnapshot()
    ret := make(map[string]string)
    for entry := range snapshot.Iterator(nil) {
        if val, ok := entry.Value.(string); ok {
            ret[string(entry.Key)] = val
        }
    }

    return ret
}

func (r *NickRegistry) SetBestPossibleNick(id, nick string) (string, error) {
    failDefault, ok := r.NickOf(id)

    if !ok {
        r.Register(id, id)
        failDefault = id
    }

    if invalidAliasRegex.MatchString(nick) || len(nick) > 42 {
        return failDefault, errors.New("A nick can only have alpha-numeric values")
    }

    i := 0
    for i = 0; i < cMaxNickAttempts && r.Register(id, nick) == false; i++ {
        nick = nick + "_"
    }

    // Try registering by appending a random number
    if i >= cMaxNickAttempts {
        nick = fmt.Sprintf("%s%d", nick, rand.Uint32())
        if !r.Register(id, nick) {
            return failDefault, errors.New("Nick already registered please choose a different nick")
        }
    }

    return nick, nil
}

func (r *NickRegistry) Register(id, nick string) bool {
    nickKey := []byte("nick:" + nick)
    idKey := []byte("id:" + id)
    if _, ok := r.registryCtrie.Lookup(nickKey); ok {
        return false
    }

    // Try setting ID for given nickKey.
    // There might be a race condition
    // Last writer will be a winner
    r.registryCtrie.Insert(nickKey, id)

    // Ensure the last writer is a winner
    if registeredIdInf, ok := r.registryCtrie.Lookup(nickKey); ok {
        if registeredId, ok := registeredIdInf.(string); !ok || strings.Compare(registeredId, id) != 0 {
            return false
        }
    } else {
        return false
    }

    r.registryCtrie.Insert(idKey, nick)
    return true
}

func (r *NickRegistry) Unregister(id string) bool {
    idKey := []byte("id:" + id)
    nick, ok := r.registryCtrie.Remove(idKey)
    if !ok {
        return false
    }

    if nickString, ok := nick.(string); ok {
        nickId := []byte("nick:" + nickString)
        if _, ok := r.registryCtrie.Remove(nickId); ok {
            return true
        }
    }

    return false
}

func (r *NickRegistry) NickOf(id string) (string, bool) {
    idKey := []byte("id:" + id)
    nick, ok := r.registryCtrie.Lookup(idKey)
    if !ok {
        return "", false
    }

    nickString, ok := nick.(string)
    return nickString, ok
}

func (r *NickRegistry) IdOf(nick string) (string, bool) {
    nickKey := []byte("nick:" + nick)
    idInf, ok := r.registryCtrie.Lookup(nickKey)
    if !ok {
        return "", false
    }

    id, ok := idInf.(string)
    return id, ok
}
