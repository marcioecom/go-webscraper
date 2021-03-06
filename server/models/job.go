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
	Url         string         `json:"url"`
	Offers      string         `json:"offers" gorm:"not null"`
	Interested  string         `json:"interested" gorm:"not null"`
	SeenAt      time.Time      `json:"seenAt" gorm:"not null"`
	TimeLeft    int64          `json:"timeLeft" gorm:"not null"`
	PublishedAt int64          `json:"publishedAt" gorm:"not null"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func Jobs(page int) ([]Job, int64, error) {
	var jobs []Job
	var totalJobs int64

	db.Table("jobs").Count(&totalJobs)
	if page == 0 {
		result := db.Order("updated_at desc").Find(&jobs)

		if result.Error != nil {
			return []Job{}, 0, result.Error
		}

		return jobs, totalJobs, nil
	}

	page -= 1
	result := db.Order("updated_at desc").Offset(page * 10).Limit(10).Find(&jobs)
	if result.Error != nil {
		return []Job{}, 0, result.Error
	}

	return jobs, totalJobs, nil
}

func CreateJob(job Job) (Job, error) {
	tx := db.Create(&job)
	if tx.Error != nil {
		return Job{}, tx.Error
	}

	return job, nil
}

func CreateJobs(jobs []Job) error {
	tx := db.Create(&jobs)
	return tx.Error
}

func UpdateJob(job Job) {
	if db.Model(&job).Where("title = ?", job.Title).Updates(&job).RowsAffected == 0 {
		tx := db.Create(&job)
		if tx.Error != nil {
			panic(tx.Error)
		}
	}
}
