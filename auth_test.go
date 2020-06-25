package auth

import "testing"

func TestHello(t *testing.T) {
	want := "Hello world."

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestBasicAuth(t *testing.T) {
	want := true

	basicRequest := BasicRequest{"test", "test"}
	basicLocal := BasicLocal{"test", "test"}

	if got, _ := BasicAuth(basicRequest, basicLocal); got != want {
		t.Errorf("BasicAuth() = %v, want %v", got, want)
	}
}

func TestFailBasicAuth(t *testing.T) {
	want := false

	basicRequest := BasicRequest{"test1", "test"}
	basicLocal := BasicLocal{"test", "test"}

	if got, _ := BasicAuth(basicRequest, basicLocal); got != want {
		t.Errorf("BasicAuth() = %v, want %v", got, want)
	}
}
