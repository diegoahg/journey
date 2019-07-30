package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Field string `schema:"field"`
}

func TestNewSchema(t *testing.T) {
	var tStruct TestStruct
	s := NewSchema()
	err := s.NewDecoder()
	assert.Nil(t, err)

	m := map[string][]string{"field": []string{"any"}}

	err = s.Decode(&tStruct, m)
	assert.Nil(t, err)
	assert.Equal(t, "any", tStruct.Field)
}
