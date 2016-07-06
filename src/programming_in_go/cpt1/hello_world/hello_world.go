package hello_world

import (
	"fmt"
	"os"
	"strings"
)

func Hello_world()  {
	who :="world!"
	if len(os.Args)>1{
		who=strings.Join(os.Args[1:]," ")
	}
	fmt.Println("Hello",who)
}
