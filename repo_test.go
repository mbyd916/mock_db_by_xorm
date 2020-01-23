package dbmock

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPersonGet(t *testing.T) {
	Convey("Setup", t, func() {
		session, mock := getSession()
		repo := NewPersonRepo(session)
		id, name := 1, "John"
		Convey("get some person by id", func() {
			mock.ExpectQuery("SELECT (.+) FROM `person`").
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name))
			person, err := repo.Get(id)
			So(err, ShouldBeNil)
			So(person, ShouldResemble, &Person{ID: id, Name: name})
		})

		Reset(func() {
			So(mock.ExpectationsWereMet(), ShouldBeNil)
		})
	})

}

func TestPersonCreate(t *testing.T) {
	Convey("Setup", t, func() {
		session, mock := getSession()
		repo := NewPersonRepo(session)
		id, name := 1, "John"
		Convey("create a person", func() {
			mock.ExpectExec("INSERT INTO `person`").
				WithArgs(id, name).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Create(id, name)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			So(mock.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}

func TestPersonUpdate(t *testing.T) {
	Convey("Setup", t, func() {
		session, mock := getSession()
		repo := NewPersonRepo(session)
		id, name := 1, "John"
		Convey("update a person", func() {
			mock.ExpectExec("UPDATE `person`").
				WithArgs(name, id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Update(id, name)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			So(mock.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}

func TestPersonDelete(t *testing.T) {
	Convey("Setup", t, func() {
		session, mock := getSession()
		repo := NewPersonRepo(session)
		id := 1
		Convey("delete a person", func() {
			mock.ExpectExec("DELETE FROM `person`").
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.Delete(id)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			So(mock.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}

func getSession() (*xorm.Session, sqlmock.Sqlmock) {

	db, mock, err := sqlmock.New()
	So(err, ShouldBeNil)

	eng, err := xorm.NewEngine("mysql", "root:123@/test?charset=utf8")
	So(err, ShouldBeNil)

	eng.DB().DB = db
	eng.ShowSQL(true)

	return eng.NewSession(), mock
}
