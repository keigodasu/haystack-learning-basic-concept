package haystack_learning_basic_concept

type Value interface {
	MarshallJSON() ([]byte, error)
}

