package infrastructure

import (
	"github.com/gorilla/schema"
)

type Schema struct {
	decoder *schema.Decoder
}

// NewSchema instace a *Schema struct
func NewSchema() *Schema {
	return &Schema{}
}

// NewDecoder initialice decoder variable
func (s *Schema) NewDecoder() error {
	s.decoder = schema.NewDecoder()
	return nil
}

// Decode match interface with a map input (normally in HTTP request)
func (s *Schema) Decode(payload interface{}, src map[string][]string) error {
	return s.decoder.Decode(payload, src)
}
