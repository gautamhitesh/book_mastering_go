package main

import (
 "fmt"
 "errors"
)

func returnError(a, b int) error { // a and b are of int type and return type of function is error
	if a == b {
		err:= errors.New("A and B are equal")
		return err
	}else{
		return nil
	}
}

func main(){ //redeclaration warning can be ignored 
	fmt.Println("Hello")
	err:=returnError(1,2)
	if err==nil{
		fmt.Println("return Error function ended normally. A and B arent equal")
	}else{
		fmt.Println(err)
	}

	err=returnError(10,10)
	if err==nil{
		fmt.Println("Return error function ended normally. A and B are equal.")
	}else{
		fmt.Println("err")
	}

	if err.Error()=="A and B are equal"{ //err.Error converts error to string
		fmt.Println("New error variable returned")
	}
}