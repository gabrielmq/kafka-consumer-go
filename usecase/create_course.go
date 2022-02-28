package usecase

import "github.com/gabrielmq/kafka-consumer-go/entity"

type CreateCourse struct {
	Repository entity.CourseRepository
}

func NewCreateCourse(repository entity.CourseRepository) CreateCourse {
	return CreateCourse{Repository: repository}
}

func (c CreateCourse) Execute(input CreateCourseInputDto) (CreateCourseOutputDto, error) {
	course := entity.NewCourseOf(
		input.Name,
		input.Description,
		input.Status,
	)

	err := c.Repository.Insert(course)
	if err != nil {
		return NewCreateCourseOutputDto(), err
	}

	return NewCreateCourseOutputDtoOf(course), nil
}
