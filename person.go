package dbmock

type Person struct {
	ID   int    `xorm:"pk id"`
	Name string `xorm:"name"`
}

func (p *Person) TableName() string {
	return "person"
}
