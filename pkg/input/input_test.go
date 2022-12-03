package input

import (
	"reflect"
	"strconv"
	"testing"
)

func TestArrayMapAndSumOverNonEmptyLines(t *testing.T) {
	type args struct {
		lines []string
		fn    func(line string) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "cube",
			args: args{
				lines: []string{"1", "2", "3", "4", "5"},
				fn: func(line string) int {
					i, err := strconv.Atoi(line)

					if err != nil {
						panic(err)
					}

					return i * i * i
				},
			},
			want: 225,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayMapAndSumOverNonEmptyLines(tt.args.lines, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayMapAndSumOverNonEmptyLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args struct {
		lines     []string
		chunkSize int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test",
			args: args{
				lines:     []string{"a", "b", "c", "d"},
				chunkSize: 2,
			},
			want: [][]string{{"a", "b"}, {"c", "d"}},
		},
		{
			name: "test with newline",
			args: args{
				lines:     []string{"a", "b", "c", "d", ""},
				chunkSize: 2,
			},
			want: [][]string{{"a", "b"}, {"c", "d"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.lines, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
