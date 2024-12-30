package main

import (
	db "BlogManagementSystem/config"
	"BlogManagementSystem/controller"
	"BlogManagementSystem/middleware"
	"BlogManagementSystem/repository"
	"BlogManagementSystem/services"
	"fmt"
	"net/http"
)

func main() {
	// Initialize the database
	db.InititalizeDB()

	// Set up repository, service, and controller
	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := services.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// Define blog routes
	blogRoutes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			blogController.CreateBlog(w, r)
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			fmt.Println("id received:", id)
			if id == "" {
				blogController.GetAllBlogs(w, r)
			} else {
				blogController.GetBlog(w, r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Define parameterized routes
	paramRoutes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			blogController.UpdateBlog(w, r)
		case http.MethodDelete:
			blogController.DeleteBlog(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Set up protected routes
	protectedMux := http.NewServeMux()
	protectedMux.Handle("/blog", blogRoutes)
	protectedMux.Handle("/blog/", paramRoutes)

	protectedRoutes := middleware.AuthorizationMiddleware(db.GetDB(), protectedMux)

	// Set up logging middleware
	loggedMux := middleware.LoggingMiddleware(http.DefaultServeMux)

	// Register routes
	http.Handle("/blog", protectedRoutes)
	http.Handle("/blog/", protectedRoutes)

	// Start the server
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		fmt.Println("Server failed to start", err)
	}
}
