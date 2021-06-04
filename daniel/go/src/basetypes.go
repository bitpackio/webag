
package main

import (
    "fmt"
)

func main() {

    // Deklaration und leere Default-Initialisierung von Basis-Datentypen
    // Die leere Default-Initialisierung erfolgt automatisch mit definierten Werten
    var i int 
    var f float32
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)

    // Überschreiben der Default-Initialisierung durch Zuweisung eines neuen Wertes
    i = 1
    f = 0.1
    b = true
    s = "*"
    fmt.Printf("%v %v %v %q\n", i, f, b, s)

    // Deklaration von Basis-Datentypen und explizite Initialisierung mit eigenen Werten
    var i1 int32 = 2
    var f1 float32 = 0.2
    var b1 bool = false
    var s1 string = "#"
    fmt.Printf("%v %v %v %q\n", i1, f1, b1, s1)

    // Die Standard Basis-Datentypen haben Wertebereiche
    // Wird der Wertebereich überschritten...
    //
    // int8      8-bit integers range (-128 to 127)
    // int16    16-bit integers range (-32768 to 32767)
    // int32    32-bit integers range (-2147483648 to 2147483647)
    // int64    64-bit integers range (-9223372036854775808 to 9223372036854775807)

    var i8 int8 = 127
    fmt.Printf("%v\n", i8)
    i8++
    fmt.Printf("%v\n", i8)
    // i8 = 128
    // constant -128 overflows int8

    var i16 int16 = 32767
    fmt.Printf("%v\n", i16)
    i16++
    fmt.Printf("%v\n", i16)
    // i16 = 32768
    // constant 32768 overflows int16

    var i32 int32 = 2147483647
    fmt.Printf("%v\n", i32)
    i32++
    fmt.Printf("%v\n", i32)
    // i32 = 2147483648
    // constant 2147483648 overflows int32

    var i64 int64 = 9223372036854775807
    fmt.Printf("%v\n", i64)
    i64++
    fmt.Printf("%v\n", i64)
    // i64 = 9223372036854775808
    // constant 9223372036854775808 overflows int64

    var ui8 uint8 = 0
    fmt.Printf("%v\n", ui8)
    ui8--
    fmt.Printf("%v\n", ui8)
    // ui8 = -1
    // constant -1 overflows uint

    var ui16 uint16 = 0
    fmt.Printf("%v\n", ui16)
    ui16--
    fmt.Printf("%v\n", ui16)
    // ui16 = -1
    // constant -1 overflows uint

    var ui32 uint32 = 0
    fmt.Printf("%v\n", ui32)
    ui32--
    fmt.Printf("%v\n", ui32)
    // ui32 = -1
    // constant -1 overflows uint

    var ui64 uint64 = 0
    fmt.Printf("%v\n", ui64)
    ui64--
    fmt.Printf("%v\n", ui64)
    // ui64 = -1
    // constant -1 overflows uint

    // Ist uint uint32 oder unit64?
    var ui uint = 0
    fmt.Printf("%v\n", ui)
    ui--
    fmt.Printf("%v\n", ui)

    // Wir versuchen zu landen. Wird es gelingen?
    //var reduceSpeedForLandingCounter byte
    //for reduceSpeedForLandingCounter = 250; reduceSpeedForLandingCounter <= 255; reduceSpeedForLandingCounter++ {
    //    fmt.Printf("%d %c\n", reduceSpeedForLandingCounter, reduceSpeedForLandingCounter)
    //}
}
