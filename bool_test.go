package haystack_learning_basic_concept

import (
	"reflect"
	"testing"
)

func TestNewBool(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want *Bool
	}{
		{
			name: "haystack bool is true",
			args: args{true},
			want: &Bool{val: true},
		},
		{
			name: "haystack bool is false",
			args: args{false},
			want: &Bool{val: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBool(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool_ToHaystackJsonValue(t *testing.T) {
	type fields struct {
		val bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "get true of haystack json value",
			fields: fields{true},
			want:   "b:true",
		},
		{
			name:   "get false of haystack json value",
			fields: fields{false},
			want:   "b:false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bool{
				val: tt.fields.val,
			}
			if got := b.ToHaystackJsonValue(); got != tt.want {
				t.Errorf("ToHaystackJsonValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
