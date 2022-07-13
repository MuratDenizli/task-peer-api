package models

import (
	"strings"
	u "task-peer-api/utils"

	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Type   string `json:"type" gorm:"not null"`
	Status string `json:"status" gorm:"not null"`
}

func (task *Task) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(task.Name, "") {
		return u.Message(false, "Task name cannot be empty"), false
	}

	if len(task.Name) > 500 {
		return u.Message(false, "Task name too long"), false
	}
	return nil, false
}

func (task *Task) Create() map[string]interface{} {
	if resp, ok := task.Validate(); !ok {
		return resp
	}

	GetDB().Create(task)

	if task.ID <= 0 {
		return u.Message(false, "Bağlantı hatası oluştu. Task oluşturulamadı.")
	}
	response := u.Message(true, "Hesap başarıyla yaratıldı!")
	response["task"] = task
	return response
}
