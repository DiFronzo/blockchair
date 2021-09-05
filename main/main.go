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
	//MapOfEthUncles := []string{"0xedc7a92c2a8aa140b0afa26db4ce8e05994a67d6fc3d736ddd77210b0ba565bb", "0x5cd50096dbb856a6d1befa6de8f9c20decb299f375154427d90761dc0b101109"}
	MapOfEthTxs := []string{"0xd628780ba231cefe6a4f6c3da3b683b16f6151dc9753afd8773d3c2d74ac10c8","0x77025c5c7ff5eeb4bb164a4be84dd49192e12086cc321199f73888830c3ecd9e"}
	resp, e := c.GetTransactionsEth("ethereum", MapOfEthTxs)
	if e != nil {
		fmt.Print(e)
	}

	for i := range resp.Data{
		fmt.Println(resp.Data[i].Transaction.Hash)
	}
}
