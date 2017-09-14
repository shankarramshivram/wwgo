package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"wwgo/controllers"
	"wwgo/views"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
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
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	h := http.HandlerFunc(notfound)
	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}
