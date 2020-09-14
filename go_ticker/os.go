

package main

import (
	"fmt"
	"os/exec"
)

func main(){
	c := exec.Command("/bin/sh", "-c", "sudo apt-get install prometheus-node-exporter")
	if err := c.Run(); err != nil {
		fmt.Println("Error in installation: ", err)
	}
}


