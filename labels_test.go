package idx_parser

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParserLabels(t *testing.T) {
	filename := "t10k-labels.idx1-ubyte"
	filenameGz := "t10k-labels-idx1-ubyte.gz"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := downloadFile(filenameGz, "http://yann.lecun.com/exdb/mnist/t10k-labels-idx1-ubyte.gz")
		if err != nil {
			t.Fatal(err.Error())
		}

		err = decompress(filenameGz, filename)
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	labels, err := ReadLabels(filename)
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, 10000, labels.GetCount())
	assert.Equal(t, uint8(9), labels.GetLabel(7))
}
