package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	//"github.com/gorilla/mux"
	"github.com/gorilla/mux/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	FileName string
	//path
}

type LUrls struct {
	LongUrl string
	//ShortUrl string
}

type SUrls struct {
	ShortUrl string
}

//func (d *Database) OpenDatabase() (*gorm.DB, error) {
func OpenDatabase() (*gorm.DB, error) {

	//log.Info("Inside OpenDatabase function")
	fmt.Println("Inside OpenDatabase function")
	db, err := gorm.Open(sqlite.Open("url.db"), &gorm.Config{})
	if err != nil {
		//logger.GetLogger().Error(fmt.Sprintf("ERROR is: %s", err))
		//l.Error("ERROR is: %s", err)
		fmt.Println("Error is:", err)
		return nil, err
	}
	//defer db.Close()
	db.Migrator().CreateTable(&LUrls{})

	return db, nil
}

func WelcomePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to Url Shortner --> Created by Saher AA Kadri!\n")
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	lurl := LUrls{}
	err := decoder.Decode(&lurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(lurl.LongUrl)
	OpenDatabase()
	fmt.Println("Tiny URL API consumption")
	/*
		if len(os.Args) != 2 {
			fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
			os.Exit(1)
		}
	*/
	baseUrl := "http://tinyurl.com/api-create.php?url="
	//urlToShorten := os.Args[1]
	urlToShorten := lurl.LongUrl
	getReqUrl := baseUrl + urlToShorten

	response, err := http.Get(getReqUrl)
	//fmt.Println(response.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/welcome-url-shortner", WelcomePage).Methods("GET")
	r.HandleFunc("/create-short-url", CreateShortUrl).Methods("POST")
	http.ListenAndServe(":8001", r)
}
