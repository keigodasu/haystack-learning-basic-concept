package haystack_learning_basic_concept

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewEntity(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *Entity
	}{
		{
			name: "set uuid",
			args: args{id: "1111"},
			want: &Entity{"1111", "", map[string]Value{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEntity(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetTag(t *testing.T) {
	type args struct {
		l string
		v Value
	}
	tests := []struct {
		name string
		args args
		want *Entity
	}{
		{
			name: "set string tag value",
			args: args{
				"testLabelValue",
				&Str{
					val: "testTagValue",
				},
			},
			want: &Entity{"1111", "", map[string]Value{
				"testLabelValue": &Str{
					val: "testTagValue",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEntity("1111")
			e.Tags = map[string]Value{tt.args.l: tt.args.v}

			if !reflect.DeepEqual(e, tt.want) {
				t.Errorf("NewEntity() = %v, want %v", e, tt.want)
			}
		})
	}
}

func TestEntity_GetID(t *testing.T) {
	type fields struct {
		id   string
		dis  string
		Tags map[string]Value
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "refer to ID",
			fields: fields{id: "1111"},
			want:   "1111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{
				id:   tt.fields.id,
				dis:  tt.fields.dis,
				Tags: tt.fields.Tags,
			}
			if got := e.GetID(); got != tt.want {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_GetDis(t *testing.T) {
	type fields struct {
		ID   string
		dis  string
		Tags map[string]Value
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "refer to dis",
			fields: fields{
				ID:   "1111",
				dis:  "test",
				Tags: nil,
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entity{
				id:   tt.fields.ID,
				dis:  tt.fields.dis,
				Tags: tt.fields.Tags,
			}
			if got := e.GetDis(); got != tt.want {
				t.Errorf("GetDis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_MarshallJSON(t *testing.T) {
	type fields struct {
		ID   string
		dis  string
		Tags map[string]Value
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "marshall entity with only Str",
			fields: fields{
				ID:   "1111",
				dis:  "test",
				Tags: map[string]Value{"testTagKey": &Str{val: "testTagValue"}},
			},
			want:    []byte(`{"dis":"test","id":"1111","testTagKey":"s:testTagValue"}`),
			wantErr: false,
		},
		{
			name: "marshall entity with only Bool",
			fields: fields{
				ID:   "1111",
				dis:  "test",
				Tags: map[string]Value{"testTagKey": &Bool{val: true}},
			},
			want:    []byte(`{"dis":"test","id":"1111","testTagKey":"b:true"}`),
			wantErr: false,
		},
		{
			name: "marshall entity with all type",
			fields: fields{
				ID:   "1111",
				dis:  "test",
				Tags: map[string]Value{"testTagKey": &Bool{val: true}, "testTagKeyStr": &Str{val: "testTagValue"}},
			},
			want:    []byte(`{"dis":"test","id":"1111","testTagKey":"b:true","testTagKeyStr":"s:testTagValue"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entity{
				id:   tt.fields.ID,
				dis:  tt.fields.dis,
				Tags: tt.fields.Tags,
			}
			got, err := e.MarshallJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshallJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				fmt.Println(string(got))
				fmt.Println(string(tt.want))
				fmt.Println(e)
				t.Errorf("MarshallJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
