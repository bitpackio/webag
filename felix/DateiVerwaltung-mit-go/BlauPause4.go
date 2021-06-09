package main

import ("fmt"; . "os")


func main (){
	
	var DName string
	var f *File
	var b []byte = make ([]byte,1)
	var x,y uint64
	var xs,ys string
	
	fmt.Print("Dateiname: ")
	fmt.Scanln(&DName)
	fmt.Print("Erster uint64-Wert: ")
	fmt.Scanln(&x)
	fmt.Print("Zweiter uint64-Wert: ")
	fmt.Scanln(&y)
	
	xs = fmt.Sprint(x)
	ys = fmt.Sprint(y)
	
	f,_ = Create(DName)
	
	for i:=0; i<len(xs);i++ {
		b[0]=xs[i]
		f.Write(b)
	}
	b[0]='\n'
	f.Write(b)
	
	for i:=0; i<len(ys);i++ {
		b[0]=ys[i]
		f.Write(b)
	}
	b[0]='\n'
	f.Write(b)
	
	f.Close()
	
	
	
}
