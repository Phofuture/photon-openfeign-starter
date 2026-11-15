package openfeign

import "testing"

func TestStringBody(t *testing.T) {
	var rb RequestBody = (&StringBody{}).AddData("hello")
	if rb.ContentType() != "text/plain" {
		t.Fatalf("unexpected content type: %s", rb.ContentType())
	}
	b, err := rb.Encode()
	if err != nil || string(b) != "hello" {
		t.Fatalf("unexpected encode result: %q, %v", b, err)
	}
}
