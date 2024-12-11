package utils

import (
	"bytes"

	"github.com/gocql/gocql"
)

func UUIDBinarySearch(slc *[]gocql.UUID, it gocql.UUID) (uint64, bool) {
	// UUID === [16]byte
	low := uint64(0)
	high := uint64(len(*slc) - 1)
	for low <= high {
		mid := low + (high-low)/2
		if bytes.Equal((*slc)[mid].Bytes(), it.Bytes()) { // (*slc)[mid] == it
			return mid, true
		} else if bytes.Compare((*slc)[mid].Bytes(), it.Bytes()) == -1 { // (*slc)[mid] < it
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return uint64(0), false
}
