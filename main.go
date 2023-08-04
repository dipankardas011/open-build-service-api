package main

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	credential string
)

func GenerateCreds() {

	username := os.Getenv("OSC_USERNAME")
	password := os.Getenv("OSC_PASSWORD")

	credential = b64.StdEncoding.EncodeToString([]byte(username + ":" + password))
}

func ApiCall(method, endpoint string) ([]byte, error) {
	req, err := http.NewRequest(method, "https://api.opensuse.org"+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/xml; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", credential))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func getGPGKeyURL() string {
	return "/source/home%3Adipankardas/_pubkey"
}

func getKsctlLogsFedora38() string {
	return "/build/home%3Adipankardas/Fedora_38/x86_64/ksctl/_log"

}

func main() {
	GenerateCreds()
	url := getKsctlLogsFedora38()
	meth := "GET"
	resp, err := ApiCall(meth, url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

	url = getGPGKeyURL()
	meth = "GET"
	resp, err = ApiCall(meth, url)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

}
