package aes

import (
	"slices"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
	}{
		{
			name:      "nil",
			plaintext: nil,
		},
		{
			name:      "empty",
			plaintext: []byte{},
		},
		{
			name:      "1",
			plaintext: []byte("s"),
		},
		{
			name:      "2",
			plaintext: []byte("sa"),
		},
		{
			name:      "3",
			plaintext: []byte("sad"),
		},
		{
			name:      "4",
			plaintext: []byte("sadm"),
		},
		{
			name:      "metadata.json",
			plaintext: []byte("metadata.json"),
		},
		{
			name:      "Shell functions",
			plaintext: []byte(`Shell functions are a way to group commands for later execution using a single name for the group. They are executed just like a "regular" command. When the name of a shell function is used as a simple command name, the list of commands associated with that function name is executed. Shell functions are executed in the current shell context; no new process is created to interpret them.`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key := range slices.Values([]AESKey{
				AESKey16Bytes{},
				AESKey24Bytes{},
				AESKey32Bytes{},
			}) {
				ciphertext, err := EncryptAES(key, tt.plaintext)
				if err != nil {
					t.Fatal(err)
				}
				plaintext, err := DecryptAES(key, ciphertext)
				if err != nil {
					t.Fatal(err)
				}
				if !slices.Equal(tt.plaintext, plaintext) {
					t.Fatalf("decrypted and plaintext are different, decrypted: %v, plaintext: %v", plaintext, tt.plaintext)
				}
				t.Logf("ciphertext: %x", ciphertext)
			}
		})
	}
}

func TestEncryptDecryptBase64(t *testing.T) {
	tests := []struct {
		name      string
		plaintext []byte
	}{
		{
			name:      "nil",
			plaintext: nil,
		},
		{
			name:      "empty",
			plaintext: []byte{},
		},
		{
			name:      "1",
			plaintext: []byte("s"),
		},
		{
			name:      "2",
			plaintext: []byte("sa"),
		},
		{
			name:      "3",
			plaintext: []byte("sad"),
		},
		{
			name:      "4",
			plaintext: []byte("sadm"),
		},
		{
			name:      "metadata.json",
			plaintext: []byte("metadata.json"),
		},
		{
			name:      "Shell functions",
			plaintext: []byte(`Shell functions are a way to group commands for later execution using a single name for the group. They are executed just like a "regular" command. When the name of a shell function is used as a simple command name, the list of commands associated with that function name is executed. Shell functions are executed in the current shell context; no new process is created to interpret them.`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key := range slices.Values([]AESKey{
				AESKey16Bytes{},
				AESKey24Bytes{},
				AESKey32Bytes{},
			}) {
				ciphertext, err := EncryptAESBase64(key, tt.plaintext)
				if err != nil {
					t.Fatal(err)
				}
				plaintext, err := DecryptAESBase64(key, ciphertext)
				if err != nil {
					t.Fatal(err)
				}
				if !slices.Equal(tt.plaintext, plaintext) {
					t.Fatalf("decrypted and plaintext are different, decrypted: %v, plaintext: %v", plaintext, tt.plaintext)
				}
				t.Logf("ciphertext: %x", ciphertext)
			}
		})
	}
}
