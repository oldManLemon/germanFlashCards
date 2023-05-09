// package sqlutils

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	"github.com/oldManLemon/germanFlashCards/structs"
// )

// type Card = structs.Card

// func Insert_data(c Card) {
// 	// Open Sqlite Database file
// 	db, err := sql.Open("sqlite3", "sqlite/ger_dict.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	// // Create a new table to store data
// 	// _, err = db.Exec("CREATE TABLE ger_dict (id INTERGER PRIMARY KEY, ger_article TEXT NOT NULL, ger_word TEXT NOT NULL UNIQUE, eng_word TEXT NOT NULL);")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	//Prepare the statement
// 	statement, err := db.Prepare("INSERT INTO ger_dict (ger_article, ger_word, eng_word) VALUES (?, ?, ?);")
// 	if err != nil {
// 		// log.Fatal(err)
// 		fmt.Println(err)
// 	} else {
// 		defer statement.Close()
// 		_, err = statement.Exec(c.Article, c.WordGerman, c.WordEnglish)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 	}

// }
