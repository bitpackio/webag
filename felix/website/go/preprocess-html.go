package main

import ("fmt"; "os"; "io"; "path/filepath"; "regexp"; "strings")



func read2string(name string) string {
	b := make([]byte, 1)
	text := make([]byte, 0)
	var n int
	
	file, err := os.Open(name)

	if err != nil {
		fmt.Println(err)
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
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
	
	return files
}


func main() {
	srcdir := "../html/srchtml/"
	outputdir := "../html/outputhtml/"
	rootdir := "../html/"
	
	os.RemoveAll(outputdir)
	os.Mkdir(outputdir, 1644)
	
	importre := regexp.MustCompile(`<import +src="(.*?)" ?\/?>(<\/import>)?\n?`)
	pathre := regexp.MustCompile(`^(.*\/)[^/]*`)

	err := filepath.Walk(srcdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		outputpath := strings.Replace(path, srcdir, outputdir, 1)
		
		if info.IsDir() {
			os.Mkdir(outputpath, 1644)
		} else {
			content := importre.ReplaceAllStringFunc(
				read2string(path),
				func(match string) string {
					importname := importre.FindStringSubmatch(match)[1]
					var importpath string

					if strings.Index(importname, "/") == 0 {
						importpath = rootdir + importname[1:]
					} else {
						importpathmatches := pathre.FindStringSubmatch(outputpath)
						if importpathmatches == nil {
							importpath = outputdir
						} else {
							importpath = importpathmatches[1]
						}
						importpath += importname
						
					}
					return read2string(importpath)
				}
			)
			
			outputfile, err := os.Create(outputpath)
			if err != nil {
				fmt.Println(err)
				return err
			}
			outputfile.WriteString(content)
			outputfile.Close()
		}
		
		return nil
	})
	
	
	fmt.Println(filelist("../"))
}
