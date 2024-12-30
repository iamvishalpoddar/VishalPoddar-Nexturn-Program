package repository

import (
	"BlogManagementSystem/model"
	"database/sql"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(DB *sql.DB) *BlogRepository {
	return &BlogRepository{DB}
}

func (repo *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare(`INSERT INTO blog (title, content, author) VALUES (?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Content, blog.Author)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	blog.ID = int(id)
	return blog, nil
}

func (repo *BlogRepository) GetBlog(id int) (*model.Blog, error) {
	row := repo.DB.QueryRow(`SELECT * FROM blog WHERE id = ?`, id)
	blog := &model.Blog{}
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	rows, err := repo.DB.Query(`SELECT * FROM blog`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repo *BlogRepository) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare(`UPDATE blog SET title = ?, content = ?, author = ? WHERE id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, blog.ID)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (repo *BlogRepository) DeleteBlog(id int) error {
	stmt, err := repo.DB.Prepare(`DELETE FROM blog WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
