package haystack_learning_basic_concept

import (
	"reflect"
	"testing"
)

func TestNewStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *Str
	}{
		{
			name: "new haystack string",
			args: args{"test"},
			want: &Str{val: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStr_MarshallJSON(t *testing.T) {
	type fields struct {
		val string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name:    "marshall to haystack json",
			fields:  fields{val: "test"},
			want: []byte(`"s: test"`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Str{
				val: tt.fields.val,
			}
			got, err := s.MarshallJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshallJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshallJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}