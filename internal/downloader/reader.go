package download

import "io"

type reader struct {
	io.Reader
	n int
}

func (r *reader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	r.n += n
	return n, err
}
