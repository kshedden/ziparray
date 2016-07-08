ziparray provides utility functions for reading and writing Go slices
as gzipped files and streams.  Currently, a limited set of primitive
slice types is supported: `[]float64`, `[]int64`, and `[]string`.

Example of writing a `[]int64` slice named `vec`:

```
err := WriteInt64Array(vec, "file.bin.gz")
if err != nil {
 	  panic(err)
}
```

Example of reading a `[]int64` slice named `vec`:

```
vec, err := ReadInt64Array("file.bin.gz")
if err != nil {
 	  panic(err)
}
```