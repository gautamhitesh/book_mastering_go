//WAP that finds the average value of its float command line arguments

package main

import (
	"fmt"
	"os"
	// "log"
	"strconv"
)

func main(){
	arguments:=os.Args
	sum:=0.0
	count:=0.0
	for i:=1; i<len(arguments);i++ {
		number, err:=strconv.ParseFloat(arguments[i], 64)
		if err!=nil{
			// log.Fatal("Not a number", err) //stops the execution of the program
			fmt.Println("Not a number", err)
		}else{
			sum+=number
			count+=1.0
		}
	}
	average:=sum/count
	fmt.Println("Average of entered numbers is ", average)
	return
}