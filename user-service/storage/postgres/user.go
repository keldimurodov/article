package postgres

import (
	"fmt"
	u "projects/article/user-service/genproto/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *u.User) (*u.User, error) {

	var res u.User
	query := `
		INSERT INTO users(
			first_name, 
			last_name,
			username, 
			avatar_url,
			bio,
			email,
			password
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING 
			id, 
			first_name, 
			last_name,
			username,
			avatar_url,
			bio,
			email,
			password,
			created_at,
			updeted_at
	`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.AvatarUrl,
		user.Bio,
		user.Email,
		user.Password).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.AvatarUrl,
		&res.Bio,
		&res.Email,
		&res.Password,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		fmt.Println("Error Scanning")
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) Get(id *u.GetUserRequest) (*u.User, error) {
	var user u.User
	query := `
	SELECT
		id,
		first_name,
		last_name,
		username,
		avatar_url,
		bio,
		email,
		password,
		created_at,
		updeted_at
	FROM 
		users
	WHERE
		id=$1 
	AND deleted_at IS NULL
	`
	err := r.db.QueryRow(query, id.UserId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.AvatarUrl,
		&user.Bio,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Update(user *u.User) (*u.User, error) {
	var res u.User
	query := `
	UPDATE
		users
	SET
		first_name=$1,
		last_name=$2,
		username=$3,
        avatar_url=$4,
        bio=$5,
		updeted_at=CURRENT_TIMESTAMP
	WHERE
		id=$5
	returning
		id, 
		first_name,
		last_name,
		username,
		avatar_url,
		bio,
		created_at,
		updeted_at
	`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Username,
        user.AvatarUrl,
        user.Bio,
        user.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.AvatarUrl,
		&res.Bio,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *userRepo) Delete(user *u.GetUserRequest) (*u.User, error) {
	var res u.User
	query := `
	UPDATE
		users
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	RETURNING
		id, 
		first_name, 
		last_name,
		username,
		avatar_url,
		bio,
		created_at,
		updeted_at,
		deleted_at
	`
	err := r.db.QueryRow(query, user.UserId).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Username,
		&res.AvatarUrl,
		&res.Bio,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}

func (r *userRepo) GetAll(user *u.GetAllRequest) (*u.GetAllResponse, error) {
	var allUser u.GetAllResponse
	query := `
	SELECT
		id,
		first_name,
		last_name,
		username,
		avatar_url,
		bio,
		created_at,
		updeted_at
	FROM 
		users 
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`
	offset := user.Limit * (user.Page - 1)
	rows, err := r.db.Query(query, user.Limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user u.User
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.AvatarUrl,
			&user.Bio,
			&user.CreatedAt,
			&user.UpdetedAt)
		if err != nil {
			return nil, err
		}
		allUser.Users = append(allUser.Users, &user)
	}
	return &allUser, nil
}
