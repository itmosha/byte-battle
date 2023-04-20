package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	config       *Config
	db           *sql.DB
	usersRepo    *UserRepository
	coursesRepo  *CourseRepository
	profilesRepo *ProfileRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	dbConnectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		s.config.DatabaseHost, s.config.DatabasePort, s.config.DatabaseName,
		s.config.DatabaseUser, s.config.DatabasePassword, s.config.DatabaseSSLMode)

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.usersRepo != nil {
		return s.usersRepo
	}

	s.usersRepo = &UserRepository{
		store: s,
	}

	return s.usersRepo
}

func (s *Store) Course() *CourseRepository {
	if s.coursesRepo != nil {
		return s.coursesRepo
	}

	s.coursesRepo = &CourseRepository{
		store: s,
	}

	return s.coursesRepo
}

func (s *Store) Profile() *ProfileRepository {
	if s.profilesRepo != nil {
		return s.profilesRepo
	}

	s.profilesRepo = &ProfileRepository{
		store: s,
	}

	return s.profilesRepo
}
