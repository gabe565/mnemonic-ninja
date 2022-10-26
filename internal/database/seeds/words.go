package seeds

import (
	"bufio"
	"bytes"
	"compress/gzip"
	_ "embed"
	"github.com/gabe565/mnemonic-ninja/internal/database/models"
	"gorm.io/gorm"
	"log"
	"time"
)

const ImportBatchSize = 500

func SeedWords(db *gorm.DB, cmudictGz []byte) error {
	var err error

	log.Println("Loading words")
	startTime := time.Now()
	gz, err := gzip.NewReader(bytes.NewReader(cmudictGz))
	if err != nil {
		return err
	}
	defer func(gz *gzip.Reader) {
		_ = gz.Close()
	}(gz)

	s := bufio.NewScanner(gz)
	var inserted int64
	if err := db.Transaction(func(db *gorm.DB) error {
		words := make([]*models.Word, 0, ImportBatchSize)
		for s.Scan() {
			w, err := models.FromCmudict(s.Bytes())
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
				words = make([]*models.Word, 0, ImportBatchSize)
			}
		}
		if err := s.Err(); err != nil {
			return s.Err()
		}

		result := db.Create(words)
		inserted += result.RowsAffected
		return result.Error
	}); err != nil {
		return err
	}
	timeTaken := time.Since(startTime)

	log.Printf("Loaded %d words in %s\n", inserted, timeTaken)

	return nil
}
