package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	r1, err := ioutil.ReadFile("/log/jenkins-log-duarte2")

	if err != nil {
		fmt.Println("File not created yet")
	} else {
		fmt.Println(string(r1))
	}

	l1 := []byte("hello\ngo\n")
	err1 := ioutil.WriteFile("/log/jenkins-log-duarte", l1, 0644)
	check(err1)

	f, err := os.OpenFile("/log/jenkins-log-duarte2", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	check(err)

	defer f.Close()

	l2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(l2)
	check(err)
	fmt.Printf("Wrote %d bytes\n", n2)

	n3, err := f.WriteString("Wrote string\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("Wrote %d bytes\n", n4)

	w.Flush()

}
