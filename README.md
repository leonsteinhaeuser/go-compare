# go-diff

A library that can compare data types based on different "match" criterias.

Supported match criterias:

- less than
- less than or equal
- greater than
- greater than or equal
- percentage deviation
- regex
- range
- equal
- not equal
- not empty
- contains

## Example

```go
func main() {
    val := Validation{
        MatchType: MatchTypeEqual
        ExpectedValue: "foo",
    }
    isMatch, err := val.Matches("foo")
    if err != nil {
        panic(err)
    }
    fmt.Printf("Compared %q against %q and got %v", "foo", "foo", isMatch)
}
```
