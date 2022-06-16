package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	flagContractAddress = flag.String("contract.address", "", "Contract Address for the collection")
	flagTokenCount      = flag.Int("token.count", 10000, "# of NFTs in this collection")
	flagTokenType       = flag.String("token.type", "erc721", "erc721 or erc1511")
	flagAPIKey          = flag.String("api.key", "", "API Key")
)

func main() {

	flag.Parse()

	if v := *flagContractAddress; v != "" {
	} else {
		log.Fatal("ERROR: missing the Contract Address")
	}
	if v := *flagTokenCount; v != 0 {
	} else {
		log.Fatal("ERROR: missing Token Count")
	}
	if v := *flagTokenType; v != "" {
	} else {
		log.Fatal("ERROR: missing the Token Type")
	}
	if v := *flagAPIKey; v != "" {
	} else {
		log.Fatal("ERROR: missing the API Key")
	}

	url := "https://eth-mainnet.alchemyapi.io/v2/" + *flagAPIKey + "/getNFTMetadata"
	method := "GET"

	if *flagTokenType == "erc721" {
		for i := 0; i < *flagTokenCount; i += 1 {
			tokenId := strconv.Itoa(i)

			client := &http.Client{}
			req, err := http.NewRequest(
				method,
				url+"?contractAddress="+*flagContractAddress+"&tokenId="+tokenId+"&tokenType="+*flagTokenType,
				nil,
			)
			if err != nil {
				fmt.Println(err)
				return
			}

			res, err := client.Do(req)
			fmt.Print(res)
			if err != nil {
				fmt.Println(err)
				return
			}

			defer res.Body.Close()
			out, err := os.Create(*flagContractAddress + "-" + tokenId + ".json")
			if err != nil {
				fmt.Println(err)
				return
			}
			// body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println(string(body))
			io.Copy(out, res.Body)
		}

	} else {
	}
}
