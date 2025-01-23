package user

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-task/pkg/repository"
)

type Repo repository.Repo

type UserHandler struct {
	repo   Repo
	router *http.ServeMux
}

func NewUserHandler(repo Repo, router *http.ServeMux) {
	handler := &UserHandler{
		repo:   repo,
		router: router,
	}

	err := Migrate(handler.repo)
	if err != nil {
		log.WithField("error", err).Fatal("User table migration failed")
	}

	router.HandleFunc("POST /user", handler.create())
	router.HandleFunc("GET /user", handler.readAll())
	router.HandleFunc("PUT /user", handler.update())
	router.HandleFunc("DELETE /user", handler.delete())
}

func (u *UserHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := readBody(w, r)
		if err != nil {
			return
		}

		newUser := &User{}
		err = deserializeJSON(w, body, newUser)
		if err != nil {
			return
		}

		err = newUser.create(u.repo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.WithField("error", err).Error("Error creating user")
			return
		}

		w.WriteHeader(http.StatusCreated)
		log.Info("User created")
	}
}

func (u *UserHandler) readAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr := User{}
		allUsers, err := usr.readAll(u.repo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.WithField("error", err).Error("Error read all users")
			return
		}

		response, err := serializeJSON(w, allUsers)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
		if err != nil {
			log.Errorln("Error writing response", err)
			return
		}
		log.Info("Read all users")
	}
}

func (u *UserHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := readBody(w, r)
		if err != nil {
			return
		}

		updUser := &User{}
		err = deserializeJSON(w, body, updUser)
		if err != nil {
			return
		}

		err = updUser.update(u.repo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.WithField("error", err).Error("Error updating user")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		log.Info("User updated")
	}
}

func (u *UserHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := readBody(w, r)
		if err != nil {
			return
		}

		delUser := &User{}
		err = deserializeJSON(w, body, delUser)
		if err != nil {
			return
		}

		err = delUser.delete(u.repo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.WithField("error", err).Error("Error deleting user")
			return
		}

		w.WriteHeader(http.StatusNoContent)
		log.Info("User deleted")
	}
}
