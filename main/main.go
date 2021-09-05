package main

import (
	"blockchair"
	"fmt"
)

func main() {
	c, _ := blockchair.New()
	//mapOfShit := []string{"f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16", "4133a884b8c71369956f54dbc559dc0d60f27386a8c8f68e55973fc5fd9b134e"}
	//mapOfShit2 := []string{"0", "1", "2"}
	mapOfShit3 := []string{"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa","12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX"}
	resp, e := c.GetAddresses("bitcoin", mapOfShit3)
	if e != nil {
		fmt.Print(e)
	}

	for i := range resp.Data.Addresses{
		fmt.Println(resp.Data.Addresses[i].Balance)
	}
}
