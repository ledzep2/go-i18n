package json

import (
	"github.com/ledzep2/go-i18n/src/pkg/msg"
	"io"
	"encoding/json"
)

type Writer struct {

}

func NewWriter() msg.Writer {
	return &Writer{}
}

func (w *Writer) WriteMessages(iw io.Writer, m []msg.Message) error {
	json, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	_, err = iw.Write(json)
	return err
}
