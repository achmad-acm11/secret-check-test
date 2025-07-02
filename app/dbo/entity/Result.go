package entity

import (
	"gorm.io/gorm"
	"time"
)

type Result struct {
	Id           int            `gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	ProjectId    int            `gorm:"column:project_id;type:int"`
	ProjectKey   string         `gorm:"column:project_key;type:varchar(255)"`
	Path         string         `gorm:"column:path;type:varchar(255)"`
	Category     string         `gorm:"column:category;type:varchar(255)"`
	Rule         string         `gorm:"column:rule;type:varchar(255)"`
	ScanType     string         `gorm:"column:scan_type;type:varchar(255)"`
	Title        string         `gorm:"column:title;type:text"`
	Severity     string         `gorm:"column:severity;type:varchar(255)"`
	LastFoundAt  string         `gorm:"column:last_found_at;type:varchar(255)"`
	StatusResult int            `gorm:"column:status_result;type:int"`
	ScanVersion  int            `gorm:"column:scan_version;type:int"`
	Code         string         `gorm:"column:code;type:varchar(255)"`
	StartLine    int            `gorm:"column:start_line;type:int"`
	EndLine      int            `gorm:"column:end_line;type:int"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;->"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null;->"`
}

type LineCode struct {
	Number      int    `json:"Number"`
	Content     string `json:"Content"`
	IsCause     bool   `json:"IsCause"`
	Annotation  string `json:"Annotation"`
	Truncated   bool   `json:"Truncated"`
	Highlighted string `json:"Highlighted,omitempty"`
	FirstCause  bool   `json:"FirstCause"`
	LastCause   bool   `json:"LastCause"`
}

type SecretCode struct {
	Lines []LineCode `json:"Lines"`
}

func (Result) TableName() string {
	return "results"
}
