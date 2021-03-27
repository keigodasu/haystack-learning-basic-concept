package haystack_learning_basic_concept

import (
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
			want: &Entity{"1111", "", map[Label]*Value{}},
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
		l Label
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
				l: "testLabelValue",
				v: Value{
					kind: HaystackStr,
					str:  "testTagValue",
				},
			},
			want: &Entity{"1111", "", map[Label]*Value{
				"testLabelValue": &Value{
					kind: HaystackStr,
					str:  "testTagValue",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEntity("1111")
			e.Tags = map[Label]*Value{tt.args.l: &tt.args.v}

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
		Tags map[Label]*Value
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
		id   string
		dis  string
		Tags map[Label]*Value
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "refer to dis",
			fields: fields{
				id:   "1111",
				dis:  "test",
				Tags: nil,
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Entity{
				id:   tt.fields.id,
				dis:  tt.fields.dis,
				Tags: tt.fields.Tags,
			}
			if got := e.GetDis(); got != tt.want {
				t.Errorf("GetDis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKindString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Value
	}{
		{
			name: "set string tag",
			args: args{s: "test"},
			want: &Value{
				kind: HaystackStr,
				str:  "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KindString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KindString() = %v, want %v", got, tt.want)
			}
		})
	}
}