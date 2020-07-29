# IDX FILE PARSER

Go version of idx file parser. Not completed implementation. Only for unsigned byte data types. For labels and images.

### Usage

```
go get github.com/maglink/idx-parser
```

```go
package main

import (
    "log"
    "github.com/maglink/idx-parser"
)

func main() {
    images, err := idx_parser.ReadImages("filename.idx")
    if err != nil {
        log.Fatal(err)
    }
    
    err = images.SaveToFile(0, "image.bmp")
    if err != nil {
        log.Fatal(err)
    }
    
    println("done")
}


```

### THE IDX FILE FORMAT

the IDX file format is a simple format for vectors and multidimensional matrices of various numerical types.
<br>The basic format is
  
<br>magic number
<br>size in dimension 0
<br>size in dimension 1
<br>size in dimension 2
<br>.....
<br>size in dimension N
<br>data

<br>The magic number is an integer (MSB first). The first 2 bytes are always 0.
  
<br>The third byte codes the type of the data:
<br><strong>0x08: unsigned byte</strong>
<br>0x09: signed byte
<br>0x0B: short (2 bytes)
<br>0x0C: int (4 bytes)
<br>0x0D: float (4 bytes)
<br>0x0E: double (8 bytes)
  
<br>The 4-th byte codes the number of dimensions of the vector/matrix: 1 for vectors, 2 for matrices....

<br>The sizes in each dimension are 4-byte integers (MSB first, high endian, like in most non-Intel processors).
  
<br>The data is stored like in a C array, i.e. the index in the last dimension changes the fastest.
