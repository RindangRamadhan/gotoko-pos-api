package main

import "gotoko-pos-api/cmd"

// @contact.name   Developer
// @contact.email  rindangramadhan10@gmail.com

// @securityDefinitions.apikey  JWTBearer
// @in                          header
// @name                        Authorization
// @description                 Token for access api
func main() {
	cmd.Execute()
}
