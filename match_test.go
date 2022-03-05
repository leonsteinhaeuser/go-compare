package compare

import (
	"reflect"
	"testing"
)

func TestValidation_Matches(t *testing.T) {
	type fields struct {
		MatchType     MatchType
		MatchValue    string
		ExpectedValue interface{}
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// ============================ smaller than
		{
			name: "smaller than int64 (number = 9)",
			fields: fields{
				MatchType:     MatchTypeLessThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(9),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "smaller than int64 (number = 10)",
			fields: fields{
				MatchType:     MatchTypeLessThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(10),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "smaller than int64 (number = 10000)",
			fields: fields{
				MatchType:     MatchTypeLessThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(100),
			},
			want:    false,
			wantErr: false,
		},
		// ============================ smaller than or equal
		{
			name: "smaller than or equal int64 (number = 9)",
			fields: fields{
				MatchType:     MatchTypeLessThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(9),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "smaller than or equal int64 (number = 10)",
			fields: fields{
				MatchType:     MatchTypeLessThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(10),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "smaller than or equal int64 (number = 10000)",
			fields: fields{
				MatchType:     MatchTypeLessThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(100),
			},
			want:    false,
			wantErr: false,
		},
		// ============================ greater than
		{
			name: "greater than int64 (number = 11)",
			fields: fields{
				MatchType:     MatchTypeGreaterThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(11),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "greater than int64 (number = 100)",
			fields: fields{
				MatchType:     MatchTypeGreaterThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(100),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "greater than int64 (number = 10)",
			fields: fields{
				MatchType:     MatchTypeGreaterThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(10),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "greater than int64 (number = 0)",
			fields: fields{
				MatchType:     MatchTypeGreaterThan,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(0),
			},
			want:    false,
			wantErr: false,
		},
		// ============================ greater than or equal
		{
			name: "greater than or equal int64 (number = 11)",
			fields: fields{
				MatchType:     MatchTypeGreaterThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(11),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "greater than or equal int64 (number = 100)",
			fields: fields{
				MatchType:     MatchTypeGreaterThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(100),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "greater than or equal int64 (number = 10)",
			fields: fields{
				MatchType:     MatchTypeGreaterThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(10),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "greater than or equal int64 (number = 0)",
			fields: fields{
				MatchType:     MatchTypeGreaterThanOrEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(0),
			},
			want:    false,
			wantErr: false,
		},
		// ============================ % offset
		{
			name: "offset in % same value",
			fields: fields{
				MatchType:     MatchTypePercentageDeviation,
				MatchValue:    "10%",
				ExpectedValue: `{"value":5,"input":"hello"}`,
			},
			args: args{
				value: `{"value":5,"input":"hello"}`,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "offset in % matchvalue = 40%",
			fields: fields{
				MatchType:     MatchTypePercentageDeviation,
				MatchValue:    "40%",
				ExpectedValue: `{"value":5,"input":"hello"}`,
			},
			args: args{
				value: `{"value":10,"input":"hello","offset":5}`,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "offset in % matchvalue = 10%",
			fields: fields{
				MatchType:     MatchTypePercentageDeviation,
				MatchValue:    "10%",
				ExpectedValue: `{"value":5,"input":"hello"}`,
			},
			args: args{
				value: `{"value":10,"input":"hello","offset":5}`,
			},
			want:    false,
			wantErr: false,
		},
		// ============================ regex
		{
			name: "regex (string = 'abc')",
			fields: fields{
				MatchType:  MatchTypeRegex,
				MatchValue: "abc",
			},
			args: args{
				value: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "regex (string = '[0-9]+')",
			fields: fields{
				MatchType:  MatchTypeRegex,
				MatchValue: "[0-9]+",
			},
			args: args{
				value: "hello world 123 becko",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "regex (string = '[a-zA-Z0-9]+')",
			fields: fields{
				MatchType:  MatchTypeRegex,
				MatchValue: "[a-zA-Z0-9]+",
			},
			args: args{
				value: "helloWorld123",
			},
			want:    true,
			wantErr: false,
		},
		// ============================ Range
		{
			name: "range (number = 10)",
			fields: fields{
				MatchType:  MatchTypeRange,
				MatchValue: "10-20",
			},
			args: args{
				value: int64(10),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "range (number = 15)",
			fields: fields{
				MatchType:  MatchTypeRange,
				MatchValue: "10-20",
			},
			args: args{
				value: int64(15),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "range (number = 20)",
			fields: fields{
				MatchType:  MatchTypeRange,
				MatchValue: "10-20",
			},
			args: args{
				value: int64(15),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "range (number = 9)",
			fields: fields{
				MatchType:  MatchTypeRange,
				MatchValue: "10-20",
			},
			args: args{
				value: int64(9),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "range (number = 21)",
			fields: fields{
				MatchType:  MatchTypeRange,
				MatchValue: "10-20",
			},
			args: args{
				value: int64(21),
			},
			want:    false,
			wantErr: false,
		},
		// ============================ Equal
		{
			name: "equal (string = 'abc')",
			fields: fields{
				MatchType:     MatchTypeEqual,
				ExpectedValue: "abc",
			},
			args: args{
				value: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "equal (string = 'abc1')",
			fields: fields{
				MatchType:     MatchTypeEqual,
				ExpectedValue: "abc",
			},
			args: args{
				value: "abc1",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "equal (int64 = '100')",
			fields: fields{
				MatchType:     MatchTypeEqual,
				ExpectedValue: int64(100),
			},
			args: args{
				value: int64(100),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "equal (int64 = '101')",
			fields: fields{
				MatchType:     MatchTypeEqual,
				ExpectedValue: int64(100),
			},
			args: args{
				value: int64(101),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "equal ([]byte = '{0,1,2,1}')",
			fields: fields{
				MatchType:     MatchTypeEqual,
				ExpectedValue: []byte{0, 1, 2, 1},
			},
			args: args{
				value: []byte{0, 1, 2, 1},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "equal ([]byte = '{0,1,2,1,2}')",
			fields: fields{
				MatchType:     MatchTypeEqual,
				ExpectedValue: []byte{0, 1, 2, 1},
			},
			args: args{
				value: []byte{0, 1, 2, 1, 2},
			},
			want:    false,
			wantErr: false,
		},
		// ============================ Not Equal
		{
			name: "not equal (string = 'abc')",
			fields: fields{
				MatchType:     MatchTypeNotEqual,
				ExpectedValue: "abc",
			},
			args: args{
				value: "abc",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "not equal (string = 'abc1')",
			fields: fields{
				MatchType:     MatchTypeNotEqual,
				ExpectedValue: "abc",
			},
			args: args{
				value: "abc1",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "not equal (int64 = '10')",
			fields: fields{
				MatchType:     MatchTypeNotEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(10),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "not equal (int64 = '11')",
			fields: fields{
				MatchType:     MatchTypeNotEqual,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(11),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "not equal ([]byte = '{0,1,2,1}')",
			fields: fields{
				MatchType:     MatchTypeNotEqual,
				ExpectedValue: []byte{0, 1, 2, 1},
			},
			args: args{
				value: []byte{0, 1, 2, 1},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "not equal ([]byte = '{0,1,2,1,2}')",
			fields: fields{
				MatchType:     MatchTypeNotEqual,
				ExpectedValue: []byte{0, 1, 2, 1},
			},
			args: args{
				value: []byte{0, 1, 2, 1, 2},
			},
			want:    true,
			wantErr: false,
		},
		// ============================ Empty
		{
			name: "empty (string = 'abc')",
			fields: fields{
				MatchType: MatchTypeEmpty,
			},
			args: args{
				value: "abc",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "empty (string = '')",
			fields: fields{
				MatchType: MatchTypeEmpty,
			},
			args: args{
				value: "",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "empty (nil)",
			fields: fields{
				MatchType: MatchTypeEmpty,
			},
			args: args{
				value: nil,
			},
			want:    true,
			wantErr: false,
		},
		// ============================ Not Empty
		{
			name: "not empty (string = 'abc')",
			fields: fields{
				MatchType: MatchTypeNotEmpty,
			},
			args: args{
				value: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "not empty (string = '')",
			fields: fields{
				MatchType: MatchTypeNotEmpty,
			},
			args: args{
				value: "",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "not empty (nil)",
			fields: fields{
				MatchType: MatchTypeNotEmpty,
			},
			args: args{
				value: nil,
			},
			want:    false,
			wantErr: false,
		},
		// ============================ Contains
		{
			name: "contains (string = 'abc')",
			fields: fields{
				MatchType:     MatchTypeContains,
				ExpectedValue: "abc",
			},
			args: args{
				value: "abc",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "contains (string = 'a')",
			fields: fields{
				MatchType:     MatchTypeContains,
				ExpectedValue: "abc",
			},
			args: args{
				value: "a",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "contains (int64 = '10')",
			fields: fields{
				MatchType:     MatchTypeContains,
				ExpectedValue: int64(10),
			},
			args: args{
				value: int64(100),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "contains (int64 = '00')",
			fields: fields{
				MatchType:     MatchTypeContains,
				ExpectedValue: int64(00),
			},
			args: args{
				value: int64(1000),
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Validation{
				MatchType:     tt.fields.MatchType,
				MatchValue:    &tt.fields.MatchValue,
				ExpectedValue: tt.fields.ExpectedValue,
			}
			got, err := d.Matches(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validation.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Validation.Matches() got = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_valueToInt64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "valid int64 (number = 1)",
			args: args{
				value: int64(1),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "valid int64 (number = 100)",
			args: args{
				value: int64(100),
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "value not a number",
			args: args{
				value: "echo",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valueToInt64(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("valueToInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("valueToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rangeParser(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [2]int64
	}{
		{
			name: "valid syntax",
			args: args{
				input: "1-10",
			},
			want: [2]int64{1, 10},
		},
		{
			name: "valid syntax with spaces",
			args: args{
				input: "1-100",
			},
			want: [2]int64{1, 100},
		},
		{
			name: "empty value",
			args: args{
				input: "",
			},
			want: [2]int64{},
		},
		{
			name: "non-numeric value",
			args: args{
				input: "sample",
			},
			want: [2]int64{},
		},
		{
			name: "invalid range",
			args: args{
				input: "1-100-1",
			},
			want: [2]int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeParser(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rangeParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
