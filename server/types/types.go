package types

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceStore interface {
	GetTodaysClasses() ([]string, error)
	GetAllSubjects() ([]string, error)
	MarkAttendance(*AttendanceRequest, uint) error
	GetAttendanceSummary(userId uint) (*AttendanceSummary, error)
}

type AttendanceRequest struct {
	Date     string            `json:"date"`
	Subjects []SubjectsRequest `json:"subjects"`
}
type SubjectsRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type UserStore interface {
	CreateSubjectsForUser(tx *gorm.DB, userID uint) error
	CreateUser(*User) error
	GetUserByIdentifier(string) (*User, error)
	GetUserById(uint) (*User, error)
	UpdateUser(*User) error
	DeleteUser(uint) error
}
type LoginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

// User Model (You already have this, but included for completeness)
type User struct {
	gorm.Model
	Register string `gorm:"unique" json:"register_no" validate:"required,min=12,max=12"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required,min=4,max=24"`
	Password string `json:"password" validate:"required"` // Hashed password
	Role     string `json:"role"`                         // student/teacher/admin
}

// Subject Model - Represents a subject a student is enrolled in
type Subject struct {
	gorm.Model
	Name       string    `gorm:"not null" json:"name"`
	UserID     uint      `gorm:"not null;index" json:"user_id"` // Each user has their own subjects
	StartDate  time.Time `gorm:"not null" json:"start_date"`
	TotalTaken int       `gorm:"default:0" json:"total_taken"`
	Attended   int       `gorm:"default:0" json:"attended"`
	Missed     int       `gorm:"default:0" json:"missed"`
}

// Schedule Model - Defines which days a subject has classes
type Schedule struct {
	gorm.Model
	SubjectName string `json:"subject_name"` // Links to Subject
	DayOfWeek   string `json:"day_of_week"`  // 0 = Sunday, 6 = Saturday
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

// Attendance Model - Tracks if a student attended or bunked a class
type Attendance struct {
	gorm.Model
	UserID    uint      `json:"user_id" gorm:"index:idx_user_subject_date,unique"`
	SubjectID uint      `json:"subject_id" gorm:"index:idx_user_subject_date,unique"`
	Date      time.Time `json:"date" gorm:"index:idx_user_subject_date,unique"`
	Status    string    `json:"status" validate:"required,oneof=present absent leave not_taken"`
}

type SubjectSummary struct {
	Name         string `json:"name"`
	Taken        int    `json:"taken"`
	Attended     int    `json:"attended"`
	Skipped      int    `json:"skipped"`
	AllowedSkips int    `json:"allowed_skips"`
}

type SubjectStats struct {
	SubjectName string `json:"subject_name"`
	Total       int    `json:"total"`
	Attended    int    `json:"attended"`
	Skipped     int    `json:"skipped"`
}

type AttendanceSummary struct {
	TotalClasses int            `json:"total_classes"`
	Attended     int            `json:"attended"`
	Skipped      int            `json:"skipped"`
	AllowedSkips int            `json:"allowed_skips"`
	Subjects     []SubjectStats `json:"subjects"`
}
