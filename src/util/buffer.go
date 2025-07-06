package util

import (
	"io"
	"log"
)

// DiscardBuffer will discard the buffer.
func DiscardBuffer(src io.Reader) {
	if _, err := io.Copy(io.Discard, src); err != nil {
		log.Printf("ERROR discard buffer : %s", err.Error())
	}
}

// CloseBuffer will close the buffer.
func CloseBuffer(rc io.Closer) {
	if err := rc.Close(); err != nil {
		log.Printf("ERROR close buffer : %s", err.Error())
	}
}
