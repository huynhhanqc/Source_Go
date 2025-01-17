package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = []string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) String() string {
	return allItemStatus[*item]
}

func ParseStringToItemStatus(str string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == str {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("scan error: %s", value))
	}
	v, err := ParseStringToItemStatus(string(bytes))
	if err != nil {
		return errors.New(fmt.Sprintf("scan error: %s", value))
	}
	*item = v
	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := string(data)
	v, err := ParseStringToItemStatus(str)
	if err != nil {
		return err
	}
	*item = v
	return nil
}
