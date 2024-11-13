package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	/// Writing
	size, err := file.WriteString("\nHello World")
	_, err2 := file.Write([]byte("\nWriting bytes in the file"))
	if err != nil {
		panic(err)
	}
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("File succesfully created! Size in bytes %d", size)
	file.Close()

	/// Reading
	file2, err3 := os.ReadFile("file.txt")
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(string(file2))

	/// Reading step by step opening the file
	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(f)
	/// 10 bytes in 10 bytes
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
