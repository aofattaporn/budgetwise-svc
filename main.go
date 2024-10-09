package main

import (
	"github.com/goproject/cmd"
	_ "github.com/goproject/docs"
)


// @title User API by Fiber and Swagger
// @version 1.0
// @description API user management Server by Fiber | Doc by Swagger.

// @contact.name admin
// @contact.url http://subalgo.com/support
// @contact.email admin@subalgo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Excecute()
}
