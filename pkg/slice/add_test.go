package slice

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	type args[T any] struct {
		src     []T
		index   int
		element T
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		want    []T
		wantErr bool
	}
	tests := []testCase[int]{
		// TODO: Add test cases.
		{
			name: "success",
			args: args[int]{
				src:     []int{1, 2, 3},
				index:   1,
				element: 4,
			},
			want:    []int{1, 4, 2, 3},
			wantErr: false,
		},
		{
			name: "failed",
			args: args[int]{
				src:     []int{1, 2, 3},
				index:   3,
				element: 4,
			},
			want:    []int{1, 2, 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.src, tt.args.index, tt.args.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
