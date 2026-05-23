package main

import (
	// "bufio"
	"fmt"
	"os"
)

// This program demonstrates how to use the os.FileInfo interface to get information about a file.
// func main() {
// 	f, err := os.Open("example.txt")
// 	if err != nil {
// 		// log the error
// 		panic(err)
// 	}

// 	fileInfo, err := f.Stat()
// 	if err != nil {
// 		// log the error
// 		panic(err)
// 	}

// 	fmt.Println("file name:", fileInfo.Name())
// 	fmt.Println("file or folder:", fileInfo.IsDir())
// 	fmt.Println("file size:", fileInfo.Size())
// 	fmt.Println("file mode:", fileInfo.Mode())
// 	fmt.Println("file mod time:", fileInfo.ModTime())
// }

// read file content
// func main() {
// 	f, err := os.Open("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	buf := make([]byte, 26)
// 	d, err := f.Read(buf)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for i := 0; i < len(buf); i++ {
// 		println("data", d, string(buf[i]))
// 	}

// 	fmt.Println("data", d, buf)
// }

// read file content using os.ReadFile
// func main() {
// 	data, err := os.ReadFile("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(string(data))
// }

// read folder
// func main() {
// 	dir, err:= os.Open("../")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer dir.Close()

// 	fileInfo, err := dir.Readdir(0)
// 	for _, file := range fileInfo {
// 		fmt.Println(file.Name(), file.IsDir())
// 	}
// }

// create a file and write to it
// func main() {
// 	f, err := os.Create("example2.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	// write string to file
// 	// f.WriteString("GOAT VIRAT KOHLI")

// 	// write byte slice to file
// 	bytes := []byte("GOAT VIRAT KOHLI.")

// 	f.Write(bytes)
// }

// read and write to another file (steaming fashion)
// func main() {
// 	f, err := os.Create("example2.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	source, err := os.Open("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer source.Close()

// 	distFile, err := os.Create("example2.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer distFile.Close()

// 	reader := bufio.NewReader(source)
// 	writer := bufio.NewWriter(distFile)

// 	for {
// 		b, err := reader.ReadByte()
// 		if err != nil {
// 			if err.Error() != "EOF" {
// 				panic(err)
// 			}

// 			break
// 		}

// 		e := writer.WriteByte(b)
// 		if e != nil {
// 			panic(e)
// 		}
// 	}

// 	writer.Flush()

// 	fmt.Println("written to new file successfully")
// }

// delete a file
func main() {

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")
}
