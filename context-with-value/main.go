package main

import (
	"context"
	"fmt"
)





func main(){
	ctx:=context.WithValue(context.Background(),"username","Alex")
	ctx=context.WithValue(ctx,"password","passw")


	result:=ctx.Value("username")
	if result==nil{
		fmt.Println("user not found")
		return
	}
    
	password_result:=ctx.Value("password")
	if password_result==nil{
		fmt.Println("password not set")
		return
	}
	var username string = result.(string)
	var mypass string = password_result.(string)


	fmt.Printf("Hello %s!",username)
	fmt.Println()
	fmt.Printf("Password %s!",mypass)
}