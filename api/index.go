package handler

import (
	"io/ioutil"
	"net/http"
	"os"
)

// Handler : main
func Handler(w http.ResponseWriter, r *http.Request) {
	apiFlashEndpoint := "https://api.apiflash.com/v1/urltoimage"
	request, _ := http.NewRequest("GET", apiFlashEndpoint, nil)
	query := request.URL.Query()
	query.Add("access_key", os.Getenv("APIFLASH_ACCESS_KEY"))
	query.Add("url", "https://sugarshin.net")
	query.Add("width", "854")
	query.Add("height", "505")
	query.Add("format", "png")
	request.URL.RawQuery = query.Encode()
	client := &http.Client{}
	response, _ := client.Do(request)
	image, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}
