package dbmock

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

type Repository interface {
	Get(id int) (*Person, error)
	Create(id int, name string) error
	Update(id int, name string) error
	Delete(id int) error
}

type repo struct {
	session *xorm.Session
}

func NewPersonRepo(session *xorm.Session) Repository {
	return repo{session}
}

func (r repo) Get(id int) (person *Person, err error) {
	person = &Person{ID: id}
	has, err := r.session.Get(person)
	if err != nil {
		return
	}
	if !has {
		err = fmt.Errorf("person[id=%d] not found", id)
		return
	}

	return
}

func (r repo) Create(id int, name string) (err error) {
	person := &Person{ID: id, Name: name}
	affected, err := r.session.Insert(person)
	if err != nil {
		return
	}

	if affected == 0 {
		err = fmt.Errorf("insert err, because of 0 affected")
		return
	}

	return
}

func (r repo) Update(id int, name string) (err error) {
	_, err = r.session.ID(id).Cols("name").Update(&Person{Name: name})
	return
}

func (r repo) Delete(id int) (err error) {
	affected, err := r.session.ID(id).Delete(&Person{})
	if err != nil {
		return
	}

	if affected == 0 {
		err = fmt.Errorf("delete err, because of 0 affected")
		return
	}

	return
}
