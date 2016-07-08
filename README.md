ziparray provides utility functions for reading and writing Go slices
as gzipped files and streams.  Currently, a limited set of primitive
slice types is supported: `[]float64`, `[]int64`, and `[]string`.

Example of writing a `[]int64` slice named `vec`:

```
err := ziparray.WriteInt64Array(vec, "file.bin.gz")
if err != nil {
 	  panic(err)
}
```

Example of reading a `[]int64` slice named `vec`:

```
vec, err := ziparray.ReadInt64Array("file.bin.gz")
if err != nil {
 	  panic(err)
}
```

The data are read and written in native binary format without any
meta-data.  The user is responsible for ensuring that the data have
the proper type.