package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Asset struct {
	ID                string          `json:"id,omitempty"`
	TypeID            *string         `json:"type_id,omitempty" form:"type_id"`
	ModelID           *string         `json:"model_id,omitempty" form:"model_id"`
	ModelNameID       *string         `json:"model_name_id,omitempty" form:"model_name_id"`
	SerialNo          string          `gorm:"index;unique;size:25;not null;" json:"serial_no,omitempty" form:"serial_no"`
	Price             float64         `json:"price,omitempty" form:"price"`
	ReceiveDate       time.Time       `json:"receive_date,omitempty" form:"receive_date"`
	AssetCode         string          `gorm:"size:25;" json:"asset_code,omitempty" form:"asset_code"`
	AccountNo         string          `gorm:"size:25;" json:"account_no,omitempty" form:"account_no"`
	OwnerID           *string         `json:"owner_id,omitempty" form:"owner_id"`
	OperatingSystemID *string         `json:"operating_system_id,omitempty" form:"operating_system_id"`
	OfficeVersionID   *string         `json:"office_version_id,omitempty" form:"office_version_id"`
	HasAntivirus      bool            `json:"has_antivirus,omitempty" form:"has_antivirus"`
	UseInternet       bool            `json:"use_internet,omitempty" form:"use_internet"`
	UseEmail          bool            `json:"use_email,omitempty" form:"use_email"`
	WindowIsUpdate    bool            `json:"window_is_update,omitempty" form:"window_is_update"`
	IpAddress         string          `json:"ip_address,omitempty" form:"ip_address"`
	Description       string          `json:"description,omitempty" form:"description" binding:"required"`
	IsActive          bool            `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt         time.Time       `json:"created_at,omitempty" default:"now"`
	UpdatedAt         time.Time       `json:"updated_at,omitempty" default:"now"`
	AssetType         AssetType       `gorm:"foreignKey:TypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset_type,omitempty"`
	AssetModel        AssetModel      `gorm:"foreignKey:ModelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset_model,omitempty"`
	AssetModelName    AssetModelName  `gorm:"foreignKey:ModelNameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset_model_name,omitempty"`
	Profile           Profile         `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner,omitempty"`
	OperatingSystem   OperatingSystem `gorm:"foreignKey:OperatingSystemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"operating_system,omitempty"`
	OfficeVersion     OfficeVersion   `gorm:"foreignKey:OfficeVersionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"office_version,omitempty"`
}

func (obj *Asset) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ImageAsset struct {
	ID        string    `json:"id,omitempty"`
	AssetID   *string   `json:"asset_id,omitempty" form:"asset_id" binding:"required"`
	UrlImage  string    `json:"url_image,omitempty" form:"url_image" binding:"required"`
	IsActive  bool      `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty" default:"now"`
	UpdatedAt time.Time `json:"updated_at,omitempty" default:"now"`
	Asset     Asset     `gorm:"foreignKey:AssetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset,omitempty"`
}

func (obj *ImageAsset) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type RepairRequest struct {
	ID              string        `json:"id,omitempty"`
	RequestDate     time.Time     `json:"request_date,omitempty" form:"request_date"`
	RequestNo       string        `gorm:"size:25;" json:"request_no" form:"request_no"`
	RequestTypeID   *string       `json:"request_type_id,omitempty" form:"request_type_id"`
	RequestByID     *string       `json:"request_by_id,omitempty" form:"request_by_id"`
	AssetID         *string       `json:"asset_id,omitempty" form:"asset_id" binding:"required"`
	ProbleDetail    string        `json:"proble_detail,omitempty" form:"proble_detail"`
	RequestStatusID *string       `json:"request_status_id,omitempty" form:"request_status_id" binding:"required"`
	IsActive        bool          `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt       time.Time     `json:"created_at,omitempty" default:"now"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" default:"now"`
	User            User          `gorm:"foreignKey:RequestByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"request_by,omitempty"`
	Asset           Asset         `gorm:"foreignKey:AssetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset,omitempty"`
	RequestStatus   RequestStatus `gorm:"foreignKey:RequestStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"status,omitempty"`
	RequestType     RequestType   `gorm:"foreignKey:RequestTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"request_type,omitempty"`
}

func (obj *RepairRequest) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ImageRequest struct {
	ID              string        `json:"id,omitempty"`
	RepairRequestID *string       `json:"repair_request_id,omitempty" form:"repair_request_id" binding:"required"`
	UrlImage        string        `json:"url_image,omitempty" form:"url_image" binding:"required"`
	IsActive        bool          `json:"is_active,omitempty" form:"is_active" binding:"required"`
	CreatedAt       time.Time     `json:"created_at,omitempty" default:"now"`
	UpdatedAt       time.Time     `json:"updated_at,omitempty" default:"now"`
	RepairRequest   RepairRequest `gorm:"foreignKey:RepairRequestID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"request_id,omitempty"`
}

func (obj *ImageRequest) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
