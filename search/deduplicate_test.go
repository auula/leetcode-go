package search

import (
	"reflect"
	"testing"
)

func TestRemoveRep(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test01", args{
			a: []string{"123", "245", "345", "345"},
		}, []string{"123", "245", "345"}},
		{"test02", args{
			a: []string{"123", "245", "345", "123", "245", "345", "123", "245", "345"},
		}, []string{"123", "245", "345"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveRep(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeRep() = %v, want %v", got, tt.want)
			}
		})
	}
}
