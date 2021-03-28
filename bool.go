package haystack_learning_basic_concept

type Bool struct {
	val bool
}

func NewBool(b bool) *Bool {
	return &Bool{val: b}
}

func (b Bool) ToHaystackJsonValue() string {
	var boolString string
	if b.val {
		boolString = "true"
	} else {
		boolString = "false"
	}
	return "b:" + boolString
}
