package models

import (
	"time"

	g "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Asset struct {
	ID                string          `json:"id"`
	TypeID            *string         `json:"type_id" form:"type_id"`
	ModelID           *string         `json:"model_id" form:"model_id"`
	ModelNameID       *string         `json:"model_name_id" form:"model_name_id"`
	SerialNo          string          `gorm:"index;unique;size:25;not null;" json:"serial_no" form:"serial_no"`
	Price             float64         `json:"price" form:"price"`
	ReceiveDate       time.Time       `json:"receive_date" form:"receive_date"`
	AssetCode         string          `gorm:"size:25;" json:"asset_code" form:"asset_code"`
	AccountNo         string          `gorm:"size:25;" json:"account_no" form:"account_no"`
	OwnerID           *string         `json:"owner_id" form:"owner_id"`
	OperatingSystemID *string         `json:"operating_system_id" form:"operating_system_id"`
	OfficeVersionID   *string         `json:"office_version_id" form:"office_version_id"`
	HasAntivirus      bool            `json:"has_antivirus" form:"has_antivirus"`
	UseInternet       bool            `json:"use_internet" form:"use_internet"`
	UseEmail          bool            `json:"use_email" form:"use_email"`
	WindowIsUpdate    bool            `json:"window_is_update" form:"window_is_update"`
	IpAddress         string          `json:"ip_address" form:"ip_address"`
	Description       string          `json:"description" form:"description" binding:"required"`
	IsActive          bool            `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt         time.Time       `json:"created_at" default:"now"`
	UpdatedAt         time.Time       `json:"updated_at" default:"now"`
	AssetType         AssetType       `gorm:"foreignKey:TypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset_type"`
	AssetModel        AssetModel      `gorm:"foreignKey:ModelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset_model"`
	AssetModelName    AssetModelName  `gorm:"foreignKey:ModelNameID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset_model_name"`
	Profile           Profile         `gorm:"foreignKey:OwnerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"owner"`
	OperatingSystem   OperatingSystem `gorm:"foreignKey:OperatingSystemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"operating_system"`
	OfficeVersion     OfficeVersion   `gorm:"foreignKey:OfficeVersionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"office_version"`
}

func (obj *Asset) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ImageAsset struct {
	ID        string    `json:"id"`
	AssetID   *string   `json:"asset_id" form:"asset_id" binding:"required"`
	UrlImage  string    `json:"url_image" form:"url_image" binding:"required"`
	IsActive  bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt time.Time `json:"created_at" default:"now"`
	UpdatedAt time.Time `json:"updated_at" default:"now"`
	Asset     Asset     `gorm:"foreignKey:AssetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset"`
}

func (obj *ImageAsset) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type RepairEquipmentRequest struct {
	ID              string        `json:"id"`
	RequestDate     time.Time     `json:"request_date" form:"request_date"`
	RequestNo       string        `gorm:"size:25;" json:"request_no" form:"request_no"`
	RequestTypeID   *string       `json:"request_type_id" form:"request_type_id"`
	RequestByID     *string       `json:"request_by_id" form:"request_by_id"`
	AssetID         *string       `json:"asset_id" form:"asset_id" binding:"required"`
	ProbleDetail    string        `json:"proble_detail" form:"proble_detail"`
	RequestStatusID *string       `json:"request_status_id" form:"request_status_id" binding:"required"`
	IsActive        bool          `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt       time.Time     `json:"created_at" default:"now"`
	UpdatedAt       time.Time     `json:"updated_at" default:"now"`
	User            User          `gorm:"foreignKey:RequestByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"request_by"`
	Asset           Asset         `gorm:"foreignKey:AssetID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"asset"`
	RequestStatus   RequestStatus `gorm:"foreignKey:RequestStatusID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"status"`
	RequestType     RequestType   `gorm:"foreignKey:RequestTypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"request_type"`
}

func (obj *RepairEquipmentRequest) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ImageRequest struct {
	ID                     string                 `json:"id"`
	RepairRequestID        *string                `json:"repair_request_id" form:"repair_request_id" binding:"required"`
	UrlImage               string                 `json:"url_image" form:"url_image" binding:"required"`
	IsActive               bool                   `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt              time.Time              `json:"created_at" default:"now"`
	UpdatedAt              time.Time              `json:"updated_at" default:"now"`
	RepairEquipmentRequest RepairEquipmentRequest `gorm:"foreignKey:RepairRequestID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"request_id"`
}

func (obj *ImageRequest) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}

type ConfirmEquipment struct {
	ID                     string                 `json:"id"`
	RepairRequestID        *string                `json:"repair_request_id" form:"repair_request_id" binding:"required"`
	ConfirmByID            *string                `json:"accep_by_id" form:"accep_by_id"`
	Recording              string                 `json:"recording" form:"recording"`
	Remark                 string                 `json:"remark" form:"remark"`
	IsActive               bool                   `json:"is_active" form:"is_active" binding:"required"`
	CreatedAt              time.Time              `json:"created_at" default:"now"`
	UpdatedAt              time.Time              `json:"updated_at" default:"now"`
	RepairEquipmentRequest RepairEquipmentRequest `gorm:"foreignKey:RepairRequestID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"request_id"`
	Profile                Profile                `gorm:"foreignKey:ConfirmByID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"confirm_by"`
}

func (obj *ConfirmEquipment) BeforeCreate(tx *gorm.DB) (err error) {
	id, _ := g.New()
	obj.ID = id
	return
}
