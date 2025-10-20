package main

import (
	"testing"
)

func MustSuccess(b *testing.B, err error) {
	if err != nil {
		b.Fatal(err)
	}
}

func MustTrue(b *testing.B, bb bool) {
	if !bb {
		b.Fatal("should be true, but got", bb)
	}
}

func MustEqualInt64(t *testing.B, a, b int64) {
	if a != b {
		t.Fatalf("%d != %d", a, b)
	}
}

func MustEqual(t *testing.B, a, b string) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}
