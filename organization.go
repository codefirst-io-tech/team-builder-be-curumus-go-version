package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Role int

const (
	ADMIN Role = iota
	USER
)

type Organization struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Members []Member
}

type Member struct {
	ID             uint `gorm:"primaryKey"`
	Name           string
	Surname        string
	Role           Role
	OrganizationID uint
}

func main() {
	dsn := "host=tai.db.elephantsql.com user=hnwhqtul password=sub7Txm66RYypQFNO5ChapHVgXPcNCXV dbname=hnwhqtul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Migrator().DropTable(&Member{}, &Organization{})

	// Migrate the schema
	db.AutoMigrate(&Member{}, &Organization{})

	db.Create(&Organization{Name: "muazzam_org"})

	var organization Organization
	db.First(&organization, "name = ?", "muazzam_org")

	member1 := Member{Role: ADMIN, Name: "member1", OrganizationID: organization.ID}
	db.Create(&member1)

	member2 := Member{Role: ADMIN, Name: "member2", OrganizationID: organization.ID}
	db.Create(&member2)

	var membersOfOrganization []Member

	db.Find(&membersOfOrganization, "organization_id = ?", organization.ID)

	fmt.Println(membersOfOrganization)
}
