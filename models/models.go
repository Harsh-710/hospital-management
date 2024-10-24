package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ID        int      `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string   `json:"firstname" gorm:"not null" binding:"required"`          // Required field
	LastName  string   `json:"lastname" gorm:"not null" binding:"required"`           // Required field
	Age       int      `json:"age" gorm:"not null" binding:"required,gt=0"`           // Required, must be greater than 0
	Gender    string   `json:"gender" gorm:"not null" binding:"required"`             // Required field
	Phone     string   `json:"phone" gorm:"not null" binding:"required"`              // Required field
	Email     string   `json:"email" gorm:"not null;unique" binding:"required,email"` // Required, must be unique, email format
	Diagnosis string   `json:"diagnosis" gorm:"not null" binding:"required"`          // Required field
	Condition string   `json:"condition" gorm:"not null" binding:"required"`          // Required field
	Tests     []string `json:"tests"`
}

type User struct {
	gorm.Model
	ID        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"firstname" gorm:"not null" binding:"required"`          // Required field
	LastName  string `json:"lastname" gorm:"not null" binding:"required"`           // Required field
	Age       int    `json:"age" gorm:"not null" binding:"required,gt=0"`           // Required, must be greater than 0
	Gender    string `json:"gender" gorm:"not null" binding:"required"`             // Required field
	Phone     string `json:"phone" gorm:"not null" binding:"required"`              // Required field
	Email     string `json:"email" gorm:"not null;unique" binding:"required,email"` // Required, must be unique, email format
	Password  string `json:"-" gorm:"not null" binding:"required"`                  // Required field
	Role      string `json:"role" gorm:"default:receptionist"`                      // Default role value
}

type Appointment struct {
	gorm.Model
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	PatientID     int    `json:"patientID" gorm:"not null" binding:"required"` // Required field
	DoctorID      int    `json:"doctorID" gorm:"not null" binding:"required"`  // Required field
	Date          string `json:"date" gorm:"not null" binding:"required"`      // Required field
	Time          string `json:"time" gorm:"not null" binding:"required"`      // Required field
	Reason        string `json:"reason" gorm:"not null" binding:"required"`    // Required field
	Status        string `json:"status" gorm:"default:pending"`                // Default status value = pending
}

type CreatePatientPayload struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Phone     string   `json:"phone"`
	Email     string   `json:"email"`
	Diagnosis string   `json:"diagnosis"`
	Condition string   `json:"condition"`
	Tests     []string `json:"tests"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type PatientStore interface {
	GetPatientByID(id int) (*Patient, error)
	GetPatientsByID(ids []int) ([]Patient, error)
	GetPatients() ([]*Patient, error)
	CreatePatient(CreatePatientPayload) error
	UpdatePatient(Patient) error
}
