package controllers

import (
	"fmt"
	"net/http"
	"wwgo/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "/Users/sshivram/go-workspace/src/wwgo/views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// Get /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a temporary response")
}
