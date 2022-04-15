package main

//todo
//1.  create config file to store api key to be hidden from .git
//2.  add if statement to check tokenType and for erc1155 don't loop, check just tokenID

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

var (
	contractAddress = "0xf97DF1e168c27e22Eedd34C05AE0615605C5DcbF"
	tokenCount      = 5555
	tokenType       = "erc721" //erc721 or erc1511

)

func main() {

	url := "https://eth-mainnet.alchemyapi.io/v2/9KfaK3_skwhS81PyS1wEN4GPfuFVyPWX/getNFTMetadata"
	method := "GET"

	if tokenType == "erc721" {
		for i := 0; i < tokenCount; i += 1 {
			tokenId := strconv.Itoa(i)

			//Build url

			// url := "https://eth-mainnet.alchemyapi.io/v2/demo/getNFTMetadata?contractAddress=0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d&tokenId=2&tokenType=erc721"
			// payload := strings.NewReader("contractAddress=" + contractAddress + "&tokenId=" + tokenId + "&tokenType=" + tokenType)
			// fmt.Print(payload)
			client := &http.Client{}
			req, err := http.NewRequest(
				method,
				url+"?contractAddress="+contractAddress+"&tokenId="+tokenId+"&tokenType="+tokenType,
				nil,
			)
			// fmt.Print(req)
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
			out, err := os.Create(contractAddress + "-" + tokenId + ".json")
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
		tokenId := "63763212585141476199723449601564140078511354041472126739157516503384422613942"
		client := &http.Client{}
		req, err := http.NewRequest(
			method,
			url+"?contractAddress="+contractAddress+"&tokenId="+tokenId+"&tokenType="+tokenType,
			nil,
		)
		fmt.Print(req)
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
		out, err := os.Create(contractAddress + "-" + tokenId + ".json")
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
}
