//write a program to find the sum of all numeric command line arguments


//Sol
// loop arguments from 1
// if parsed then ok otherwise log err
// use error handling, logging

package main

import (
	"fmt" // to print
	"os" // to import arguments
	// "errors" //to throw errors
	"strconv" //convert arguments from string to float    
	"log"
	// "log/syslog"
)

func main() {
	// k:=1
	// arguments:=os.Args
	// for err!=nil{
	// 	if k>=len(){
	// 		fmt.Println("None of the arguments are numeric")
	// 		return
	// 	}
	// 	n, err := strconv.ParseFloat(arguments[k], 64)
	// 	k++
	// }
	sum:=0.0
	arguments:=os.Args
	for i:=1;i<len(arguments);i++{
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err==nil{
			sum=sum+n
			// syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, "sum of numbers")
		}else{
			fmt.Println("Not a number", arguments[i])
			log.Fatal("Not a number", arguments[i], err)
		}
	}
	fmt.Println("The sum of numbers is",sum)
	return
}

// This was edition 1
// throw an error when none of the entered numbers are numerics
// use a logging file