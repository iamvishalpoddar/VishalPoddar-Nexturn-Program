package controller

import (
	"BlogManagementSystem/model"
	"BlogManagementSystem/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type BlogController struct {
	BlogService *services.BlogService
}

func NewBlogController(service *services.BlogService) *BlogController {
	return &BlogController{service}
}

func (controller *BlogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	createdBlog, err := controller.BlogService.CreateBlog(&blog)
	if err != nil {
		http.Error(w, "Error creating the blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdBlog); err != nil {
		http.Error(w, "Error encoding the blog: "+err.Error(), http.StatusInternalServerError)
	}
}

func (controller *BlogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	blogId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	blog, err := controller.BlogService.GetBlog(blogId)
	if err != nil {
		http.Error(w, "Error getting the blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(blog); err != nil {
		http.Error(w, "Error encoding the blog: "+err.Error(), http.StatusInternalServerError)
	}
}

func (controller *BlogController) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := controller.BlogService.GetAllBlogs()
	if err != nil {
		http.Error(w, "Error getting the blogs: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		http.Error(w, "Error encoding the blogs: "+err.Error(), http.StatusInternalServerError)
	}
}

func (controller *BlogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	blogId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var blog model.Blog
	if err := json.NewDecoder(r.Body).Decode(&blog); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	blog.ID = blogId
	updatedBlog, err := controller.BlogService.UpdateBlog(&blog)
	if err != nil {
		http.Error(w, "Error updating the blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedBlog); err != nil {
		http.Error(w, "Error encoding the blog: "+err.Error(), http.StatusInternalServerError)
	}
}

func (controller *BlogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	blogId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if err := controller.BlogService.DeleteBlog(blogId); err != nil {
		http.Error(w, "Error deleting the blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Blog deleted successfully"); err != nil {
		http.Error(w, "Error encoding the response: "+err.Error(), http.StatusInternalServerError)
	}
}
