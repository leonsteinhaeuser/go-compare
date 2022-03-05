package compare

import (
	"reflect"
	"testing"
)

func Test_bytesDifferent(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Reports
		wantErr bool
	}{
		{
			name: "a smaller than b",
			args: args{
				a: []byte{1, 3, 3, 4, 5},
				b: []byte{1, 2, 3, 4, 5, 6},
			},
			want: Reports{
				{
					Type:     "uint8",
					Index:    5,
					Original: nil,
					New:      6,
				},
			},
			wantErr: false,
		},
		{
			name: "a bigger than b",
			args: args{
				a: []byte{1, 3, 3, 4, 5, 6},
				b: []byte{1, 2, 3, 4, 5},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bytesDifferent(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("bytesDifferent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bytesDifferent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ByteDifferent(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want *Report
	}{
		{
			name: "number a=1, b=1",
			args: args{
				a: 1,
				b: 1,
			},
			want: nil,
		},
		{
			name: "number a=1, b=2",
			args: args{
				a: 1,
				b: 2,
			},
			want: &Report{
				Index:    0,
				Original: func() *byte { b := byte(1); return &b }(),
				New:      2,
			},
		},
		{
			name: "char a=x, b=x",
			args: args{
				a: 'x',
				b: 'x',
			},
			want: nil,
		},
		{
			name: "char a=x, b=y",
			args: args{
				a: 'x',
				b: 'y',
			},
			want: &Report{
				Index:    0,
				Original: func() *byte { b := byte('x'); return &b }(),
				New:      'y',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteDifferent(tt.args.a, tt.args.b)
			if tt.want != nil {
				tt.want.Type = reflect.TypeOf(tt.args.a).String()
			}
			if tt.want == got {
				// both nil
				return
			}
			if !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("ByteDifferent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_BytesDifferent(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Reports
		wantErr bool
	}{
		{
			name: "a must be smaller than b",
			args: args{
				a: []byte{1, 3, 3, 4, 5, 6},
				b: []byte{1, 2, 3, 4, 5},
			},
			want: Reports{
				{
					Type:     "uint8",
					Index:    5,
					Original: nil,
					New:      6,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BytesDifferent(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesDifferent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesDifferent() = %v, want %v", got, tt.want)
			}
		})
	}
}
