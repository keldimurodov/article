package postgres

import (
	"log"
	p "projects/article/post-service/genproto/post"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) Create(post *p.Post) (*p.Post, error) {

	idd := uuid.NewString()
	post.Id = idd

	var res p.Post

	query := `INSERT INTO Post(
		id, 
		picture, 
		title, 
		article, 
		owner_id,
		created_at, 
        updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING 
		id, 
		picture, 
		title, 
		article,
		created_at,
		updeted_at`
	err := r.db.QueryRow(query, post.Id, post.Picture, post.Title, post.Article, post.OwnerId, post.CreatedAt, post.UpdetedAt).Scan(
		&res.Id,
		&res.Picture,
		&res.Title,
		&res.Article,
		&res.OwnerId,
		&res.CreatedAt,
		&res.UpdetedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) Get(pp *p.GetPostRequest) (*p.Post, error) {

	var res p.Post
	query := `
	UPDATE
		post
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	RETURNING
		id,
		picture, 
		title, 
		article,
		owner_id,
        created_at,
        updeted_at`
	err := r.db.QueryRow(query, pp.Id).Scan(
		&res.Id,
		&res.Picture,
		&res.Title,
		&res.Article,
		&res.OwnerId,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *postRepo) GetAll(pp *p.GetAllRequest) (*p.GetAllResponse, error) {
	var allPost p.GetAllResponse
	query := `
	SELECT
		id,
		picture,
		title,
		article,
		owner_id,
		created_at,
		updeted_at
	FROM 
		post
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2`

	offset := pp.Limit * (pp.Page - 1)

	rows, err := r.db.Query(query, pp.Limit, offset)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var pr p.Post

		err := rows.Scan(
			&pr.Id,
			&pr.Picture,
			&pr.Title,
			&pr.Article,
			&pr.OwnerId,
			&pr.CreatedAt,
			&pr.UpdetedAt)

		if err != nil {
			return nil, err
		}

		allPost.Posts = append(allPost.Posts, &pr)
	}
	return &allPost, nil
}

func (r *postRepo) Update(prr *p.Post) (*p.Post, error) {
	query := `
	UPDATE
		post
	SET
		title = $1,
		article = $2,
		updeted_at = CURRENT_TIMESTAMP
	WHERE
		id = $4
	RETURNING
	id,
	picture, 
	title, 
	article, 
	created_at,
	updeted_at`

	var respUser p.Post
	err := r.db.QueryRow(query, prr.Title, prr.Article, prr.Id).Scan(
		&respUser.Id,
		&respUser.Picture,
		&respUser.Title,
		&respUser.Article,
		&respUser.CreatedAt,
		&respUser.UpdetedAt,
	)

	if err != nil {
		log.Println("Error updating user in postgres")
		return nil, err
	}
	return &respUser, nil
}

func (r *postRepo) Delete(pr *p.GetPostRequest) (*p.Post, error) {

	var res p.Post

	query := `
	UPDATE
		post
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	RETURNING
	id,
	picture, 
    title, 
    article, 
    created_at,
	updeted_at,
	deleted_at`

	err := r.db.QueryRow(query, pr.Id).Scan(
		&res.Id,
		&res.Picture,
		&res.Title,
		&res.Article,
		&res.CreatedAt,
		&res.UpdetedAt,
		&res.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil

}
