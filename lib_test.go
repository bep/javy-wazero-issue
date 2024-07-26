package javywazeroissue

import (
	"testing"
)

func TestRunFoo(t *testing.T) {
	s, err := RunFoo()
	if err != nil {
		t.Fatal(err)
	}
	if s != `{"foo":3,"newBar":"baz!"}` {
		t.Fatalf("unexpected output: %s", s)
	}
}
