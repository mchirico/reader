package main

import (
	"fmt"
	"io"
)

type R struct {
	count int
}

func MyReader(r io.Reader) {
	_, err := r.Read([]byte("hello"))
	if err != nil {
		fmt.Println(err)
	}
}

func (r *R) Read(b []byte) (int, error) {
	r.count += 1
	fmt.Println("Read called")
	if r.count > 10 {
		return r.count, nil
	}
	n, err := r.Loop()

	if err != nil {
		return n, err
	}
	return len(b), nil

}

func (r *R) Loop() (int, error) {
	MyReader(r)
	return r.count, nil
}

func main() {
	r := &R{}
	r.Loop()
}
