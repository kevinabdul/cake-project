package data

import (
	"database/sql"
	"errors"
	"privy/internal/validator"
	"time"
)

type Cake struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Rating      float64   `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_ar"`
}

func ValidateCake(v *validator.Validator, cake *Cake) {
	v.Check(len(cake.Title) > 0, "title", "must not be empty")
	v.Check(len(cake.Title) <= 100, "title", "must be less than 100 bytes of character")

	v.Check(len(cake.Description) > 0, "description", "must not be empty")
	v.Check(len(cake.Description) <= 1000, "description", "must be less than 1000 bytes of character")

	v.Check(cake.Rating >= 0, "rating", "must be positive number")
	v.Check(cake.Rating <= 10, "rating", "must be less than 10")
}

type CakeRepository interface {
	GetAll() ([]*Cake, error)
	GetOneByID(int) (*Cake, error)
	Insert(*Cake) error
	UpdateByID(*Cake) error
	DeleteByID(int) error
}

type DefaultCakeRepo struct {
	db *sql.DB
}

func (r *DefaultCakeRepo) GetAll() ([]*Cake, error) {
	query := `select id, title, description, rating, image, created_at, updated_at
	from cakes order by rating, title`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cakes := []*Cake{}

	for rows.Next() {
		var cake Cake
		err := rows.Scan(
			&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&cake.CreatedAt,
			&cake.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		cakes = append(cakes, &cake)
	}

	return cakes, nil
}

func (r *DefaultCakeRepo) GetOneByID(id int) (*Cake, error) {
	query := `select id, title, description, rating, image, created_at, updated_at
	from cakes where id = ?`

	var cake Cake
	err := r.db.QueryRow(query, id).Scan(
		&cake.ID,
		&cake.Title,
		&cake.Description,
		&cake.Rating,
		&cake.Image,
		&cake.CreatedAt,
		&cake.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &cake, nil
}

func (r *DefaultCakeRepo) Insert(cake *Cake) error {
	query := `insert into cakes (title, description, rating, image, created_at, updated_at)
	values (?, ?, ?, ?, ?, ?)`

	now := time.Now()
	args := []interface{}{cake.Title, cake.Description, cake.Rating, cake.Image, now, now}

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultCakeRepo) UpdateByID(cake *Cake) error {
	query := `update cakes set
	title = ?,
	description = ?,
	rating = ?,
	image = ?,
	updated_at = ?
	where id = ?
	`
	args := []interface{}{cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), cake.ID}

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultCakeRepo) DeleteByID(id int) error {
	query := `
	delete from cakes where id = ?`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func NewDefaultCakeRepository(db *sql.DB) CakeRepository {
	return &DefaultCakeRepo{db: db}
}
