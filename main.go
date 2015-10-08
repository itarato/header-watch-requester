// @TODO
// - make it secure (maybe POST or JWT)

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var ()

// ServerHandler handles incoming requests.
type ServerHandler struct {
}

// CrawlingLocation keeps the location info in incoming JSON payload.
type CrawlingLocation struct {
	URL string `json:"url"`
}

// CrawlingInfoJSON is the format of the incoming JSON.
type CrawlingInfoJSON struct {
	Locations []CrawlingLocation `json:"locations"`
}

// CrawlingLocationResult keeps records of location crawl results.
type CrawlingLocationResult struct {
	URL        string              `json:"url"`
	StatusCode int                 `json:"status_code"`
	Protocol   string              `json:"protocol"`
	Headers    map[string][]string `json:"headers"`
}

// CrawlingResponseJSON is the returned object.
type CrawlingResponseJSON struct {
	Locations []CrawlingLocationResult `json:"locations"`
}

func (sh *ServerHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("incoming request")

	payload := &CrawlingInfoJSON{}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(400)
		io.WriteString(rw, "Unable to read payload.")
		return
	}

	if err = json.Unmarshal(b, payload); err != nil {
		rw.WriteHeader(400)
		io.WriteString(rw, "Unable to parse JSON. "+err.Error())
		return
	}

	responseJSON := &CrawlingResponseJSON{}
	for _, location := range payload.Locations {
		log.Println(location.URL)
		locationInfo := CrawlingLocationResult{URL: location.URL}

		resp, err := LoadURL(location.URL)
		if err != nil {
			log.Println("unable to fetch header for: " + location.URL)
		} else {
			locationInfo.Headers = resp.Header
			locationInfo.StatusCode = resp.StatusCode
			locationInfo.Protocol = resp.Proto
		}

		responseJSON.Locations = append(responseJSON.Locations, locationInfo)
	}

	jsonOut, err := json.Marshal(responseJSON)
	if err != nil {
		rw.WriteHeader(400)
		io.WriteString(rw, "Unable to generate JSON. "+err.Error())
		return
	}

	rw.WriteHeader(200)
	rw.Write(jsonOut)
}

// LoadURL loads the path and returns the header info.
func LoadURL(path string) (*http.Response, error) {
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func main() {
	log.Println("crawler has been started")

	server := &http.Server{
		// @TODO make it a configuration
		Addr:           "localhost:8080",
		Handler:        &ServerHandler{},
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 0,
	}

	log.Println("server has been initialized")

	err := server.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}

	log.Println("listen loop has been started")

	// @todo figure out if this is effective
	for {
	}
}
