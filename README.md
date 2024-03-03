# Rosina -- Go test helpers

## Status

- The project currently does not accept Pull Requests.
- Work in progress, API heavily unstable, not ready for use.

## Approach

For the time being, this attempts to extend https://github.com/go-quicktest/qt. This approach might or not work. Everything can break.

## Usage

```go
import (
    "testing"
    "github.com/go-quicktest/qt"
    "github.com/marco-m/rosina"
)

func TestFoo(t *testing.T) {
	err := WriteFile("the-file", "the-contents")

	qt.Assert(t, qt.IsNil(err))
	qt.Assert(t, rosina.FileContains("the-file", "contents"))
}
```

## About the name Rosina

The name `Rosina` is a humble homage to _Rosina Ferrario_, the first Italian woman pilot. She received her license in 1913. From the Wikipedia page in [Italian](https://it.wikipedia.org/wiki/Rosina_Ferrario) you can find also other languages.
