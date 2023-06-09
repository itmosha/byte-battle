package store

import (
	"byte-battle_backend/internal/app/model"
	"fmt"
)

type ProfileRepository struct {
	store *Store
}

func (r *ProfileRepository) CreateInstance(profile *model.Profile) (*model.Profile, error) {

	query := fmt.Sprintf(
		"INSERT INTO profiles (id, username, email) VALUES (%d, '%s', '%s') RETURNING id;",
		profile.ID, profile.Username, profile.Email,
	)

	var id int

	err := r.store.db.QueryRow(query).Scan(&id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *ProfileRepository) GetInstance(id int) (*model.Profile, error) {

	query := fmt.Sprintf(
		"SELECT id, username, email, status, image_url, tasks_completed, active_courses, completed_courses FROM profiles WHERE id=%d;",
		id,
	)

	var profile model.Profile

	err := r.store.db.QueryRow(query).Scan(
		&profile.ID, &profile.Username, &profile.Email, &profile.Status,
		&profile.ImageUrl, &profile.TasksCompleted, &profile.ActiveCourses, &profile.CompletedCourses,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}
