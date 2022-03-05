package compare

import (
	"reflect"
	"testing"
)

func TestParsePercentageValueFromString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    *Percent
		wantErr bool
	}{
		{
			name: "parse 0.4%",
			args: args{
				in: "0.4%",
			},
			want: &Percent{
				value: 0.4,
			},
			wantErr: false,
		},
		{
			name: "parse 10%",
			args: args{
				in: "10%",
			},
			want: &Percent{
				value: 10,
			},
			wantErr: false,
		},
		{
			name: "parse 10.34%",
			args: args{
				in: "10.34%",
			},
			want: &Percent{
				value: 10.34,
			},
			wantErr: false,
		},
		{
			name: "parse 55.4875%",
			args: args{
				in: "55.4875%",
			},
			want: &Percent{
				value: 55.4875,
			},
			wantErr: false,
		},
		{
			name: "parse 99.1457%",
			args: args{
				in: "99.1457%",
			},
			want: &Percent{
				value: 99.1457,
			},
			wantErr: false,
		},
		{
			name: "parse 100%",
			args: args{
				in: "100%",
			},
			want: &Percent{
				value: 100,
			},
			wantErr: false,
		},
		// error inputs
		{
			name: "parse -1000%",
			args: args{
				in: "-1000%",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse -100%",
			args: args{
				in: "-100%",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse -0.000000001%",
			args: args{
				in: "-0.000000001%",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse 100.00000001%",
			args: args{
				in: "100.00000001%",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse 101%",
			args: args{
				in: "101%",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse 1001%",
			args: args{
				in: "1001%",
			},
			want:    nil,
			wantErr: true,
		},
		// non valid string inputs (% is missing)
		{
			name: "parse -1000",
			args: args{
				in: "-1000",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse -100",
			args: args{
				in: "-100",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse -0.000000001",
			args: args{
				in: "-0.000000001",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse 100.00000001",
			args: args{
				in: "100.00000001",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse 101",
			args: args{
				in: "101",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "parse 1001",
			args: args{
				in: "1001",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePercentageValueFromString(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePercentageValueFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePercentageValueFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}
