package php

import (
	"io"
)

type xorReadCloser struct {
	rc  io.ReadCloser
	key []byte
}

func newXorReadCloser(rc io.ReadCloser, key []byte) io.ReadCloser {
	x := new(xorReadCloser)
	x.rc = rc
	x.key = key
	return x
}

func (x *xorReadCloser) Read(p []byte) (n int, err error) {
	n, err = x.rc.Read(p)
	c := x.key[0]
	for i := 0; i < n; i++ {
		p[i] ^= c
	}

	return n, err
}

func (x *xorReadCloser) Close() error {
	return x.rc.Close()
}
