package main

import (
	"reflect"
	"testing"
)

func Test_task1(t *testing.T) {
	type args struct {
		input *[]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zeroes",
			args: args{
				input: &[]int{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "task input",
			args: args{
				input: &[]int{1, 2, 3, 4, 5, 6},
			},
			want: 3.5,
		},
		{
			name: "empty input",
			args: args{
				input: &[]int{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task1(tt.args.input); got != tt.want {
				t.Errorf("task1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task2(t *testing.T) {
	type args struct {
		input *[]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "task input",
			args: args{
				input: &[]int64{1, 2, 5, 15},
			},
			want: []int64{15, 5, 2, 1},
		},
		{
			name: "empty input",
			args: args{
				input: &[]int64{},
			},
			want: []int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task2(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("task2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_task3(t *testing.T) {
	type args struct {
		input *map[int]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty input",
			args: args{
				input: &map[int]string{},
			},
			want: []string{},
		},
		{
			name: "task input 1",
			args: args{
				input: &map[int]string{2: "a", 0: "b", 1: "c"},
			},
			want: []string{"b", "c", "a"},
		},
		{
			name: "task input 2",
			args: args{
				input: &map[int]string{10: "aa", 0: "bb", 500: "cc"},
			},
			want: []string{"bb", "aa", "cc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := task3(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("task3() = %v, want %v", got, tt.want)
			}
		})
	}
}
