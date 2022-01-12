package word

import (
	"bufio"
	_ "embed"
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
		words := make([]*WordModel, 0, ImportBatchSize)
		for s.Scan() {
			w, err := FromCmudict(s.Text())
			if err != nil {
				return err
			}

			words = append(words, w)
			if len(words) >= ImportBatchSize {
				err = db.Create(words).Error
				if err != nil {
					return err
				}
				words = make([]*WordModel, 0, ImportBatchSize)
			}
			lineCount += 1
		}
		if err := s.Err(); err != nil {
			return s.Err()
		}

		err = db.Create(words).Error
		return err
	})
	if err != nil {
		return err
	}
	timeTaken := time.Since(startTime)

	var count int64
	err = db.Model(&WordModel{}).Count(&count).Error
	if err != nil {
		return err
	}

	if count != lineCount {
		log.Fatalf("Not all words loaded sucessfully. missed %d words.\n", lineCount-count)
	}

	log.Printf("Loaded %d words in %s\n", count, timeTaken)

	return nil
}
