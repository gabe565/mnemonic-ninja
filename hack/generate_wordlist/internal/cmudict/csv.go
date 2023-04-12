package cmudict

import (
	"bufio"
	"encoding/csv"
	"io"

	"github.com/gabe565/mnemonic-ninja/hack/generate_wordlist/internal/models"
)

func Csv(w io.Writer, r io.Reader) (int, error) {
	var count int
	s := bufio.NewScanner(r)

	csvw := csv.NewWriter(w)

	for s.Scan() {
		w, err := models.FromCmudict(s.Bytes())
		if err != nil {
			return count, err
		}

		if err := csvw.Write([]string{w.Word, w.Number}); err != nil {
			return count, err
		}

		count += 1
	}
	if err := s.Err(); err != nil {
		return count, err
	}

	csvw.Flush()
	if err := csvw.Error(); err != nil {
		return count, err
	}

	return count, nil
}
