package store

import (
	"byte-battle_backend/internal/app/model"
	jwtfuncs "byte-battle_backend/pkg/jwt_funcs"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) CreateInstance(user *model.User) (*model.User, error) {

	/**

	This function creates a User row in users table with all the data required for authorization actions.
	Also when the row is inserted, this functions creates a row in Profiles table with all the user's account info.

	*/

	query := fmt.Sprintf(
		"INSERT INTO users (username, email, role, encrypted_pwd) VALUES ('%s', '%s', %d, '%s') RETURNING id;",
		user.Username, user.Email, user.Role, user.EncryptedPwd,
	)

	err := r.store.db.QueryRow(query).Scan(&user.ID)

	if err != nil {
		return nil, err
	}

	profile := &model.Profile{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	profile, err = r.store.Profile().CreateInstance(profile)

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

func (r *UserRepository) Login(user *model.User) (string, error) {

	/**

	Find the requested user's info in the db, check if the password is valid, make a JWT and send it back.
	Return error if that user doesn't exist or the password is wrong.

	*/

	var (
		query          string
		err            error
		encryptedPwdDb string
	)

	// Construct the query depending on what was provided: usename or email

	if user.Username == "" {
		query = fmt.Sprintf("SELECT id, username, role, encrypted_pwd FROM users WHERE email='%s';", user.Email)
		err = r.store.db.QueryRow(query).Scan(&user.ID, &user.Username, &user.Role, &encryptedPwdDb)
	} else if user.Username != "" {
		query = fmt.Sprintf("SELECT id, email, role, encrypted_pwd FROM users WHERE username='%s';", user.Username)
		err = r.store.db.QueryRow(query).Scan(&user.ID, &user.Email, &user.Role, &encryptedPwdDb)
	} else {
		return "", fmt.Errorf("No username or email was provided")
	}

	if err != nil {
		return "", err
	}

	// Check if the password is valid

	if user.EncryptedPwd != encryptedPwdDb {
		return "", fmt.Errorf("Incorrect password provided")
	}

	token, err := jwtfuncs.GenerateJWT(user.Username, user.Email, int8(user.Role))

	if err != nil {
		return "", err
	}

	return token, nil
}
