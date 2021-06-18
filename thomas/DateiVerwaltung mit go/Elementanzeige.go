package main

import "fmt"

type Objekt interface {}

func String (a Objekt) string {
	switch v:=a.(type) {
		case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64,bool, float32, float64, complex64, complex128:
			return fmt.Sprint (v)
		default:
			return "<Objekt ohne Anzeige>"
	}
} 
		 


func main () {
	var a Objekt
	a = 4711
	if _,ok :=a.(fmt.Stringer); ok {
		fmt.Println ("JA, 'a' erfüllt das Interface 'Stringer'!")
	} else {
		fmt.Println ("NEIN, 'a' erfüllt das Interface 'Stringer' NICHT!")
	}
	fmt.Println (String(a))
	a = 47.11
	fmt.Println (String(a))
}
	
