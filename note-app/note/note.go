package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const storageDir = "./notes"

type Note struct {
    Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display()  {
	fmt.Printf("Title: %v  \nContent: \n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fmt.Sprintf("%s/%s.json", storageDir, fileName))

	jsonData, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 06444)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}
