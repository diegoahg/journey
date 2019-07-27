package app

import (
	"github.com/gorilla/schema"
)

type Schema struct {
	decoder *schema.Decoder
}

func NewSchema() *Schema {
	return &Schema{}
}

func (s *Schema) NewDecoder() error {
	s.decoder = schema.NewDecoder()
	return nil
}

func (s *Schema) Decode(payload interface{}, src map[string][]string) error {
	return s.decoder.Decode(payload, src)
}
