package member

import (
	"golang/model"
)

type Member struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Surname        string
	Role           Role
	OrganizationID uint
	Strength       float32
	Position       Position
	Avatar         string
}

type Role int
type Position int

const (
	ADMIN Role = iota
	USER
)
const (
	GOALKEEPER Position = iota
	DEFENDER
	MIDFIELDER
	FORWARD
)

func CreateDropTable() {
	db := model.GetConnection()
	db.Migrator().DropTable(&Member{})
	db.AutoMigrate(&Member{})
}

func Save(member Member) {
	model.GetConnection().Create(&member)
}

func FindById(id uint) Member {
	var member Member
	model.GetConnection().Find(&member, "id = ?", id)
	if &member == nil {
		panic("there is not member by id : " + string(id))
	}
	return member
}

func FindByName(name string) Member {
	var member Member
	model.GetConnection().Find(&member, "name = ?", name)
	if &member == nil {
		panic("there is not member by name : " + name)
	}
	return member
}
