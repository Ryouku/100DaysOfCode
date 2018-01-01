package main

import "testing"

func TestHandleInputURL(t *testing.T) {
	testRawURL := "http://some.server.com"

	path, err := handleInputURL(testRawURL)
	if err != nil {
		t.Fatal(err)
	}

	if path.schema != "http" {
		t.Errorf(
			"Schema is wrong. Have: %s, want: %s.",
			path.schema,
			"http",
		)
	}

	if path.host != "some.server.com" {
		t.Errorf(
			"Hostname is wrong. Have: %s, want: %s.",
			path.host,
			"some.server.com",
		)
	}
}
