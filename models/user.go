package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserName  string    `gorm:"not null;column:username;index;unique;size:10" json:"username,omitempty" form:"username"`
	Email     string    `gorm:"not null;unique;size:50;" json:"email,omitempty" form:"email"`
	Password  string    `gorm:"not null;unique;size:60;" json:"-" form:"password"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Administrator struct {
	ID        string    `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	UserID    *string   `gorm:"unique;" json:"user_id,omitempty" form:"user_id"`
	IsActive  bool      `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
	User      User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (obj *Administrator) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Profile struct {
	ID           string     `gorm:"primaryKey;size:21;" json:"id,omitempty"`
	AvatarURL    string     `json:"avatar_url,omitempty" form:"avatar_url"`
	PrefixNameID *string    `json:"prefix_name_id,omitempty" form:"prefix_name_id"`
	FirstName    string     `gorm:"not null;size:50;" json:"first_name,omitempty" form:"first_name"`
	LastName     string     `gorm:"not null;size:50;" json:"last_name,omitempty" form:"last_name"`
	PhoneNo      string     `json:"phone_no,omitempty" form:"phone_no"`
	UserID       *string    `gorm:"not null;unique;" json:"user_id,omitempty" form:"user_id"`
	PositionID   *string    `json:"position_id,omitempty" form:"position_id"`
	DepartmentID *string    `json:"department_id,omitempty" form:"department_id"`
	SectionID    *string    `json:"section_id,omitempty" form:"section_id"`
	WhsID        *string    `json:"whs_id,omitempty" form:"whs_id"`
	IsActive     bool       `gorm:"null" json:"is_active,omitempty" form:"is_active" default:"false"`
	CreatedAt    time.Time  `json:"created_at,omitempty" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty" default:"now"`
	User         User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Whs          Whs        `gorm:"foreignKey:WhsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"whs,omitempty"`
	Position     Position   `gorm:"foreignKey:PositionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"position,omitempty"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"department,omitempty"`
	Section      Section    `gorm:"foreignKey:SectionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"section,omitempty"`
	PrefixName   PrefixName `gorm:"foreignKey:PrefixNameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"prefix_name,omitempty"`
}

func (obj *Profile) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
