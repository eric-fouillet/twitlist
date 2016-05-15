package server

import (
	"fmt"
	"time"

	"github.com/eric-fouillet/anaconda"
)

type DummyTwitterClient struct {
	api            *anaconda.TwitterApi
	lists          []anaconda.List
	lastUpdateTime time.Time
	authenticated  bool
}

func (tc *DummyTwitterClient) authenticate() error {
	return nil
}

func (tc *DummyTwitterClient) close() {
}

func (tc *DummyTwitterClient) getSelfId() (int64, error) {
	return 1, nil
}

func (tc *DummyTwitterClient) GetListMembers(id int64) ([]anaconda.User, error) {
	users := make([]anaconda.User, 0, 10)
	var add func()
	var idu int64 = 1
	add = func() {
		users = append(users, anaconda.User{Id: idu, Name: fmt.Sprint("user", id)})
		idu++
	}
	for i := 0; i < 10; i++ {
		add()
	}
	return users, nil
}

func (tc *DummyTwitterClient) GetAllLists() ([]anaconda.List, error) {
	lists := make([]anaconda.List, 0, 5)
	var add func()
	var id int64 = 1
	add = func() {
		lists = append(lists, anaconda.List{Id: id, Name: fmt.Sprint("list", id)})
		id++
	}
	for i := 0; i < 5; i++ {
		add()
	}
	tc.lists = lists
	return tc.lists, nil
}
