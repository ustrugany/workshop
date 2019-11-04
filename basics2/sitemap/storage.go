package sitemap

import (
	"encoding/json"
	"os"
)

// create your domain errors
type StorageError struct {
	Description string
	Err         error
}

// implementing error interface
func (es StorageError) Error() string {
	return es.Description
}

// interface
type Storage interface {
	Store(s Sitemap) error
}

type FileStorage struct {
	path string
}

// usually when using new constructor return pointer to struct
func NewFileStorage(path string) *FileStorage {
	return &FileStorage{path: path}
}

// method receiver and interface implementation
func (fs *FileStorage) Store(s Sitemap) error {
	file, err := os.Create(fs.path)
	if err != nil {
		return StorageError{Description: "could not create file", Err: err}
	}

	// language mechanism that puts your function call into a stack, to be called before return
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	// serializing struct into json
	body, err := json.Marshal(s)
	if err != nil {
		return StorageError{Description: "could not marshal json", Err: err}
	}

	// muting first return argument
	_, err = file.WriteString(string(body))
	if err != nil {
		return StorageError{Description: "could write into file", Err: err}
	}

	return nil
}
