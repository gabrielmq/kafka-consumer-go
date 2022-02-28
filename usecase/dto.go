package usecase

import "github.com/gabrielmq/kafka-consumer-go/entity"

type CreateCourseInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type CreateCourseOutputDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewCreateCourseOutputDtoOf(course entity.Course) CreateCourseOutputDto {
	return CreateCourseOutputDto{
		ID:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Status:      course.Status,
	}
}

func NewCreateCourseOutputDto() CreateCourseOutputDto {
	return CreateCourseOutputDto{}
}
