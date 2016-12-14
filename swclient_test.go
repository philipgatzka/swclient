package swclient

import "testing"

// TODO: write tests...

func TestNew(t *testing.T) {
	_, err := New("user", "key", "https://shop.ware/api")
	if err != nil {
		t.Error(err)
	}
}

func TestNewShouldThrowError(t *testing.T) {
	_, err := New("", "key", "https://shop.ware/api")
	if err == nil {
		t.Error("Swclient.New() didn't throw an error on empty user!")
	}

	_, err = New("user", "", "https://shop.ware/api")
	if err == nil {
		t.Error("Swclient.New() didn't throw an error on empty key!")
	}

	_, err = New("user", "key", "")
	if err == nil {
		t.Error("Swclient.New() didn't throw an error on empty URI!")
	}

	_, err = New("", "", "")
	if err == nil {
		t.Error("Swclient.New() didn't throw an error on empty arguments!")
	}
}

func TestCerror(t *testing.T) {
	err := cerror{"test", "test", "test"}
	if err.Error() != "\ntest, test: test" {
		t.Error("cerror didn't print out the error message it was given!")
	}
}
