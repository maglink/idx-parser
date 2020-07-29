package idx_parser

import (
	"encoding/binary"
	"errors"
	"golang.org/x/image/bmp"
	"image"
	"os"
)

type IdxImages struct {
	width, height uint32
	images        []idxImage
}

func (f *IdxImages) GetType() FileType {
	return Images
}

func (f *IdxImages) GetCount() int {
	return len(f.images)
}

func (f *IdxImages) GetImage(i int) []byte {
	if i > len(f.images) - 1 {
		return nil
	}
	return f.images[i].pixels
}

func (f *IdxImages) ToImage(i int) *image.Gray {
	if i > len(f.images) - 1 {
		return nil
	}
	img := image.NewGray(image.Rect(0,0, int(f.width), int(f.height)))
	for j, pix := range f.images[i].pixels {
		img.Pix[j] = 255 - pix //reverse color
	}
	return img
}

func (f *IdxImages) SaveToFile(i int, filename string) error {
	img := f.ToImage(i)
	if img == nil {
		return errors.New("i is out of bounds")
	}

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	err = bmp.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}

type idxImage struct {
	pixels []byte
}

func readImages(r []byte) (*IdxImages, error) {
	idxFile := &IdxImages{}

	imagesCount := binary.BigEndian.Uint32(r[4:8])
	idxFile.width = binary.BigEndian.Uint32(r[8:12])
	idxFile.height = binary.BigEndian.Uint32(r[12:16])
	imageLen := idxFile.width * idxFile.height

	for i := uint32(0); i < imagesCount; i++ {
		idxFile.images = append(idxFile.images, idxImage{
			pixels: r[16+i*imageLen : 16+(i+1)*imageLen],
		})
	}

	return idxFile, nil
}