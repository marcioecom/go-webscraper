package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Job struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Tags        datatypes.JSON `json:"tags"` //  gorm:"foreignKey:JobId;type:json"
	// KeyWords    []Word         `json:"keyWords"`
	Offers      string         `json:"offers" gorm:"not null"`
	SeenAt      time.Time      `json:"seenAt" gorm:"not null"`
	TimeLeft    int64          `json:"timeLeft" gorm:"not null"`
	PublishedAt int64          `json:"publishedAt" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func Jobs() ([]Job, error) {
	var jobs []Job

	result := db.Find(&jobs)

	if result.Error != nil {
		return []Job{}, result.Error
	}

	return jobs, nil
}
