package main

import (
	"ecommerce-server/cmd"
	// "ecommerce-server/util"
	// "fmt"
)


func main() {
	cmd.Serve();
	// jwt, err := util.CreateJwt("my-secret", util.Payload{
	// 	Sub: 45,
	// 	FirstName: "Radid",
	// 	LastName: "Boktier",
	// 	Email: "boktier@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt)
}

