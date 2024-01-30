##### 1. Bagaimanakan dependency management dalam golang ?
Menggunakan sistem modul terdesentralisasi dengan menggunakan import dari pkg.go.dev

##### 2. Jelaskan kegunaan function fmt.Sprintln! Jelaskan perbedaannya dengan fmt.Println dan berikan contoh code lalu salin output-nya!

fmt.Println digunakan untuk menampilkan parameter yang diberikan pada fungsi tersebut pada terminal, perbedaan dengan fmt.Sprintln adalah pada fungsi fmt.Sprintln tidak menampilkan langsung pada terminal, sehingga bisa disimpan pada variabel.

```
package main

import "fmt"

func main() {
    fmt.Println("Hello, world! dengan Println")
    
    output := fmt.Sprintln("Hello, world! dengan Sprintln")
    fmt.Print(output)
}
```

##### 3. Jelaskan kegunaan function fmt.Errorf! Jelaskan perbedaannya dengan errors.New dan berikan contoh code lalu salin output-nya!
Perbedaannya ada pada module yang digunakan, fmt.Errorf dari module fmt, errors.New dari module errors.
Selain itu fungsi fmt.Errorf bisa menggunakan format string

```
package main

import "fmt"
import "errors"

func main() {
    
    err := fmt.Errorf("Terjadi error: %s", "String error")
    fmt.Println(err)
    
    err2 := errors.New("Unexpected error")
    fmt.Println(err2)

}
```

##### Output penggunaan
.

```
$ go run tugas-1.go
Hello, world! dengan Println
Hello, world! dengan Sprintln

Terjadi error: String error
Unexpected error

```
