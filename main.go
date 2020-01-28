package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
	// "os"
)

var db *sql.DB
var err error

type user struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
	Status    string
}

type contant struct {
	ID      int
	Email  	string
	Pesan 	string
}

type articles struct {
	ID      int
	Title  	string
	Isi 	string
	Status  string
}


func connect_db() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1)/go_db")

/* 	dbDriver := "mysql"
    dbUser := "root"
    dbPass := "root"
    dbName := "go_db"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
     */

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

//var tmpl = template.Must(template.ParseGlob("views/*"))


func routes() {
	http.HandleFunc("/", mulai)
	http.HandleFunc("/artikelbaru1", artikelbaru1)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/listuser",listuser)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/aboutus", aboutus)
	http.HandleFunc("/contact", contact)	
	http.HandleFunc("/listartikel", listartikel)
	http.HandleFunc("/editartikel", editartikel)
	http.HandleFunc("/updateartikel", updateartikel)
	http.HandleFunc("/Deleteartikel", Deleteartikel)
	http.HandleFunc("/contactus", contactus)
	http.HandleFunc("/home", home)
}

func main() {
	connect_db()
	routes()

	defer db.Close()

/* 	http.Handle("/static/", 
        http.StripPrefix("/static/", 
            http.FileServer(http.Dir("assets"))))
 */
	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}

