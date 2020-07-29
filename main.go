package idx_parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
)

type FileType uint32

const (
	Labels FileType = 2049
	Images FileType = 2051
)

type IdxFile interface {
	GetType() FileType
}

func ReadImages(filename string) (*IdxImages, error) {
	idx, err := ReadAll(filename)
	if err != nil {
		return nil, err
	}
	images, ok := idx.(*IdxImages)
	if !ok {
		return nil, errors.New("mapping error")
	}
	return images, nil
}

func ReadLabels(filename string) (*IdxLabels, error) {
	idx, err := ReadAll(filename)
	if err != nil {
		return nil, err
	}
	labels, ok := idx.(*IdxLabels)
	if !ok {
		return nil, errors.New("mapping error")
	}
	return labels, nil
}

func ReadAll(filename string) (IdxFile, error) {
	r, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if bytes.Compare(r[0:2], []byte{0,0}) != 0 {
		return nil, errors.New("invalid start of file")
	}

	idxType := FileType(binary.BigEndian.Uint32(r[0:4]))
	switch idxType {
	case Labels:
		return readLabels(r)
	case Images:
		return readImages(r)
	}

	return nil, errors.New(fmt.Sprintf("unknown idx type = %d", idxType))
}