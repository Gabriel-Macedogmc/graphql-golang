package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryId  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("INSERT INTO course (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryId: categoryId}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM course")

	if err != nil {
		return nil, err
	}

	var courses = []Course{}

	defer rows.Close()

	for rows.Next() {
		var id, name, description, categoryId string

		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryId: categoryId})
	}

	return courses, nil
}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM course WHERE category_id = $1", categoryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var courses = []Course{}

	for rows.Next() {
		var id, name, description, categoryID string

		if err := rows.Scan(&id, &name, &description, &categoryID); err != nil {
			return nil, err
		}

		courses = append(courses, Course{ID: id, Name: name, Description: description, CategoryId: categoryID})
	}

	return courses, nil
}
