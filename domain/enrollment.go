package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID        string           `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID    string           `json:"user_id,omitempty" gorm:"type:char(50)"`
	User      *User            `json:"user,omitempty" gorm:"-"`
	CourseId  string           `json:"course_id,omitempty" gorm:"type:char(50);not null"`
	Course    *Course          `json:"course,omitempty" gorm:"-"`
	Status    EnrollmentStatus `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time       `json:"-"`
	UpdatedAt *time.Time       `json:"-"`
}

type EnrollmentStatus string

const (
	Pending  EnrollmentStatus = "P"
	Active   EnrollmentStatus = "A"
	Studying EnrollmentStatus = "S"
	Inactive EnrollmentStatus = "I"
)

func (c *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}
