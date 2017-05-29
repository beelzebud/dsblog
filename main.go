package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	// "github.com/NabeelSarwar/dsblog/article" //when using go build
	// "time"
	"article" // when using goapp serve
)

var agg = article.Aggregate()
var templates = template.Must(template.ParseFiles(
	"static/html/templatepost.html", "static/html/search_page.html", "static/html/mainpage.html"))

func ArticleTemplate(w http.ResponseWriter, r *http.Request, a article.Article) {
	err := templates.ExecuteTemplate(w, "template_post", a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/article/"):]
	article, err := article.LoadArticleTitle(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ArticleTemplate(w, r, article)
}

func sendEmailHandler(w http.ResponseWriter, r *http.Request) {
	sender := r.FormValue("name")
	email := r.FormValue("email")
	subject := r.FormValue("subject")
	host := "aspmx.l.google.com"
	body := r.FormValue("message")

	header := make(map[string]string)
	header["From"] = (&mail.Address{Name: sender, Address: email}).String()
	header["To"] = (&mail.Address{Name: "Nabeel Sarwar", Address: "nabeelsarwar200@gmail.com"}).String()
	header["Subject"] = subject

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body
	auth := smtp.PlainAuth("", "nabeelsarwar200@gmail.com", "password", host)
	err := smtp.SendMail(host+":25", auth, email, []string{"nabeelsarwar200@gmail.com"}, []byte(message))
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/contact.html", http.StatusFound)
}

func SearchBarHandler(w http.ResponseWriter, r *http.Request) {
	searchTerms := r.FormValue("searchbar")
	searchResults := article.NewSearchResults(searchTerms)
	fmt.Println(searchTerms)
	fmt.Println(searchResults)
	err := templates.ExecuteTemplate(w, "search_page", searchResults)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func HomePageFunc(w http.ResponseWriter, r *http.Request) {
	// load main page and print it
	err := templates.ExecuteTemplate(w, "main_page", agg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AboutPageFunc(w http.ResponseWriter, r *http.Request) {
	AboutPage, err := ioutil.ReadFile("static/html/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", AboutPage)
}

func ContactPageFunc(w http.ResponseWriter, r *http.Request) {
	ContactPage, err := ioutil.ReadFile("static/html/contact.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", ContactPage)
}

func init() {

	/**
	Title := "Hello"
	Url := "hello.html"
	Author := "Nabeel"
	Tags := []string{"pew", "miracle"}
	Date := time.Now()
	LimitedContent := "Limited Content"
	Content := []byte("This is the content full")
	art := article.Article{Title, Url, Author, Date, Tags, Content, LimitedContent}
	article.SaveJSONArticle(art)
	*/


	http.HandleFunc("/", HomePageFunc)
	http.HandleFunc("/emailme", sendEmailHandler)
	http.HandleFunc("/search", SearchBarHandler)
	http.HandleFunc("/about.html", AboutPageFunc)
	http.HandleFunc("/contact.html", ContactPageFunc)
	http.HandleFunc("/article/", articleHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//	http.ListenAndServe(":8080", nil)
}
