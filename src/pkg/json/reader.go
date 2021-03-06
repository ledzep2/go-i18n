package json

import (
	"github.com/ledzep2/go-i18n/src/pkg/msg"
	"io"
	"io/ioutil"
	"encoding/json"
)

type Reader struct {

}

func NewReader() msg.Reader {
	return &Reader{}
}

func (r *Reader) ReadMessages(rs io.ReadSeeker) ([]msg.Message, error) {
	data, err := ioutil.ReadAll(rs)
	if err != nil {
		return nil, err
	}
	m := make([]msg.Message, 0, 100)
	err = json.Unmarshal(data, &m)
	return m, err
}
