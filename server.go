package main

import (
	"fmt"
    "serveracademia/user"
    "serveracademia/model"
    "encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

//type 

func main(){
	fmt.Println("hello pretty pretty <3 <3 <3")
	p := model.Post{
		Id:1,
		UserId:1,
		Title:"Apple",
		Content:"The great scientist newton has a story with apple",
		Comments: [] model.Comment{},
	}
	c := model.Comment{
		Id:1,
		UserId:2,
		Comment:"apple is red",
	}
	p.Comments = append(p.Comments,c)
	user.User()
	fmt.Println(p)
    data,err := json.MarshalIndent(p,"","  ")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(data))
}