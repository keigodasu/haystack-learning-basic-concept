package haystack_learning_basic_concept

type Str struct {
	val string
}

func NewStr(s string) *Str  {
	return &Str{val: s}
}

func (s Str) ToHaystackJsonValue() string  {
	return "s:" + s.val
}

