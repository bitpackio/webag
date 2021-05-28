package main

import ("fmt"; . "os")



func main () {
	var text,dateiname string
	var f *File
	var err error
	
	fmt.Print("Dateiname:  ")
	fmt.Scanln(&dateiname)
	fmt.Print("TEXT OHNE LEERZEICHEN:  ")
	fmt.Scanln(&text)
	
	f,_= Create(dateiname)
	if err!=nil {
		fmt.Println("Öffenen Datei fehlgeschlagen!")
	}
	
	b := make([]byte,1)
	
	for i:=0; i<len(text);i++ {
		b[0]=text[i]
		f.Write(b)
	}
	f.Close()
}
