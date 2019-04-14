package reader

import (
	"encoding/json"
	"io"

	"github.com/vovanushka/client-api/service"
)

// Reader provides reading collection data form json (in future other formats) stream
// and saving every item into it's service
type Reader interface {
	ReadStream()
}

// JSONReader provides reading json stream
type JSONReader struct {
	stream io.Reader
	s      service.Service
}

// NewJSONReader JSONReader constructor
func NewJSONReader(stream io.Reader, s service.Service) (*JSONReader, error) {
	return &JSONReader{stream, s}, nil
}

// ReadStream reading function. Calls service save function for each item during executing
func (r *JSONReader) ReadStream() error {
	dec := json.NewDecoder(r.stream)

	//Skipping first delimiter
	_, err := dec.Token()
	if err != nil {
		return err
	}

	for dec.More() {
		ob := map[string]interface{}{}
		//Setting object id from json map key
		ob["id"], err = dec.Token()
		if err != nil {
			return err
		}

		err = dec.Decode(&ob)
		if err != nil {
			return err
		}

		jsonString, _ := json.Marshal(ob)
		//Save json data to service
		r.s.Save(jsonString)
	}

	_, err = dec.Token()
	if err != nil {
		return err
	}

	return nil
}
