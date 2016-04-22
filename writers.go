package ziparray

import (
	"compress/gzip"
	"encoding/binary"
	"os"
)

// WriteFloat64ArrayStream writes an array of float64 values to a gzip
// stream in native format.
func WriteFloat64ArrayStream(vec []float64, stream *gzip.Writer) error {

	for _, x := range vec {
		err := binary.Write(stream, binary.LittleEndian, x)
		if err != nil {
			return err
		}
	}

	return nil
}

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

	return WriteFloat64ArrayStream(vec, gid)
}

// WriteInt64ArrayStream writes an array of int64 values to a gzip stream in
// native little endian format.
func WriteInt64ArrayStream(vec []int64, stream *gzip.Writer) error {

	for _, x := range vec {
		err := binary.Write(stream, binary.LittleEndian, x)
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

	return WriteInt64ArrayStream(vec, gid)
}

// WriteStringArrayStream writes an array of strings to a gzip stream.
func WriteStringArrayStream(vec []string, stream *gzip.Writer) error {

	for _, x := range vec {
		_, err := stream.Write([]byte(x + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteStringArray writes an array of strings to a file in compressed
// binary format.
func WriteStringArray(vec []string, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	return WriteStringArrayStream(vec, gid)
}
