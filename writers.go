package ziparray

import (
	"compress/gzip"
	"encoding/binary"
	"os"
)

// WriteFloat64Stream writes an array of float64 values to a gzip
// stream in native format.
func WriteFloat64Stream(vec []float64, stream *gzip.Writer) error {

	for _, x := range vec {
		err := binary.Write(stream, binary.LittleEndian, x)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteFloat64 writes an array to a file in compressed binary format.
func WriteFloat64(vec []float64, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	return WriteFloat64Stream(vec, gid)
}

// WriteInt64Stream writes an array of int64 values to a gzip stream
// in native little endian format.
func WriteInt64Stream(vec []int64, stream *gzip.Writer) error {

	for _, x := range vec {
		err := binary.Write(stream, binary.LittleEndian, x)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteInt64 writes an array to a file in compressed binary format.
func WriteInt64(vec []int64, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	return WriteInt64Stream(vec, gid)
}

// WriteInt8Stream writes an array of int8 values to a gzip stream in
// native little endian format.
func WriteInt8Stream(vec []int8, stream *gzip.Writer) error {

	for _, x := range vec {
		err := binary.Write(stream, binary.LittleEndian, x)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteInt8 writes an array to a file in compressed binary format.
func WriteInt8(vec []int8, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	return WriteInt8Stream(vec, gid)
}

// WriteStringStream writes an array of strings to a gzip stream.  The
// strings are delimited by newline characters.
func WriteStringStream(vec []string, stream *gzip.Writer) error {

	for _, x := range vec {
		_, err := stream.Write([]byte(x + "\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteString writes an array of strings to a file in compressed
// binary format.
func WriteString(vec []string, fname string) error {

	fid, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer fid.Close()
	gid := gzip.NewWriter(fid)
	defer gid.Close()

	return WriteStringStream(vec, gid)
}
