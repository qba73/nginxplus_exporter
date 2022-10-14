package collector_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nginxinc/nginx-prometheus-exporter/collector"
)

func TestMergeLabels(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		mapA, mapB, want map[string]string
	}{
		{
			name: "base case",
			mapA: map[string]string{"a": "is here"},
			mapB: map[string]string{"b": "is here"},
			want: map[string]string{"a": "is here", "b": "is here"},
		},
		{
			name: "overwrite key case",
			mapA: map[string]string{"a": "is here"},
			mapB: map[string]string{"b": "is here", "a": "is now here"},
			want: map[string]string{"a": "is now here", "b": "is here"},
		},
		{
			name: "empty maps case",
			mapA: nil,
			mapB: nil,
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := collector.MergeLabels(tt.mapA, tt.mapB)
			if !cmp.Equal(got, tt.want) {
				t.Errorf("%s\n", cmp.Diff(got, tt.want))
			}
		})
	}
}