func checkErr(w http.ResponseWriter, r *http.Request, err error) bool {
	if err != nil {

		fmt.Println(r.Host + r.URL.Path)

		http.Redirect(w, r, r.Host+r.URL.Path, 301)
		return false
	}

	return true
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		first_name, 
		last_name, 
		password,
		status 
		FROM users WHERE username=?
		`, username).
		Scan(
			&users.ID,
			&users.Username,
			&users.FirstName,
			&users.LastName,
			&users.Password,
			&users.Status,
		)
	return users
}

func mulai(w http.ResponseWriter, r *http.Request) {
	
	

	selDB, err1 := db.Query(`
	SELECT id, 
	title, 
	isi,  
	STATUS 
	FROM articles WHERE STATUS='publish' ORDER BY id ASC `)
    if err1 != nil {
        panic(err.Error())
	}

	var emp = articles{}
	var res = []articles{}

	for selDB.Next() {
        var ID int
        var Title, Isi, Status string
        err = selDB.Scan(&ID, &Title, &Isi, &Status)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = ID
        emp.Title = Title
		emp.Isi = Isi
		emp.Status = Status

		
		res = append(res, emp)
		
	}
	
	var t, err = template.ParseFiles("views/mulai111.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, res)
	//tmpl.ExecuteTemplate(w, "home", data)
	return

}

func home(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	var Data = map[string]string{
		"username": session.GetString("username"),
		"message":  "selamat Datang anda berhasil login !",
		"status": session.GetString("status"),
		"ind": session.GetString("ind"),
	}
	var t, err = template.ParseFiles("views/home111.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, Data)
	//tmpl.ExecuteTemplate(w, "home", data)
	return

}

func listuser(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		//http.Redirect(w, r, "/login", 301)
	}  
	selDB, err1 := db.Query(`
	SELECT id, 
	username, 
	first_name, 
	last_name, 
	status 
	FROM users ORDER BY id ASC `)
    if err1 != nil {
        panic(err.Error())
	}
	
    var emp = user{}
    var res = []user{}
    for selDB.Next() {
        var ID int
        var Username, First_Name, Last_Name, Status string
        err = selDB.Scan(&ID, &Username, &First_Name, &Last_Name, &Status)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = ID
        emp.Username = Username
		emp.FirstName = First_Name
		emp.LastName = Last_Name
		if Status == "1" {
			emp.Status = "admin"
		} else {
			emp.Status = "Guest"
		}
		
		res = append(res, emp)
		
	}
	
	var t, err = template.ParseFiles("views/lisatadmin222.html")
	if err != nil {
		fmt.Println(err.Error()) 
		return
	}
	t.Execute(w, res)
   //tmpl.ExecuteTemplate(w, "listAdmin", res)
   // defer db.Close()

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register111.html")
		//tmpl.ExecuteTemplate(w, "register", nil)
		return
	}

	username := r.FormValue("email")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")
	password := r.FormValue("password")

	users := QueryUser(username)

	if (user{}) == users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if len(hashedPassword) != 0 && checkErr(w, r, err) {
			stmt, err := db.Prepare("INSERT INTO users SET username=?, password=?, first_name=?, last_name=?")
			if err == nil {
				_, err := stmt.Exec(&username, &hashedPassword, &first_name, &last_name)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
	} else {
		http.Redirect(w, r, "/register", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) != 0 && checkErr(w, r, err) {
		http.Redirect(w, r, "/home", 302)

	}
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login111.html")
		//tmpl.ExecuteTemplate(w, "login", nil) 
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)
	
	

	//deskripsi dan compare password
	var password_tes = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_tes == nil {
		//login success 		
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("name", users.FirstName)
		session.Set("status", users.Status)
		session.Set("ind", users.ID)
	

		http.Redirect(w, r, "/home", 302)
	} else {
		//login failed
		http.Redirect(w, r, "/login", 302)
	}

}
func logout(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	session.Clear()
	sessions.Destroy(w, r)
	http.Redirect(w, r, "/", 302)
}



func Show(w http.ResponseWriter, r *http.Request) {
    
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query(`SELECT id, 
	username, 
	first_name, 
	last_name, 
	status 
	FROM users  WHERE id=?`, nId)
    if err != nil {
        panic(err.Error())
    }
    var emp = user{}
    for selDB.Next() {
		var ID int
        var Username, First_Name, Last_Name, Status string
        err = selDB.Scan(&ID, &Username, &First_Name, &Last_Name, &Status)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = ID
        emp.Username = Username
		emp.FirstName = First_Name
		emp.LastName = Last_Name
		
		if Status == "1" {
			emp.Status = "admin"
		} else {
			emp.Status = "Guest"
		}
	
		
		//emp.Status = Status
	}
	
	var t, err1 = template.ParseFiles("views/show.html")
	if err1 != nil {
		fmt.Println(err.Error()) 
		return
	}
	t.Execute(w, emp)
    //tmpl.ExecuteTemplate(w, "Show", emp)
    //defer db.Close()
}

func Edit(w http.ResponseWriter, r *http.Request) {
    
    nId := r.URL.Query().Get("id")
    selDB, err := db.Query(`SELECT id, 
		username, 
		first_name, 
		last_name, 
		password,
		status 
		FROM users WHERE id=?`, nId)
    if err != nil {
        panic(err.Error())
    }
    var emp = user{}
    for selDB.Next() {
		var ID int
        var Username, First_Name, Last_Name, Password ,Status string
        err = selDB.Scan(&ID, &Username, &First_Name, &Last_Name, &Password, &Status)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = ID
        emp.Username = Username
		emp.FirstName = First_Name
		emp.LastName = Last_Name
		emp.Status = Status
    }
	var t, err1 = template.ParseFiles("views/edit.html")
	if err1 != nil {
		fmt.Println(err.Error()) 
		return
	}
	t.Execute(w, emp)
}

func Update(w http.ResponseWriter, r *http.Request) {
    
    if r.Method == "POST" {
        Username := r.FormValue("Username")
        FirstName := r.FormValue("FirstName")
		ID := r.FormValue("uid")
		LastName := r.FormValue("LastName")
		Status := r.FormValue("Status")
		Password :=r.FormValue("Password")
		
		if Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
			insForm, err := db.Prepare(`
				UPDATE users SET Username =?, 
				First_Name =?, Last_Name =?, Password =?
				Status =? WHERE ID=?`)
				if err != nil {
					panic(err.Error())
				}
				insForm.Exec(Username, FirstName,LastName,hashedPassword,Status , ID)
		} else {
			insForm, err := db.Prepare(`
				UPDATE users SET Username =?, 
				First_Name =?, Last_Name =?, 
				Status =? WHERE ID=?`)
				if err != nil {
					panic(err.Error())
				}
				insForm.Exec(Username, FirstName,LastName,Status , ID)
		}
		
		
        log.Println("UPDATE: Username: " + Username + " | ID: " + ID)
    }
    
    http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
  
    emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare(`
	DELETE FROM users WHERE id=?`)
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE --> id:"+ emp)
    
    http.Redirect(w, r, "/", 301)
}

func aboutus(w http.ResponseWriter, r *http.Request) {
	var t, err1 = template.ParseFiles("views/aboutus.html")
	if err1 != nil {
		fmt.Println(err.Error()) 
		return
	}
	t.Execute(w, nil)
   
}

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/contact_admin1.html")	
		return
	}
	
	email := r.FormValue("email")
	pesan := r.FormValue("pesan")
	
	if (email != "" && pesan != "")  {
		
			stmt, err := db.Prepare("INSERT INTO contact SET email=?, pesan=?")
			if err == nil {
				_, err := stmt.Exec(&email, &pesan)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		
	} else {
		http.Redirect(w, r, "/contact", 302)
	}
   
}

func artikelbaru1(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.ServeFile(w, r, "views/beritabaru.html")
		
		return
	}

	
	title := r.FormValue("title")
	isi := r.FormValue("isi")

	
	if (title != "")  {
		
			stmt, err := db.Prepare(`INSERT INTO articles SET title=?, isi=?`)
			if err == nil {
				_, err := stmt.Exec(&title, &isi)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				http.Redirect(w, r, "/listartikel", http.StatusSeeOther)
				return
			}
		
	} else {
		http.Redirect(w, r, "/artikelbaru1", 302)
	}
   
}

func listartikel(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
	

	selDB, err1 := db.Query(`
	SELECT id, 
	title, 
	isi,  
	STATUS 
	FROM articles  ORDER BY id ASC `)
    if err1 != nil {
        panic(err.Error())
	}

	var emp = articles{}
	var res = []articles{}

	for selDB.Next() {
        var ID int
        var Title, Isi, Status string
        err = selDB.Scan(&ID, &Title, &Isi, &Status)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = ID
        emp.Title = Title
		emp.Isi = Isi
		emp.Status = Status

		
		res = append(res, emp)
		
	}
	
	var t, err = template.ParseFiles("views/listartikel.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, res)
	//tmpl.ExecuteTemplate(w, "home", data)
	return

}

func editartikel(w http.ResponseWriter, r *http.Request) {
    
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}

	nId := r.URL.Query().Get("id")
    selDB, err := db.Query(`SELECT id, 
		title, 
		isi, 
		status 
		FROM articles WHERE id=?`, nId)
    if err != nil {
        panic(err.Error())
    }
    var emp = articles{}
    for selDB.Next() {
		var id int
        var title, isi, status string
        err = selDB.Scan(&id, &title, &isi, &status)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = id
        emp.Title = title
		emp.Isi = isi
		emp.Status = status
		
		
	}
	
	var t, err1 = template.ParseFiles("views/editartikel.html")
	if err1 != nil {
		fmt.Println(err.Error()) 
		return
	}
	t.Execute(w, emp)
}

func updateartikel(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
	
    if r.Method == "POST" {
        title := r.FormValue("title")
        isi := r.FormValue("isi")
		ID := r.FormValue("uid")
		status := r.FormValue("status")
	
		
	
			insForm, err := db.Prepare(`
				UPDATE articles SET title=?, 
				isi =?, status =? 
				 WHERE ID=?`)
				if err != nil {
					panic(err.Error())
				}
				insForm.Exec(title, isi, status , ID)
		
		
		
        log.Println("UPDATE: id artikel: " + ID + " |judul: " + title)
    }
    
    http.Redirect(w, r, "/home", 301)
}

func Deleteartikel(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		http.Redirect(w, r, "/login", 301)
	}
    emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare(`
	DELETE FROM articles WHERE id=?`)
    if err != nil {
        panic(err.Error())
    }
    delForm.Exec(emp)
    log.Println("DELETE --> id:"+ emp)
    
    http.Redirect(w, r, "/home", 301)
}

func contactus(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
	}  
	selDB1, err1 := db.Query(`
	SELECT id, 
	email, 
	pesan  
	FROM contact ORDER BY id DESC `)
    if err1 != nil {
        panic(err.Error())
	}
	
    var emp = contant{}
    var res = []contant{}
    for selDB1.Next() {
        var ID int
        var email, pesan string
        err = selDB1.Scan(&ID, &email, &pesan)
        if err != nil {
            panic(err.Error())
        }
        emp.ID = ID
        emp.Email = email
		emp.Pesan = pesan
		
		log.Println(emp)
		res = append(res, emp)
		
	}
	log.Println("-------")
	log.Println(res)
	var t, err = template.ParseFiles("views/contactus.html")
	if err != nil {
		fmt.Println(err.Error()) 
		return
	}
	t.Execute(w, res)


}