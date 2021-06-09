package main

import ("fmt"; "os"; "io"; "path/filepath"; "regexp")



func read2string(name string) string {
	b := make([]byte, 1)
	text := make([]byte, 0)
	var n int
	
	file, err := os.Open(name)

	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei " + name + ":", err)
	}

	
	for {
		n, err = file.Read(b)
		if n == 0 && err == io.EOF {
			break
		}
		text = append(text, b[0])
	}
	
	file.Close()

	return string(text)
}

func filelist(dir string) []string {
	var files []string = make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	
	return files
}


func main (){
	text := read2string("test.html")
	re := regexp.MustCompile(`<import +src="(.*?)" *\/?>(<\/import>)?\n?`)
	text = re.ReplaceAllStringFunc(text, func (match string) string {
			filename := re.FindStringSubmatch(match)[1]
			return read2string(filename)
	})
	
	outputfile,_ := os.OpenFile("output.html", os.O_RDWR | os.O_CREATE, 0664)
	outputfile.Seek(0,0)
	outputfile.Truncate(0)
	outputfile.WriteString(text)
	outputfile.Close()
	
	fmt.Println(filelist("../"))
}
