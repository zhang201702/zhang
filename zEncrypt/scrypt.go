package zEncrypt

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"io"
	"log"
)

type scryptEn struct {
	Salt int
	Hash int
}

var Scrypt = &scryptEn{
	Salt: 8, Hash: 32,
}

func (se *scryptEn) EnCode(data string) string {
	salt := make([]byte, se.Salt)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("salt=%v\n", salt)
	}

	hash, err := scrypt.Key([]byte(data), salt, 1<<15, 8, 1, se.Hash)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}
