package utils_test

import (
	"bytes"
	"testing"

	"ardeolib.sapions.com/utils"
	"github.com/gocql/gocql"
)

func TestUUIDBinarySearch(t *testing.T) {
	findIt := gocql.UUID{}
	lastByte := []byte{0x00, 0x05, 0x0A, 0x0F, 0x14, 0x19, 0x1E, 0x23, 0x28, 0x2D, 0x32, 0x37}
	firstBytes := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	orderedUuidSlc := []gocql.UUID{}

	for _, b := range lastByte {
		orderedUuidSlc = append(orderedUuidSlc, gocql.UUID(append(firstBytes, b)))
	}

	for _, b := range lastByte {
		findIt = gocql.UUID(append(firstBytes, b))
		pos, ok := utils.UUIDBinarySearch(&orderedUuidSlc, findIt)
		if !bytes.Equal(orderedUuidSlc[pos].Bytes(), findIt.Bytes()) && ok {
			t.Fatal("Falhou, erro de lógica")
		}

	}

	findIt = gocql.UUID(append(firstBytes, 0xFF))
	pos, ok := utils.UUIDBinarySearch(&orderedUuidSlc, findIt)
	if bytes.Equal(orderedUuidSlc[pos].Bytes(), findIt.Bytes()) && !ok {
		t.Fatal("Falhou, erro de lógica")
	}
}
