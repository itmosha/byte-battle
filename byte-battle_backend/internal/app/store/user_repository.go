package store

import (
	"byte-battle_backend/internal/app/model"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateInstance(user *model.User) (*model.User, error) {

	/**

	Get a User instance without ID, perform INSERT, get ID and return instance with all fields.

	*/

	query := fmt.Sprintf(
		"INSERT INTO users (username, email, role, encrypted_pwd) VALUES ('%s', '%s', %d, '%s') RETURNING id;",
		user.Username, user.Email, user.Role, user.EncryptedPwd,
	)

	err := r.store.db.QueryRow(query).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetInstance(id int) (*model.User, error) {

	/**

	Get User instance with provided ID.

	*/

	query := fmt.Sprintf("SELECT id, username, email, role FROM users WHERE id=%d;", id)
	var user model.User

	err := r.store.db.QueryRow(query).Scan(&user.ID, &user.Username, &user.Email, &user.Role)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CheckUniqueValue(fieldName string, value string) (bool, error) {

	/**

	Perform a SELECT query into the database and check if provided value already exists.
	This is used for checking unique usernames and emails.

	*/

	query := fmt.Sprintf("SELECT id FROM users WHERE %s='%s';", fieldName, value)
	var id int

	err := r.store.db.QueryRow(query).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, nil
}
