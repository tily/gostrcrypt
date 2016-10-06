package gostrcrypt

import (
	"testing"
)

func Test_Encrypt_Success(t *testing.T) {
	plaintext := "hello, world"
	passwordCrypt := StrCrypt{[]byte("CAFEBABECAFEBABE")}
	encrypted, err := passwordCrypt.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}
	decrypted, err := passwordCrypt.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != plaintext {
		t.Fatalf("expected: %s, actuallly: %s", plaintext, decrypted)
	}
}

func Test_Encrypt_Success_41Characters(t *testing.T) {
	plaintext := "01234567890123456789012345678901234567890"
	passwordCrypt := StrCrypt{[]byte("CAFEBABECAFEBABE")}
	encrypted, err := passwordCrypt.Encrypt(plaintext)
	if err != nil {
		t.Fatal(err)
	}
	if len(encrypted) != 114 {
		t.Fatal("encrypted text size was not 114 byte")
	}
	decrypted, err := passwordCrypt.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != plaintext {
		t.Fatalf("expected: %s, actuallly: %s", plaintext, decrypted)
	}
}

func Test_Encrypt_Error_WrongKey(t *testing.T) {
	plaintext := "hello, world"
	passwordCrypt := StrCrypt{[]byte("hello")}
	_, err := passwordCrypt.Encrypt(plaintext)

	if err == nil {
		t.Fatal(err)
	} else {
		t.Log(err)
	}
}

func Test_Decrypt_Success(t *testing.T) {
	encrypted := "326c9236eb6a436291c23533a70e5ae8e6e4d26ea04dd3fde44e5aac"
	expected := "hello, world"
	passwordCrypt := StrCrypt{[]byte("CAFEBABECAFEBABE")}
	decrypted, err := passwordCrypt.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != expected {
		t.Fatalf("expected: %s, actuallly: %s", expected, decrypted)
	}
}

func Test_Decrypt_Success_41Characters(t *testing.T) {
	encrypted := "894a8c2bd1b048bd1d41f5d07c99b3daa4634878f9168c4ebd34a7faa2a88f385393b084e1a9682e96d6c365fb55a7dae47882147cd2f6f151"
	expected := "01234567890123456789012345678901234567890"
	passwordCrypt := StrCrypt{[]byte("CAFEBABECAFEBABE")}
	decrypted, err := passwordCrypt.Decrypt(encrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted != expected {
		t.Fatalf("expected: %s, actuallly: %s", expected, decrypted)
	}
}

func Test_Decrypt_Error_WrongKey(t *testing.T) {
	encrypted := "326c9236eb6a436291c23533a70e5ae8e6e4d26ea04dd3fde44e5aac"
	passwordCrypt := StrCrypt{[]byte("hello")}
	_, err := passwordCrypt.Decrypt(encrypted)
	if err == nil {
		t.Fatal(err)
	} else {
		t.Log(err)
	}
}

func Test_Decrypt_Error_WrongTarget(t *testing.T) {
	encrypted := "hogehoge"
	passwordCrypt := StrCrypt{[]byte("CAFEBABECAFEBABE")}
	_, err := passwordCrypt.Decrypt(encrypted)
	if err == nil {
		t.Fatal(err)
	} else {
		t.Log(err)
	}
}
