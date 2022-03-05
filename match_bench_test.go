package compare

import (
	"testing"
)

func BenchmarkValidation_Matches_MatchTypeLessThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeLessThan,
			ExpectedValue: int64(10),
		}.Matches(int64(9))
	}
}

func BenchmarkValidation_Matches_MatchTypeLessThanOrEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeLessThanOrEqual,
			ExpectedValue: int64(10),
		}.Matches(int64(9))
	}
}

func BenchmarkValidation_Matches_MatchTypeGreaterThan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeGreaterThan,
			ExpectedValue: int64(10),
		}.Matches(int64(11))
	}
}

func BenchmarkValidation_Matches_MatchTypeGreaterThanOrEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeGreaterThanOrEqual,
			ExpectedValue: int64(10),
		}.Matches(int64(11))
	}
}

func BenchmarkValidation_Matches_MatchTypePercentageDeviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType: MatchTypePercentageDeviation,
			MatchValue: func() *string {
				str := "10%"
				return &str
			}(),
			ExpectedValue: `{"value":5,"input":"hello"}`,
		}.Matches(`{"value":4,"input":"hello"}`)
	}
}

func BenchmarkValidation_Matches_MatchTypeRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType: MatchTypeRegex,
			MatchValue: func() *string {
				str := "abc"
				return &str
			}(),
		}.Matches("abc")
	}
}

func BenchmarkValidation_Matches_MatchTypeRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType: MatchTypeRange,
			MatchValue: func() *string {
				str := "10-20"
				return &str
			}(),
		}.Matches(int64(10))
	}
}

func BenchmarkValidation_Matches_MatchTypeEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeEqual,
			ExpectedValue: "abc",
		}.Matches("abc")
	}
}

func BenchmarkValidation_Matches_MatchTypeNotEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeNotEquals,
			ExpectedValue: "abc",
		}.Matches("abcd")
	}
}

func BenchmarkValidation_Matches_MatchTypeEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeEmpty,
			ExpectedValue: "abc",
		}.Matches("")
	}
}

func BenchmarkValidation_Matches_MatchTypeNotEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType:     MatchTypeNotEmpty,
			ExpectedValue: "abc",
		}.Matches("abc")
	}
}

func BenchmarkValidation_Matches_MatchTypeContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validation{
			MatchType: MatchTypeContains,
			MatchValue: func() *string {
				str := "abc"
				return &str
			}(),
		}.Matches("abc")
	}
}
