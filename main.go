package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"wwgo/views"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))

}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func notfound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h2>page not found</h2>")
}

func main() {
	homeView = views.NewView("bootstrap", "/Users/sshivram/go-workspace/src/wwgo/views/home.gohtml")
	contactView = views.NewView("bootstrap", "/Users/sshivram/go-workspace/src/wwgo/views/contact.gohtml")
	faqView = views.NewView("bootstrap", "/Users/sshivram/go-workspace/src/wwgo/views/faq.gohtml")
	signupView = views.NewView("bootstrap", "/Users/sshivram/go-workspace/src/wwgo/views/signup.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup", signup)
	h := http.HandlerFunc(notfound)
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}
