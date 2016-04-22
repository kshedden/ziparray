package ziparray

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"
	"strings"
)

// ReadInt64SubArrayStream reads a subarray of int64 values from a
// gzip stream.  If end < 0, reads from start to the end of the
// stream.
func ReadInt64SubArrayStream(stream *gzip.Reader, start, end int) ([]int64, error) {

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

// ReadInt64SubArray reads a subarray of int64 values from a file in
// compressed binary format.  If end < 0, reads from start to the end
// of the file.
func ReadInt64SubArray(fname string, start, end int) ([]int64, error) {

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

	return ReadInt64SubArrayStream(gid, start, end)
}

// ReadInt64Array reads an array from a file in compressed binary
// format.  The entire file is read.
func ReadInt64Array(fname string) ([]int64, error) {

	return ReadInt64SubArray(fname, 0, -1)
}

// ReadFloat64SubArray reads a subarray of float64 values from a gzip
// stream.  If end < 0, reads from start to the end of the stream.
func ReadFloat64SubArrayStream(stream *gzip.Reader, start, end int) ([]float64, error) {

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

// ReadFloat64SubArray reads a subarray of float64 values from a file
// in compressed binary format.  If end < 0, reads from start to the
// end of the file.
func ReadFloat64SubArray(fname string, start, end int) ([]float64, error) {

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

	return ReadFloat64SubArrayStream(gid, start, end)
}

// ReadFloat64Array reads an array of flaot64 values from a file in
// compressed binary format.  The entire file is read.
func ReadFloat64Array(fname string) ([]float64, error) {

	return ReadFloat64SubArray(fname, 0, -1)
}

// ReadStringSubArrayStream reads a contiguous subarray of strings from a
// gzip stream.  If end < 0, reads from start to the end
// of the file.
func ReadStringSubArrayStream(stream *gzip.Reader, start, end int) ([]string, error) {

	scanner := bufio.NewScanner(stream)

	vec := make([]string, 0)
	k := -1
	for scanner.Scan() {
		k++
		line := scanner.Text()
		line = strings.TrimRight(line, "\n")
		if k >= start {
			vec = append(vec, line)
		}
		if (end >= 0) && (k > end) {
			break
		}
	}

	return vec, nil
}

// ReadStringSubArray reads a contiguous subarray of strings from a
// file in compressed format.  If end < 0, reads from start to the end
// of the file.
func ReadStringSubArray(fname string, start, end int) ([]string, error) {

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

	return ReadStringSubArrayStream(gid, start, end)
}

// ReadStringArray reads an array of strings from a file in compressed
// format.
func ReadStringArray(fname string) ([]string, error) {

	return ReadStringSubArray(fname, 0, -1)
}
