package organization

import (
	"golang/model"
	"golang/model/member"
)

type Organization struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Members []member.Member
}

func CreateDropTable() {
	db := model.GetConnection()
	db.Migrator().DropTable(&Organization{})
	db.AutoMigrate(&Organization{})
}

func Save(organization Organization) {
	model.GetConnection().Create(&organization)
}

func FindById(id uint) Organization {
	var organization Organization
	model.GetConnection().Find(&organization, "id = ?", id)
	if &organization == nil {
		panic("there is not organization by id : " + string(id))
	}
	return organization
}

func FindByName(name string) Organization {
	var organization Organization
	model.GetConnection().Find(&organization, "name = ?", name)
	if &organization == nil {
		panic("there is not organization by name : " + name)
	}
	return organization
}
