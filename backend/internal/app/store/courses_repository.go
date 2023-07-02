package store

import (
	"byte-battle_backend/internal/app/model"
	"fmt"
)

type CourseRepository struct {
	store *Store
}

func (r *CourseRepository) CreateInstance(course *model.Course) (*model.Course, error) {

	query := fmt.Sprintf("INSERT INTO courses (title) VALUES ('%s') RETURNING id;", course.Title)

	err := r.store.db.QueryRow(query).Scan(&course.ID)

	if err != nil {
		return nil, err
	}

	return course, nil
}

func (r *CourseRepository) GetInstance(id int) (*model.Course, error) {

	query := fmt.Sprintf("SELECT id, title FROM courses WHERE id=%d;", id)
	var course model.Course

	err := r.store.db.QueryRow(query).Scan(&course.ID, &course.Title)

	if err != nil {
		return nil, err
	}

	return &course, nil
}
