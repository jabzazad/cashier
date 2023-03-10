package models

import (
	"time"

	"gorm.io/gorm"
)

// Model base model
type Model struct {
	ID        uint           `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"-" sql:"index"`
}

// ModelInterface model interface
type ModelInterface interface {
	GetID() uint
	SetID(id uint)
	Stamp()
	UpdateStamp()
	DeleteStamp()
}

// GetID get id
func (model *Model) GetID() uint {
	return model.ID
}

// SetID set id
func (model *Model) SetID(id uint) {
	model.ID = id
}

// DeleteStamp soft delete updated, deleted_at model
func (model *Model) DeleteStamp() {
	model.UpdateStamp()
	model.DeletedAt = gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}
}

// UpdateStamp current updated at model
func (model *Model) UpdateStamp() {
	model.UpdatedAt = time.Now()
}

// Stamp current time to model
func (model *Model) Stamp() {
	timeNow := time.Now()
	model.UpdatedAt = timeNow
	model.CreatedAt = timeNow
}
