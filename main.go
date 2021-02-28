package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	FileName string
	//path
}

type Urls struct {
	LongUrl  string
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
	db.Migrator().CreateTable(&Urls{})
	return db, nil
}

func main() {
	OpenDatabase()
	fmt.Println("Tiny URL API consumption")

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	baseUrl := "http://tinyurl.com/api-create.php?url="
	urlToShorten := os.Args[1]
	getReqUrl := baseUrl + urlToShorten

	response, err := http.Get(getReqUrl)
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
