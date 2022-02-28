package entity

// CourseRepository ...
type CourseRepository interface {
	Insert(course Course) error
}
