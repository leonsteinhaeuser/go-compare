package compare

import (
	"reflect"
	"testing"
)

func TestNewPercent(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		args    args
		want    *Percent
		wantErr bool
	}{
		{
			name: "non value -0.1",
			args: args{
				value: -0.1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non value -1.0",
			args: args{
				value: -1.0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non value -10.0",
			args: args{
				value: -10.0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non value -100.0",
			args: args{
				value: -100.0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non value -1000.0",
			args: args{
				value: -1000.0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "valid value 0.0",
			args: args{
				value: 0.0,
			},
			want: &Percent{
				value: 0.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 0.00000000000000000001",
			args: args{
				value: 0.00000000000000000001,
			},
			want: &Percent{
				value: 0.00000000000000000001,
			},
			wantErr: false,
		},
		{
			name: "valid value 0.1",
			args: args{
				value: 0.1,
			},
			want: &Percent{
				value: 0.1,
			},
			wantErr: false,
		},
		{
			name: "valid value 1.0",
			args: args{
				value: 1.0,
			},
			want: &Percent{
				value: 1.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 55.0",
			args: args{
				value: 55.0,
			},
			want: &Percent{
				value: 55.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 55.5",
			args: args{
				value: 55.5,
			},
			want: &Percent{
				value: 55.5,
			},
			wantErr: false,
		},
		{
			name: "valid value 99.0",
			args: args{
				value: 99.0,
			},
			want: &Percent{
				value: 99.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 99.99",
			args: args{
				value: 99.99,
			},
			want: &Percent{
				value: 99.99,
			},
			wantErr: false,
		},
		{
			name: "valid value 100.0",
			args: args{
				value: 100.0,
			},
			want: &Percent{
				value: 100.0,
			},
			wantErr: false,
		},
		{
			name: "non valid value 100.1",
			args: args{
				value: 100.1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non valid value 101.0",
			args: args{
				value: 101.0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non valid value 1000.0",
			args: args{
				value: 1000.0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPercent(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPercent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPercent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercent_Set(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "non value -0.1",
			args: args{
				value: -0.1,
			},
			wantErr: true,
		},
		{
			name: "non value -1.0",
			args: args{
				value: -1.0,
			},
			wantErr: true,
		},
		{
			name: "non value -10.0",
			args: args{
				value: -10.0,
			},
			wantErr: true,
		},
		{
			name: "non value -100.0",
			args: args{
				value: -100.0,
			},
			wantErr: true,
		},
		{
			name: "non value -1000.0",
			args: args{
				value: -1000.0,
			},
			wantErr: true,
		},
		{
			name: "valid value 0.0",
			args: args{
				value: 0.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 0.00000000000000000001",
			args: args{
				value: 0.00000000000000000001,
			},
			wantErr: false,
		},
		{
			name: "valid value 0.1",
			args: args{
				value: 0.1,
			},
			wantErr: false,
		},
		{
			name: "valid value 1.0",
			args: args{
				value: 1.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 55.0",
			args: args{
				value: 55.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 55.5",
			args: args{
				value: 55.5,
			},
			wantErr: false,
		},
		{
			name: "valid value 99.0",
			args: args{
				value: 99.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 99.99",
			args: args{
				value: 99.99,
			},
			wantErr: false,
		},
		{
			name: "valid value 100.0",
			args: args{
				value: 100.0,
			},
			wantErr: false,
		},
		{
			name: "non valid value 100.1",
			args: args{
				value: 100.1,
			},
			wantErr: true,
		},
		{
			name: "non valid value 101.0",
			args: args{
				value: 101.0,
			},
			wantErr: true,
		},
		{
			name: "non valid value 1000.0",
			args: args{
				value: 1000.0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Percent{
				value: 0.0,
			}
			if err := p.Set(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Percent.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPercent_Get(t *testing.T) {
	type fields struct {
		value float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "non value -0.1",
			fields: fields{
				value: -0.1,
			},
			want: 0,
		},
		{
			name: "non value -1.0",
			fields: fields{
				value: -1.0,
			},
			want: 0,
		},
		{
			name: "non value -10.0",
			fields: fields{
				value: -10.0,
			},
			want: 0,
		},
		{
			name: "non value -100.0",
			fields: fields{
				value: -100.0,
			},
			want: 0,
		},
		{
			name: "non value -1000.0",
			fields: fields{
				value: -1000.0,
			},
			want: 0,
		},
		{
			name: "valid value 0.0",
			fields: fields{
				value: 0.0,
			},
			want: 0.0,
		},
		{
			name: "valid value 0.00000000000000000001",
			fields: fields{
				value: 0.00000000000000000001,
			},
			want: 0.00000000000000000001,
		},
		{
			name: "valid value 0.1",
			fields: fields{
				value: 0.1,
			},
			want: 0.1,
		},
		{
			name: "valid value 1.0",
			fields: fields{
				value: 1.0,
			},
			want: 1.0,
		},
		{
			name: "valid value 55.0",
			fields: fields{
				value: 55.0,
			},
			want: 55.0,
		},
		{
			name: "valid value 55.5",
			fields: fields{
				value: 55.5,
			},
			want: 55.5,
		},
		{
			name: "valid value 99.0",
			fields: fields{
				value: 99.0,
			},
			want: 99.0,
		},
		{
			name: "valid value 99.99",
			fields: fields{
				value: 99.99,
			},
			want: 99.99,
		},
		{
			name: "valid value 100.0",
			fields: fields{
				value: 100.0,
			},
			want: 100.0,
		},
		{
			name: "non valid value 100.1",
			fields: fields{
				value: 100.1,
			},
			want: 0,
		},
		{
			name: "non valid value 101.0",
			fields: fields{
				value: 101.0,
			},
			want: 0,
		},
		{
			name: "non valid value 1000.0",
			fields: fields{
				value: 1000.0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Percent{
				value: tt.fields.value,
			}
			if got := p.Get(); got != tt.want {
				t.Errorf("Percent.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercent_checkForRangeCondition(t *testing.T) {
	type fields struct {
		value float64
	}
	type args struct {
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Percent{
				value: tt.fields.value,
			}
			if err := p.checkForRangeCondition(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Percent.checkForRangeCondition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPercentFromFloats(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    *Percent
		wantErr bool
	}{
		{
			name: "valid value 1 and 1",
			args: args{
				a: 1,
				b: 1,
			},
			want: &Percent{
				value: 0.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 1 and 0",
			args: args{
				a: 1,
				b: 0,
			},
			want: &Percent{
				value: 0.0,
			},
			wantErr: false,
		},
		{
			name: "valid value 100 and 15",
			args: args{
				a: 100,
				b: 15,
			},
			want: &Percent{
				value: 15.0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPercentFromFloats(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPercentFromFloats() error got = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPercentFromFloats() got = %v, want %v", got, tt.want)
			}
		})
	}
}
