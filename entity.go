package haystack_learning_basic_concept

import "errors"

type TagKind int
type Label string

const (
	HaystackStr TagKind = iota
)

type Entity struct {
	id   string
	dis  string
	Tags map[Label]*Value
}

type Value struct {
	kind TagKind
	str  string
}

func KindString(s string) *Value {
	return &Value{
		kind: HaystackStr,
		str:  s,
	}
}

func NewEntity(id string) *Entity {
	return &Entity{
		id:   id,
		Tags: map[Label]*Value{},
	}
}

func (e *Entity) GetID() string {
	return e.id
}

func (e *Entity) SetDis(dis string) error {
	if len(dis) >= 40 || len(dis) < 0 {
		return errors.New("Dis values should be more than 0 and less than 40")
	}
	e.dis = dis

	return nil
}

func (e Entity) GetDis() string {
	return e.dis
}
