package rica

/*
Copyright (c) 2015 Zohaib
Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"errors"
	"log"
	"regexp"
	"sync"
)

var invalidAliasRegex *regexp.Regexp = nil

type NickRegistry struct {
	sync.Mutex
	registry       map[string]string
	uniqueAliasMap map[string]string
}

func NewNickRegistry() *NickRegistry {
	if invalidAliasRegex == nil {
		invalidAliasRegex, _ = regexp.Compile("[^\\._A-Za-z0-9]")
	}

	return &NickRegistry{
		registry:       make(map[string]string),
		uniqueAliasMap: make(map[string]string),
	}
}

func (r *NickRegistry) SetNick(id, nick string) (string, error) {
	failDefault, ok := r.registry[id]

	if !ok {
		r.Register(id, id)
		failDefault = id
	}

	if invalidAliasRegex.MatchString(nick) || len(nick) > 42 {
		return failDefault, errors.New("A nick can only have alpha-numeric values")
	}

	i := 0
	for i = 0; i < 3 && r.Register(id, nick) == false; i++ {
		log.Println("Nick", nick, "already registered retry", i)
		nick = nick + "_"
	}

	if i >= 3 {
		return failDefault, errors.New("Nick already registered please choose a different nick")
	}

	return nick, nil
}

func (r *NickRegistry) Register(id, nick string) bool {
	r.Lock()
	defer r.Unlock()

	oldId, ok := r.uniqueAliasMap[nick]
	if ok && oldId != id {
		return false
	}

	if ok {
		delete(r.uniqueAliasMap, nick)
	}

	r.registry[id] = nick
	r.uniqueAliasMap[nick] = id
	return true
}

func (r *NickRegistry) Unregister(id string) bool {
	r.Lock()
	defer r.Unlock()

	nick, ok := r.registry[id]
	if !ok {
		log.Println("Unable to remove nick name registry", id)
		return false
	}

	delete(r.registry, id)
	delete(r.uniqueAliasMap, nick)
	return true
}

func (r *NickRegistry) NickOf(id string) (string, bool) {
	r.Lock()
	defer r.Unlock()
	nick, ok := r.registry[id]
	return nick, ok
}
