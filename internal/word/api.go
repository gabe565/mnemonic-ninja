package word

type WordResponse struct {
	Arpabet string `json:"arpabet"`
	Number  string `json:"number"`
}

func (w *WordModel) ToWordResponse() (WordResponse, error) {
	var err error
	entry := WordResponse{Arpabet: w.Arpabet}
	number, err := w.Number.Value()
	if err != nil {
		return entry, err
	}
	entry.Number = number.(string)
	return entry, err
}

type NumResponse struct {
	Arpabet string `json:"arpabet"`
	Word    string `json:"word"`
}

func (w *WordModel) ToNumResponse() (NumResponse, error) {
	var err error
	entry := NumResponse{Arpabet: w.Arpabet}
	word, err := w.Word.Value()
	if err != nil {
		return entry, err
	}
	entry.Word = word.(string)
	return entry, err
}
