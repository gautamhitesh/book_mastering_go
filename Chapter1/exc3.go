//WAP to accept reading integers till word STOP is given as input


package main

import (
	"fmt"
	// "strconv"
	"os"
	"bufio"
)

func main(){
	var f *os.File
	f=os.Stdin
	defer f.Close()

	scanner:=bufio.NewScanner(f)
	fmt.Println("Now program will start accepting inputs. Enter STOP to stop the program.")
	for scanner.Scan() {
		n:=scanner.Text()
		if n == "STOP" {
			return
		}
	}
}

// extension of this program can be to store the numbers in one array and non STOP words in another