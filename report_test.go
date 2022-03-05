package compare

import (
	"reflect"
	"testing"
)

func TestReports_swap(t *testing.T) {
	tests := []struct {
		name string
		r    *Reports
		want *Reports
	}{
		{
			name: "swap numbers",
			r: &Reports{
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte(2); return &b }(),
					New:      4,
				},
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte(4); return &b }(),
					New:      2,
				},
			},
			want: &Reports{
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte(4); return &b }(),
					New:      2,
				},
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte(2); return &b }(),
					New:      4,
				},
			},
		},
		{
			name: "swap chars",
			r: &Reports{
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte('a'); return &b }(),
					New:      'b',
				},
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte('b'); return &b }(),
					New:      'a',
				},
			},
			want: &Reports{
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte('b'); return &b }(),
					New:      'a',
				},
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte('a'); return &b }(),
					New:      'b',
				},
			},
		},
		{
			name: "swap chars and numbers",
			r: &Reports{
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte(1); return &b }(),
					New:      'b',
				},
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte('b'); return &b }(),
					New:      1,
				},
			},
			want: &Reports{
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte('b'); return &b }(),
					New:      1,
				},
				{
					Type:     "string",
					Index:    1,
					Original: func() *byte { b := byte(1); return &b }(),
					New:      'b',
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.r
			r.swap()
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Report.swap() failed got = \n%v\n, want \n%v", r, tt.want)
			}
		})
	}
}

func TestReport_swap(t *testing.T) {
	tests := []struct {
		name     string
		inReport Report
		want     Report
	}{
		{
			name: "swap numbers",
			inReport: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte(2); return &b }(),
				New:      4,
			},
			want: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte(4); return &b }(),
				New:      2,
			},
		},
		{
			name: "swap chars",
			inReport: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte('a'); return &b }(),
				New:      'b',
			},
			want: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte('b'); return &b }(),
				New:      'a',
			},
		},
		{
			name: "swap number and chars",
			inReport: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte(1); return &b }(),
				New:      'b',
			},
			want: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte('b'); return &b }(),
				New:      1,
			},
		},
		{
			name: "swap chars and number",
			inReport: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte('b'); return &b }(),
				New:      1,
			},
			want: Report{
				Type:     "string",
				Index:    1,
				Original: func() *byte { b := byte(1); return &b }(),
				New:      'b',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.inReport
			r.swap()
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Report.swap() failed got = \n%v\n, want \n%v", r, tt.want)
			}
		})
	}
}
