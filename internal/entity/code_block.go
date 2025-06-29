package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CodeBlock struct {
	Id             int       `json:"id"`
	Uuid           string    `json:"uuid"`
	Code           string    `json:"code"`
	InsertDateTime time.Time `json:"insert_date_time"`
	UpdateDateTime time.Time `json:"update_date_time"`
}

func (receiver *CodeBlock) BeforeCreate(tx *gorm.DB) (err error) {
	receiver.Uuid = uuid.NewString()
	receiver.InsertDateTime = time.Now()
	receiver.UpdateDateTime = time.Now()
	return nil
}

func (receiver *CodeBlock) BeforeUpdate(tx *gorm.DB) (err error) {
	receiver.UpdateDateTime = time.Now()
	return nil
}
