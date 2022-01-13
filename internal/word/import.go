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
	var inserted int64
	err = db.Transaction(func(db *gorm.DB) error {
		words := make([]*WordModel, 0, ImportBatchSize)
		for s.Scan() {
			w, err := FromCmudict(s.Text())
			if err != nil {
				return err
			}

			words = append(words, w)
			if len(words) >= ImportBatchSize {
				result := db.Create(words)
				inserted += result.RowsAffected
				if result.Error != nil {
					return result.Error
				}
				words = make([]*WordModel, 0, ImportBatchSize)
			}
		}
		if err := s.Err(); err != nil {
			return s.Err()
		}

		result := db.Create(words)
		inserted += result.RowsAffected
		return result.Error
	})
	if err != nil {
		return err
	}
	timeTaken := time.Since(startTime)

	log.Printf("Loaded %d words in %s\n", inserted, timeTaken)

	return nil
}
