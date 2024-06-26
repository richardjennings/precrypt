package precrypt

import (
	"bytes"
	"testing"
)

func TestPrecrypt_Encrypt(t *testing.T) {
	expected := []byte("hello")
	key := []byte("passphrasewhichneedstobe32bytes!")
	encrypted, err := encrypt(expected, key)
	if err != nil {
		t.Error(err)
	}
	actual, err := decrypt(encrypted, key)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(actual, expected) {
		t.Errorf("expected %s got %s", string(expected), string(actual))
	}
}
