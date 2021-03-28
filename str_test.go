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

