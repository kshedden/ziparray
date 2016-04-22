package ziparray

import (
	"compress/gzip"
	"encoding/binary"
	"os"
)

// WriteFloat64Array writes an array to a file in compressed binary
// format.
func WriteFloat64Array(vec []float64, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	for _, x := range vec {
		err = binary.Write(gid, binary.LittleEndian, x)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteInt64Array writes an array to a file in compressed binary
// format.
func WriteInt64Array(vec []int64, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	for _, x := range vec {
		err = binary.Write(gid, binary.LittleEndian, x)
		if err != nil {
			return err
		}
	}

	return nil
}
