package appointment

import (
	"github.com/Harsh-710/hospital-management/models"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateAppointment(appointment models.Appointment) (int, error) {
	result := s.db.Create(&appointment)
	if result.Error != nil {
		return 0, result.Error
	}
	return appointment.ID, nil // Return the auto-incremented ID after creating the order
}
