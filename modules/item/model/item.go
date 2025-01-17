package model

import (
	"errors"
	"golandtitorial/common"
)

var (
	ErrTitleIsBlank = errors.New("title is blank")
	ErritemDeleted = errors.New("item is deleted")
)

type TodoItem struct {
	common.SqlModel
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string {
	return "users"
}

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	//Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string {
	return TodoItem{}.TableName()
}
