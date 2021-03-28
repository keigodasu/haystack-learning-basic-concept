package haystack_learning_basic_concept

import (
	"encoding/json"
	"errors"
)

type Label string

type Entity struct {
	id   string
	dis  string
	Tags map[Label]Value
}

type JsonStructure struct {
	ID   string
	Dis  string
	Tags map[Label]string
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

func (e Entity) MarshallJSON() ([]byte, error) {
	tags := map[Label]string{}

	for l, v := range e.Tags {
		tags[l] = v.ToHaystackJsonValue()
	}

	tags["id"] = e.id
	tags["dis"] = e.dis

	return json.Marshal(tags)
}
