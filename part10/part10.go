package main

import (
	"learnEcho/part10/router"
)




func main()  {
	e := router.New()

	e.Start(":8080")

}


