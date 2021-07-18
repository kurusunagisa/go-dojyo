package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Create struct {
	name string `json:"name`
}

func homePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome to the HomePage!")
  fmt.Println("Endpoint Hit: homePage")
}


func createHandler(w http.ResponseWriter, r *http.Request)(int,string,string){
	var param Create
	err := json.NewDecoder(r.Body).Decode(&param)
	if err != nil {
        log.Fatal(err)
    }
	name := param.name
	b := string(rand.Intn(100000000000))
	id := rand.Intn(100000000000)
	encoded := base64.StdEncoding.EncodeToString([]byte(b))
	return id, name, encoded
}

func userCreate(fn func(w http.ResponseWriter,r *http.Request,db *sql.DB)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
		id, name, encoded := createHandler(w,r)
		sql := "INSERT INTO user VALUES(" + string(id) + "," +  name + "," + string(encoded) + ");";
		result, err := db.Exec(sql)
		log.Fatal(err)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}




func requestHandler(db *sql.DB){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user/create", userCreate(createHandler))
	//http.HandleFunc("/user/create", userGet)
	//http.HandleFunc("/user/create", userUpdate)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	rand.Seed(time.Now().UnixNano())
    db, err := sql.Open("mysql", "root:test@tcp(localhost:3306)/game?charset=utf8mb4")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

	err = db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }

	/*content, err := ioutil.ReadFile("ER-diagram.sql")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	line := string(content)

	result, err := db.Exec(line)
	log.Fatal(err)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println(result)*/
	var name string
    row, err := db.Query("SELECT username FROM user")
    if err != nil {
        log.Fatal(err)
    }
	for row.Next() {
		err := row.Scan(&name)
		if err != nil {
        	log.Fatal(err)
    	}
    	fmt.Println(name)
	}

	
	requestHandler(db)
}