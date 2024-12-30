package services

import (
	"BlogManagementSystem/model"
	"BlogManagementSystem/repository"
)

type BlogService struct {
	BlogRepo *repository.BlogRepository
}

func NewBlogService(repo *repository.BlogRepository) *BlogService {
	return &BlogService{repo}
}

func (service *BlogService) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	return service.BlogRepo.CreateBlog(blog)
}

func (service *BlogService) GetBlog(id int) (*model.Blog, error) {
	return service.BlogRepo.GetBlog(id)
}

func (service *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return service.BlogRepo.GetAllBlogs()
}

func (service *BlogService) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	return service.BlogRepo.UpdateBlog(blog)
}

func (service *BlogService) DeleteBlog(id int) error {
	return service.BlogRepo.DeleteBlog(id)
}
