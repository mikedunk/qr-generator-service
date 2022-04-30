package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Qr struct {
	Id          uint64    `gorm:"auto_increment" json:"id"`
	Code        uuid.UUID `gorm:"primary_keyindex;unique;not null default:uuid_generate_v4()" json:"code"`
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Data        JSON      `json:"data"`
	ScannedAt   time.Time `gorm:"default:NULL" json:"scanned_at"`
	ExpiresAt   time.Time `gorm:"default:NULL;column:expires_at" json:"expires_at"`
	ScanCount   int64     `gorm:"column:scan_count" json:"scan_count"`
	ScannerId   string    `gorm:"column:scanner_id" json:"scanner_id"`
	gorm.Model
}

// type Data struct {
// 	DataId             string    `json:"data_id,omitempty"`
// 	Description        string    `json:"description,omitempty"`
// 	ImageUrl           string    `json:"imageUrl,omitempty"`
// 	Destination        string    `json:"radix,omitempty"`
// 	AdditionalText1    string    `json:"add_info_1,omitempty"`
// 	Email              string    `json:"email,omitempty"`
// 	SourceAddress      string    `json:"source_address,omitempty"`
// 	DestinationAddress string    `json:"destination_address,omitempty"`
// 	SourcePhone        string    `json:"source_phone,omitempty"`
// 	DestinationPhone   string    `json:"destination_phone,omitempty"`
// 	ExpiresAt          time.Time `json:"expires_at,omitempty"`
// 	Consummator        string    `json:"consummator,omitempty"`
// }

//{\"destination\":\"07016182391\",\"source\":\"08067036022\",\"radix\":\"08118765106\",\"email\":\"missgbosi@gmail.com\",\"scanner\":\"08118765106\",\"expiry\":\"Jun 20 2021 00:00:00 GMT+0100\"}"
func (base *Qr) BeforeCreate(tx *gorm.DB) error {
	code, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("code", code)
	return nil
}

type JSON json.RawMessage

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}
