package cbreader

import (
	"io"
)

type CBReader struct {
	R        io.Reader
	Callback func([]byte)
}

func (cbr *CBReader) Read(b []byte) (int, error) {
	n, err := cbr.R.Read(b)
	if err != nil {
		return n, err
	}

	cbr.Callback(b[:n])

	return n, nil
}
