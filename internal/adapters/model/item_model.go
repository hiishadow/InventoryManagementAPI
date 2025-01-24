package model

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemStatus string

const (
	BUY  ItemStatus = "BUY"
	SELL ItemStatus = "SELL"
)

func (s *ItemStatus) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid data type for Status: %T", value)
	}

	switch ItemStatus(strVal) {
	case BUY, SELL:
		*s = ItemStatus(strVal)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %s", strVal)
	}
}

func (s ItemStatus) Value() (driver.Value, error) {
	switch s {
	case BUY, SELL:
		return string(s), nil
	default:
		return nil, fmt.Errorf("invalid value for Status: %s", s)
	}
}

type Item struct {
	ID          uuid.UUID    `gorm:"primaryKey" json:"id"`
	ProductName string       `gorm:"not null ;unique" json:"productName"`
	Status      ItemStatus   `gorm:"type:item_status;not null" json:"status"`
	Price       float64      `gorm:"not null" json:"price"`
	Amount      int          `gorm:"not null" json:"amount"`
	At          time.Time    `gorm:"index;not null" json:"at"`
	DeletedAt   sql.NullTime `gorm:"index" json:"-"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}
