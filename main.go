package main

import (
	"fmt"
	"github.com/gabe565/mnemonic-ninja/internal"
	"github.com/gabe565/mnemonic-ninja/internal/word"
)

func main() {
	db, err := internal.SetupDatabase()
	if err != nil {
		panic(err)
	}

	scanner, err := internal.GetCmudict()
	if err != nil {
		panic(err)
	}

	err = internal.ImportWords(db, scanner)
	if err != nil {
		panic(err)
	}

	for {
		fmt.Print("Enter a word to search for: ")
		var q string
		_, err := fmt.Scanln(&q)
		if err != nil {
			break
		}
		var words []word.Word
		if err := db.Where("word = ?", q).Find(&words).Error; err != nil {
			fmt.Print(err)
		}
		for _, w := range words {
			fmt.Printf("%s\n", w)
		}
	}
	fmt.Println("Exiting")
}
