package response_store

import "rhea/errors"

type ResponseData struct {
	Headers map[string]string
	Content []byte
}

type ResponseStore struct {
	store map[string]*ResponseData
}

func NewResponseStore() *ResponseStore {
	store := make(map[string]*ResponseData, 50)
	return &ResponseStore{store: store}
}

func (rs *ResponseStore) Add(path string, data *ResponseData) {
	rs.store[path] = data
}

func (rs *ResponseStore) GetData(path string) (*ResponseData, error) {
	if data, ok := rs.store[path]; ok {
		return data, nil
	}
	return &ResponseData{}, errors.ResponseNotInStoreError{Path: path}
}
