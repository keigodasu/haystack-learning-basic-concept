package haystack_learning_basic_concept

import "encoding/json"

type Grid struct {
	Meta     map[string]string
	Entities []*Entity
}

type GridJSON struct {
	Meta map[string]string   `json:"meta"`
	Cols []Col               `json:"cols,omitempty"`
	Rows []map[string]string `json:"rows,omitempty"`
}

type Col struct {
	Name string `json:"name"`
}

func (g Grid) MarshallJSON() ([]byte, error) {
	tags := map[string]bool{}

	for _, v := range g.Entities {
		for label := range v.Tags {
			tags[label] = true
		}
	}

	cols := make([]Col, len(tags))
	rows := make([]map[string]string, len(g.Entities))

	i := 0
	for tag := range tags {
		cols[i] = Col{Name: tag}
		i++
	}

	for _, entity := range g.Entities {
		rows = append(rows, tagsToRow(entity.Tags))
	}

	gjson := GridJSON{
		Meta: nil,
		Cols: cols,
		Rows: rows,
	}

	return json.Marshal(gjson)
}

func tagsToRow(tags map[string]Value) map[string]string {
	row := map[string]string{}
	for k, v := range tags {
		row[k] = v.ToHaystackJsonValue()
	}

	return row
}
