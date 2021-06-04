package main

import ("fmt"; . "os"; "io")




func main (){
	var Dname string
	var f *File
	var b []byte = make([]byte,1)
	var text []byte = make([]byte,0)
	var err error
	var n int
	
	fmt.Println("Dateiname:  ")

	if len(Args)>=2{
		Dname = Args[1]
	} else {
		fmt.Scanln(&Dname)
	}
	
	
	f,err = Open(Dname)
	
	if err!=nil {
		fmt.Println("Fehler beim Öffnen der Datei: ",err)
		return
	}
	
	fmt.Println("Dateiinhalt: ")
	for {
		n,err = f.Read(b)
		if n==0 && err ==  io.EOF {
			break
		}
		text = append(text,b[0])
	}
	fmt.Println(string(text))
	f.Close()
}
