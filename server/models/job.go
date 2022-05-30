package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// type Word struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	JobId     uint           `json:"jobId"`
// 	Name      string         `json:"name" gorm:"not null"`
// 	CreatedAt time.Time      `json:"createdAt"`
// 	UpdatedAt time.Time      `json:"updatedAt"`
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

type Job struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	Tags        datatypes.JSON `json:"tags"` //  gorm:"foreignKey:JobId;type:json"
	// KeyWords    []Word         `json:"keyWords"`
	Offers      uint8          `json:"offers" gorm:"not null"`
	SeenAt      time.Time      `json:"seenAt" gorm:"not null"`
	TimeLeft    time.Time      `json:"timeLeft" gorm:"not null"`
	PublishedAt time.Time      `json:"publishedAt" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func Jobs() ([]Job, error) {
	var jobs []Job

	result := db.Find(&jobs)

	// result := db.Table("jobs as jb").Select(`
	// 	title,
	// 	description,
	// 	offers,
	// 	seen_at,
	// 	time_left,
	// 	published_at,
	// 	jb.created_at,
	// 	jb.updated_at,
	// 	json_agg(
	// 		json_build_object(
	// 				'jobId', wd.job_id,
	// 				'name', wd.name
	// 		)
	// 	) as tags
	// `).Joins(`
	// 	JOIN words as wd
	// 	ON wd.job_id = jb.id
	// `).Group("jb.id").Find(&jobs)

	if result.Error != nil {
		return []Job{}, result.Error
	}

	return jobs, nil
}
