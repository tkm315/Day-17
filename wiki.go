package main

//wich one ?
//fmt.printf or fmt.Fprintf ?
import (
	"fmt"
	"os"
)

func main() {
	myfile, err := os.Create("Fprintf.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer myfile.Close()
	//look at here
	fmt.Fprintf(myfile, "this is my %s file", "lovely")
	readfile, err := os.ReadFile("Fprintf.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(readfile))
}
