package main

import (
	"blockchair"
	"fmt"
)

func main() {
	c, _ := blockchair.New()
	//mapOfShit := []string{"f4184fc596403b9d638783cf57adfe4c75c605f6356fbc91338530e9831e9e16", "4133a884b8c71369956f54dbc559dc0d60f27386a8c8f68e55973fc5fd9b134e"}
	//mapOfShit2 := []string{"0", "1", "2"}
	//mapOfShit3 := []string{"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa","12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX"}
	//MapOfEthBlocks := []string{"2345678", "2345679"}
	resp, e := c.GetBlock("litecoin", "ea55f38d14dbd878f99d9d981630a62a8108ed97df4c3bd94a6d2ae5bb28d57d")
	if e != nil {
		fmt.Print(e)
	}

	for i := range resp.Data{
		fmt.Println(resp.Data[i].Block.Hash)
	}
}
