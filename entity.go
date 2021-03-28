package haystack_learning_basic_concept

import (
	"errors"
)

type Label string

type Entity struct {
	id   string
	dis  string
	Tags map[Label]Value
}

func NewEntity(id string) *Entity {
	return &Entity{
		id:   id,
		Tags: map[Label]Value{},
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
