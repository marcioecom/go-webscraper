package models

type Goly struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goly     string `json:"goly" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
	Random   bool   `json:"random"`
}

func CreateGoly(goly Goly) error {
	tx := db.Create(&goly)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
