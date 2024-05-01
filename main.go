package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

func main() {
	BASE_PATH := "https://data.bayerncloud.digital/api/v4/endpoints"
	LOGIN := "cloudservices@banbutsu.com"
	PASSWORD := "82Rfppe3xZ3Qz3cT79vi2p0qwoXS6Gtk"
	encodedString := base64.StdEncoding.EncodeToString([]byte(LOGIN + ":" + PASSWORD))

	req, _ := http.NewRequest("GET", BASE_PATH+"/list_food", nil)
	req.Header.Add("Authorization", "Basic "+encodedString)

	client := http.Client{}
	resp, _ := client.Do(req)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
