package errors

import "fmt"

type ResponseNotInStoreError struct {
	Path string
}

func (r ResponseNotInStoreError) Error() string {
	return fmt.Sprintf("response for url %s not found in the store", r.Path)
}

type ParseUrlError struct {
	Url string
}

func (p ParseUrlError) Error() string {
	return fmt.Sprintf("error parsing url %s", p.Url)
}

type RoundtripError struct {
}

func (r RoundtripError) Error() string {
	return "round trip error"
}
