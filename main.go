// +build darwin
// +build amd64

package main

// #cgo CFLAGS: -I/usr/local/Cellar/aubio/0.4.9_1/include/
// #cgo LDFLAGS: -L/usr/local/lib/ -laubio
// #include <stdio.h>
// #include <aubio/aubio.h>
import "C"
import "os"
import "fmt"

func main() {
    hop_size := C.uint(1024)
    filename := C.CString(os.Args[1])
    source := C.new_aubio_source(filename, 0, hop_size)
    frames := C.aubio_source_get_duration(source)
    fmt.Println("Frames", frames)
    fvec := C.new_fvec(hop_size)
    read := C.uint(0)

    for ok := true; ok; ok = (read == hop_size) {
        C.aubio_source_do(source, fvec, &read)
        C.fvec_print(fvec)
        fmt.Println("Read", read)
    }

    C.del_fvec(fvec)

    C.aubio_source_close(source)
    C.del_aubio_source(source)
}
