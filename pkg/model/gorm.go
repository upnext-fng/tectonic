package model

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	ID        uint64         `json:"id" gorm:"primary_key;column:id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index;column:deleted_at"`
}

type UUID struct {
	ID        string         `json:"id" gorm:"TYPE:VARCHAR(50);NOT NULL;PRIMARY_KEY;column:id"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index;column:deleted_at"`
}
