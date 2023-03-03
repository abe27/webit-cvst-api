package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
	UserName  string    `gorm:"not null;column:username;index;unique;size:10" json:"username" form:"username"`
	Email     string    `gorm:"not null;unique;size:50;" json:"email" form:"email"`
	Password  string    `gorm:"not null;unique;size:60;" json:"-" form:"password"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Administrator struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id"`
	UserID    *string   `gorm:"unique;" json:"user_id" form:"user_id"`
	IsActive  bool      `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}

func (obj *Administrator) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Profile struct {
	ID           string     `gorm:"primaryKey;size:21;" json:"id"`
	AvatarURL    string     `json:"avatar_url" form:"avatar_url"`
	PrefixNameID *string    `json:"prefix_name_id" form:"prefix_name_id"`
	FirstName    string     `gorm:"not null;size:50;" json:"first_name" form:"first_name"`
	LastName     string     `gorm:"not null;size:50;" json:"last_name" form:"last_name"`
	PhoneNo      string     `json:"phone_no" form:"phone_no"`
	UserID       *string    `gorm:"not null;unique;" json:"user_id" form:"user_id"`
	PositionID   *string    `json:"position_id" form:"position_id"`
	DepartmentID *string    `json:"department_id" form:"department_id"`
	SectionID    *string    `json:"section_id" form:"section_id"`
	WhsID        *string    `json:"whs_id" form:"whs_id"`
	IsActive     bool       `gorm:"null" json:"is_active" form:"is_active" default:"false"`
	CreatedAt    time.Time  `json:"created_at" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at" default:"now"`
	User         User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	Whs          Whs        `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"whs"`
	Position     Position   `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"position"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"department"`
	Section      Section    `gorm:"foreignKey:SectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"section"`
	PrefixName   PrefixName `gorm:"foreignKey:PrefixNameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"prefix_name"`
}

func (obj *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
