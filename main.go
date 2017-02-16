package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/op/go-logging"
)

var (
	log = logging.MustGetLogger("giphy-k8s-demo")
)

// Pod stores the metadata.
type Pod struct {
	Annotations map[string]string
	Name        string
	GiphyID     string
}

type GiphyResp struct {
	Image GiphyImage `json:"data"`
}

type GiphyImage struct {
	ID string `json:"id"`
}

func main() {
	log.Info("Starting server on port 8080.")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Inspect the running container and render the page.
func handler(w http.ResponseWriter, r *http.Request) {
	pod, err := getPodMetadata()
	if err != nil {
		log.Errorf("Failed to get metadata from Downward API - %v", err)
	}

	gif, err := getRandomGif(pod.Annotations["giphy.com/search"])
	if err != nil {
		log.Errorf("Failed to get random gif - %v", err)
	}

	pod.GiphyID = gif.ID

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Errorf("Error loading template - %v", err)
	}

	t.Execute(w, pod)
}

// Gets pod metadata from the Downward API volume.
func getPodMetadata() (pod Pod, err error) {
	annotations, err := getAnnotations()
	if err != nil {
		log.Errorf("Failed to get annotations - %v", err)
		return
	}

	pod = Pod{
		Annotations: annotations,
		Name:        os.Getenv("POD_NAME"),
	}

	return
}

func getAnnotations() (map[string]string, error) {
	annotations := make(map[string]string)

	input, err := os.Open(os.Getenv("ANNOTATIONS_PATH"))
	if err != nil {
		log.Errorf("Failed to open annotations file - %v", err)
		return annotations, err
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "=")
		k := a[0]
		v := a[1]

		annotations[k] = strings.Replace(v, "\"", "", -1)
	}

	return annotations, err
}

func getRandomGif(search string) (image GiphyImage, err error) {
	apiURL := "http://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag="
	apiURL += strings.Replace(search, " ", "+", -1)

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Errorf("Failed to build API GET request - %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("GET request failed %s %d - %s", apiURL, resp.StatusCode, resp.Status)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Failed to read response to %s", apiURL)
		return
	}

	var data GiphyResp

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Errorf("Failed to unmarshal response to %s", apiURL)
		return
	}

	return data.Image, err
}
