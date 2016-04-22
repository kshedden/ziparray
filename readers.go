package ziparray

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// ReadInt64SubArray reads a subarray from a file in compressed binary
// format.
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

	vec := make([]int64, 0, 1000)
	for k := 0; k < end; k++ {
		var x int64
		err = binary.Read(gid, binary.LittleEndian, &x)
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

// ReadInt64Array reads an array from a file in compressed binary
// format.
func ReadInt64Array(fname string) ([]int64, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fid.Close()
	gid, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	vec_b, err := ioutil.ReadAll(gid)
	if err != nil {
		return nil, err
	}
	defer gid.Close()

	bb := bytes.NewBuffer(vec_b)
	vec := make([]int64, len(vec_b)/8)
	for j := 0; j < len(vec_b)/8; j++ {
		err = binary.Read(bb, binary.LittleEndian, &vec[j])
		if err != nil {
			return nil, err
		}
	}
	return vec, nil
}

// ReadFloat64SubArray reads a subarray from a file in compressed binary
// format.
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

	vec := make([]float64, 0, 1000)
	for k := 0; k < end; k++ {
		var x float64
		err = binary.Read(gid, binary.LittleEndian, &x)
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

// ReadFloat64Array reads an array from a file in compressed binary
// format.
func ReadFloat64Array(fname string) ([]float64, error) {

	fid, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fid.Close()
	gid, err := gzip.NewReader(fid)
	if err != nil {
		return nil, err
	}
	vec_b, err := ioutil.ReadAll(gid)
	if err != nil {
		return nil, err
	}
	defer gid.Close()

	bb := bytes.NewBuffer(vec_b)
	vec := make([]float64, len(vec_b)/8)
	for j := 0; j < len(vec_b)/8; j++ {
		err = binary.Read(bb, binary.LittleEndian, &vec[j])
		if err != nil {
			return nil, err
		}
	}
	return vec, nil
}

// ReadStringArray reads an array of strings from a file in compressed
// format.
func ReadStringArray(fname string) ([]string, error) {

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

	scanner := bufio.NewScanner(gid)

	vec := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\n")
		vec = append(vec, line)
	}

	return vec, nil
}
