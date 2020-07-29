package idx_parser

import "encoding/binary"

type IdxLabels struct {
	labels []uint8
}

func (f *IdxLabels) GetType() FileType {
	return Labels
}

func (f *IdxLabels) GetCount() int {
	return len(f.labels)
}

func (f *IdxLabels) GetLabel(i int) uint8 {
	if i > len(f.labels) - 1 {
		return 0
	}
	return f.labels[i]
}

func readLabels(r []byte) (*IdxLabels, error) {
	idxFile := &IdxLabels{}

	labelsCount := binary.BigEndian.Uint32(r[4:8])
	idxFile.labels = make([]uint8, labelsCount)
	for i := uint32(0); i < labelsCount; i++ {
		idxFile.labels[i] = r[8+i:][0]
	}

	return idxFile, nil
}
