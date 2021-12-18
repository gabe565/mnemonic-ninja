package internal

import (
	"bufio"
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"gorm.io/gorm"
	"log"
)

func ImportWords(db *gorm.DB, scanner *bufio.Scanner) error {
	var err error
	log.Println("Importing words")
	err = db.Transaction(func(tx *gorm.DB) error {
		words := make([]*word.Word, 0, 999)
		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				panic(scanner.Err())
			}

			line := scanner.Text()

			// Ignore comments
			if line[0] == '#' {
				continue
			}

			w := word.New(line)
			words = append(words, w)
			if len(words) >= 999 {
				db.Create(words)
				words = make([]*word.Word, 0, 999)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	log.Println("Import complete")
	return nil
}
