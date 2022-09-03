package main

import (
	"net/http"
	"time"

	"literal/internal/data"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var users data.User
		all, err := users.GetAll()
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "All users - success",
			Data:    envelope{"users": all},
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@doe.com",
			Password:  "password",
		}

		app.infoLog.Println("Adding user ...")

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.infoLog.Println("User added with id: ", id)
		newUser, _ := app.models.User.GetUserById(id)

		app.writeJSON(w, http.StatusOK, newUser)
	})

	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(1, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.Email = "admin@example.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(1, 60*time.Minute)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		user, err := app.models.User.GetUserById(2)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.Insert(*token, *user)
		if err != nil {
			app.errorLog.Println(err)
			return
		}

		payload := jsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJSON(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")
		valid, err := app.models.Token.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		var payload jsonResponse
		payload.Error = false
		payload.Data = valid
		payload.Message = "success supreme"

		app.writeJSON(w, http.StatusOK, payload)
	})

	return mux
}
