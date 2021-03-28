package haystack_learning_basic_concept

import "encoding/json"

type Str struct {
	val string
}

func NewStr(s string) *Str  {
	return &Str{val: s}
}

func (s *Str) MarshallJSON() ([]byte, error)  {
	return json.Marshal("s: " + s.val)
}

