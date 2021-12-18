package internal

import (
	"bufio"
	_ "embed"
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"gorm.io/gorm"
	"log"
	"strings"
)

func ImportWords(db *gorm.DB) error {
	var err error
	log.Println("Importing words")
	s := bufio.NewScanner(strings.NewReader(CMUdict))
	err = db.Transaction(func(tx *gorm.DB) error {
		words := make([]*word.Word, 0, 999)
		for s.Scan() {
			if err := s.Err(); err != nil {
				panic(s.Err())
			}

			line := s.Text()
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

	_ = body.Close()
	return nil
}
