package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t = template.Must(template.ParseGlob("views/*"))

func Start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	p := "by Andrej Rot"
	t.Execute(w, p)

}

func getUsersFromDB(w http.ResponseWriter, r *http.Request) {
	var (
		user  Author
		users []Author
	)
	rows, err := db.Query("SELECT * FROM authors")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Birth, user.Death)
		users = append(users, user)
		fmt.Fprintln(w, user.Id, user.Name, user.Gender, user.Birth, user.Death)

	}
	rows.Close()
	//	t.ExecuteTemplate(w, "index.html", users)

}

func getBooksFromDB(w http.ResponseWriter, r *http.Request) {
	var (
		book  Books
		books []Books
	)
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Published, book.Genere)
		books = append(books, book)
		fmt.Fprintln(w, book.Id, book.Title, book.Author, book.Published, book.Genere)

	}
	rows.Close()
	//	t.ExecuteTemplate(w, "index.html", users)

}

func search(w http.ResponseWriter, r *http.Request) {

	var (
		user Author
	)

	var (
		book Books
	)

	r.ParseForm()
	data := r.Form["search"][0]

	// fmt.Fprintf(w, data)

	rows, err := db.Query("SELECT * FROM authors WHERE full_name = ?", data)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Birth, &user.Death)
		fmt.Fprintln(w, user.Name, user.Id)

		if user.Id > 0 {
			var (
				book Books
			)
			rows, err := db.Query("SELECT * FROM books WHERE author = ?", user.Id)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				rows.Scan(&book.Id, &book.Title, &book.Author, &book.Published, book.Genere)
				fmt.Fprintln(w, book.Id, book.Title, book.Author, book.Published, book.Genere)

			}
			rows.Close()

		}
	}
	// ce ni resultov gremo pogledat knjige

	rows, erro := db.Query("SELECT * FROM books WHERE title = ?", data)
	if erro != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Genere, &book.Genere)
		fmt.Fprintln(w, book.Title, book.Author)

		if book.Author > 0 {
			var (
				user Author
			)
			rows, err := db.Query("SELECT * FROM authors WHERE id = ?", book.Author)
			if err != nil {
				fmt.Println(err)
			}
			for rows.Next() {
				rows.Scan(&user.Id, &user.Name, &user.Gender, &user.Birth, &user.Death)
				fmt.Fprintln(w, user.Id, user.Name, user.Gender, user.Birth, user.Death)

			}
			rows.Close()

		}
	}
}

func InsertNew(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/new_data.html")

	var user Author

	row, err := db.Query("SELECT * FROM authors")

	for row.Next() {
		row.Scan(&user.Id, &user.Name, user.Gender, &user.Birth, &user.Death)
	}

	r.ParseForm()
	name := r.FormValue("fname")
	gender := r.FormValue("gender")
	birth := r.FormValue("birth")
	death := r.FormValue("death")
	fmt.Println(name, gender, birth, death)

	insForm, err := db.Prepare("INSERT INTO authors(id,full_name, gender,birth,death) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(user.Id+1, name, gender, birth, death)

	t.Execute(w, "andrej")

}
