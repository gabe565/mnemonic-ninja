package internal

import (
	"bufio"
	_ "embed"
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

const ImportBatchSize = 999

func ImportWords(db *gorm.DB, cmudict string) error {
	var err error
	log.Println("Loading words")
	startTime := time.Now()
	s := bufio.NewScanner(strings.NewReader(cmudict))
	var lineCount int64
	err = db.Transaction(func(db *gorm.DB) error {
		words := make([]*word.WordModel, 0, ImportBatchSize)
		for s.Scan() {
			if err := s.Err(); err != nil {
				panic(s.Err())
			}

			line := s.Text()
			w, err := word.New(line)
			if err != nil {
				return err
			}

			words = append(words, w)
			if len(words) >= ImportBatchSize {
				err = db.Create(words).Error
				if err != nil {
					return err
				}
				words = make([]*word.WordModel, 0, ImportBatchSize)
			}
			lineCount += 1
		}
		err = db.Create(words).Error
		return err
	})
	if err != nil {
		return err
	}
	timeTaken := time.Since(startTime)

	var count int64
	err = db.Model(&word.WordModel{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count != lineCount {
		log.Fatalf("Not all words loaded sucessfully. missed %d words.\n", lineCount-count)
	}

	log.Printf("Loaded %d words in %s\n", count, timeTaken)

	return nil
}
