package idx_parser

import (
	"compress/gzip"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestParserImages(t *testing.T) {
	filename := "t10k-images.idx3-ubyte"
	filenameGz := "t10k-images-idx3-ubyte.gz"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := downloadFile(filenameGz, "http://yann.lecun.com/exdb/mnist/t10k-images-idx3-ubyte.gz")
		if err != nil {
			t.Fatal(err.Error())
		}

		err = decompress(filenameGz, filename)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	images, err := ReadImages(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, 10000, images.GetCount())

	err = images.SaveToFile(7, "test.bmp")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func decompress(from, to string) error {
	fromFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	r, err := gzip.NewReader(fromFile)
	if err != nil {
		return err
	}
	defer r.Close()


	out, err := os.Create(to)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, r)
	return nil
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
