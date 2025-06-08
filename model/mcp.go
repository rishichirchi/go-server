package models

import (
	"time"

	"github.com/lib/pq"
)

type MCP struct {
	Name        string         `json:"name" gorm:"not null;type:text"`
	Version     string         `json:"version" gorm:"type:text"`
	Description string         `json:"description" gorm:"type:text"`
	Author      string         `json:"author" gorm:"type:text"`
	License     string         `json:"license" gorm:"type:text"`
	Keywords    pq.StringArray `json:"keywords" gorm:"type:text[]"`
	Repository  Repository     `json:"repository" gorm:"embedded;embeddedPrefix:repository_"`
	Run         Run            `json:"run" gorm:"embedded;embeddedPrefix:run_"`
	ObjectKey   string         `json:"object_key" gorm:"type:text"`
	Overview    string         `json:"overview" gorm:"type:text"`
	Tools       string         `json:"tools" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
}

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Run struct {
	Command string         `json:"command"`
	Args    pq.StringArray `json:"args" gorm:"type:text[]"`
	Port    int            `json:"port"`
}
