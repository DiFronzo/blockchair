package main

import (
	"blockchair"
	"fmt"
)

func main() {
	c, _ := blockchair.New()
	mapOfShit := []string{"0", "1", "2", "3"}
	resp, e := c.GetBlocks("bitcoin", mapOfShit)
	if e != nil {
		fmt.Print(e)
	}
	fmt.Println(resp.Data)
}
