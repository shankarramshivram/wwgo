package controllers

import (
	"fmt"
	"net/http"
	"wwgo/models"
	"wwgo/views"
)

func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:     views.NewView("bootstrap", "users/new"),
		UserService: us,
	}
}

type Users struct {
	NewView *views.View
	models.UserService
}

// This is used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// This is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	form := SignupForm{}
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := &models.User{
		Name:  form.Name,
		Email: form.Email,
	}
	if err := u.UserService.Create(user); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, user)
}
