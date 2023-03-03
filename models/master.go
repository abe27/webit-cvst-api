package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type SystemLog struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *SystemLog) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Whs struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Whs) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type PrefixName struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *PrefixName) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Position struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Position) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Department struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Department) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Section struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *Section) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OperatingSystem struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *OperatingSystem) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type OfficeVersion struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string    `json:"description,omitempty" form:"description" binding:"required"`
	IsActive    bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt   time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *OfficeVersion) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type MaintenanceStatus struct {
	ID          string `json:"id,omitempty"`
	Title       string `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string `json:"description,omitempty" form:"description" binding:"required"`
	// red,green,yellow,blue,purple
	Color     string    `json:"color,omitempty" form:"color"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *MaintenanceStatus) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type RequestStatus struct {
	ID          string `json:"id,omitempty"`
	Title       string `gorm:"not null;size:50;unique;" json:"title,omitempty" form:"title" binding:"required"`
	Description string `json:"description,omitempty" form:"description" binding:"required"`
	// red,green,yellow,blue,purple
	Color     string    `json:"color,omitempty" form:"color"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
}

func (obj *RequestStatus) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type Approvement struct {
	ID           string     `json:"id,omitempty"`
	DepartmentID *string    `json:"department_id,omitempty" form:"department_id"`
	UserID       *string    `json:"user_id,omitempty" form:"user_id"`
	LineToken    string     `json:"line_token,omitempty" form:"line_token"`
	Description  string     `json:"description,omitempty" form:"description" binding:"required"`
	IsActive     bool       `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt    time.Time  `json:"created_at,omitempty" default:"now"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty" default:"now"`
	User         User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"department,omitempty"`
}

func (obj *Approvement) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
