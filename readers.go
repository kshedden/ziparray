package ziparray

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"
	"strings"
)

// ReadInt64StreamPart reads a subarray of int64 values from a gzip
// stream.  If end < 0, reads from start to the end of the stream.
func ReadInt64StreamPart(stream *gzip.Reader, start, end int) ([]int64, error) {

	vec := make([]int64, 0, 1000)
	for k := 0; (k < end) || (end < 0); k++ {
		var x int64
		err := binary.Read(stream, binary.LittleEndian, &x)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if k >= start {
			vec = append(vec, x)
		}
	}
	return vec, nil
}

// ReadInt64Part reads a subarray of int64 values from a file in
// compressed binary format.  If end < 0, reads from start to the end
// of the file.
func ReadInt64Part(fname string, start, end int) ([]int64, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fid.Close()
	gid, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	defer gid.Close()

	return ReadInt64StreamPart(gid, start, end)
}

// ReadInt64 reads an array from a file in compressed binary format.
// The entire file is read.
func ReadInt64(fname string) ([]int64, error) {

	return ReadInt64Part(fname, 0, -1)
}

// ReadFloat64StreamPart reads a subarray of float64 values from a
// gzip stream.  If end < 0, reads from start to the end of the
// stream.
func ReadFloat64StreamPart(stream *gzip.Reader, start, end int) ([]float64, error) {

	vec := make([]float64, 0, 1000)
	for k := 0; (k < end) || (end < 0); k++ {
		var x float64
		err := binary.Read(stream, binary.LittleEndian, &x)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if k >= start {
			vec = append(vec, x)
		}
	}
	return vec, nil
}

// ReadFloat64Part reads a subarray of float64 values from a file in
// compressed binary format.  If end < 0, reads from start to the end
// of the file.
func ReadFloat64Part(fname string, start, end int) ([]float64, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fid.Close()
	gid, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	defer gid.Close()

	return ReadFloat64StreamPart(gid, start, end)
}

// ReadFloat64 reads an array of flaot64 values from a file in
// compressed binary format.  The entire file is read.
func ReadFloat64(fname string) ([]float64, error) {

	return ReadFloat64Part(fname, 0, -1)
}

// ReadStringStreamPart reads a contiguous subarray of strings from a
// gzip stream.  If end < 0, reads from start to the end of the file.
func ReadStringStreamPart(stream *gzip.Reader, start, end int) ([]string, error) {

	scanner := bufio.NewScanner(stream)

	vec := make([]string, 0)
	k := -1
	for scanner.Scan() {
		k++
		if (end >= 0) && (k >= end) {
			break
		}
		line := scanner.Text()
		if k >= start {
			line = strings.TrimRight(line, "\n")
			vec = append(vec, line)
		}
	}

	return vec, nil
}

// ReadStringPart reads a contiguous subarray of strings from a file in
// compressed format.  If end < 0, reads from start to the end of the
// file.
func ReadStringPart(fname string, start, end int) ([]string, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fid.Close()
	gid, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer gid.Close()

	return ReadStringStreamPart(gid, start, end)
}

// ReadString reads an array of strings from a file in compressed
// format.
func ReadString(fname string) ([]string, error) {

	return ReadStringPart(fname, 0, -1)
}
