package entity

import (
	"github.com/google/uuid"
)

type Course struct {
	ID          string
	Name        string
	Description string
	Status      string
}

func NewCourseOf(name, description, status string) Course {
	return Course{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Status:      status,
	}
}
