package exec

import (
	"reflect"
	"testing"
)

func TestFZFArgs(t *testing.T) {
	tests := []struct {
		name         string
		preselectAll bool
		want         []string
	}{
		{
			name:         "no preselect all",
			preselectAll: false,
			want:         []string{"--multi", "--bind", "alt-a:select-all,alt-d:deselect-all"},
		},
		{
			name:         "with preselect all",
			preselectAll: true,
			want:         []string{"--multi", "--bind", "start:select-all,alt-a:select-all,alt-d:deselect-all"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := fzfArgs(tt.preselectAll)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("fzfArgs(%v) = %v, want %v", tt.preselectAll, got, tt.want)
			}
		})
	}
}
