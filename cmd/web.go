package cmd

import (
	"crypto/sha256"
	_ "embed"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/wheresalice/invidious2newpipe/lib"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run invidious2newpipe as a webserver",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Println("invidious2newpipe -- Starting ")
		log.Printf("http://localhost:%v\n", port)

		os.Mkdir("subs", 0700)

		if value, ok := os.LookupEnv("PORT"); ok {
			port = value
		}

		http.HandleFunc("/", reqHandler)
		log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
	},
}

var (
	port = "5000"
)

//go:embed index.html
var index string

func handleGetSubs(w http.ResponseWriter, r *http.Request) {
	path := filepath.Clean(r.URL.Path)
	log.Printf("GET %v", path)

	if (path == "/") || (path == "/index.html") {
		fmt.Fprintf(w, index)
	} else {
		// otherwise, if the requested paste exists, we serve it...

		subs, err := os.ReadFile("subs/" + path)
		if err != nil {
			http.NotFound(w, r)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(subs))
	}

}

func getHash(content string) string {
	h := sha256.New()
	h.Write([]byte(content))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func handlePutSubs(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseForm()
	err2 := r.ParseMultipartForm(int64(2 * 4096))

	if err1 != nil && err2 != nil {
		// Invalid POST -- let's serve the default file
		fmt.Fprintf(w, index)
	} else {
		file, handler, err := r.FormFile("subs")
		if err != nil {
			fmt.Fprintf(w, "error retrieving file: %v", err)
			return
		}
		defer file.Close()
		log.Printf("Uploaded File: %+v\n", handler.Filename)
		log.Printf("File Size: %+v\n", handler.Size)
		log.Printf("MIME Header: %+v\n", handler.Header)
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Fprintf(w, "failed to read file: %v", err)
		}

		hash := getHash(string(fileBytes))

		var opml lib.Opml
		err = xml.Unmarshal(fileBytes, &opml)
		if err != nil {
			fmt.Fprintf(w, "failed to parse: %v", err)
		}

		var newpipe lib.NewPipe
		for _, s := range opml.Body.Outline.Outline {
			newpipe.Subscriptions = append(newpipe.Subscriptions, lib.Subscriptions{
				Name:      s.Title,
				URL:       lib.XmlUrlToChanelUrl(s.XmlUrl),
				ServiceID: 0,
			})
		}

		output, err := json.Marshal(newpipe)
		if err != nil {
			fmt.Fprintf(w, "failed to marshal content: %v", err)
		}
		os.WriteFile("subs/"+hash, output, 0700)
		http.Redirect(w, r, hash, 301)
	}

}

func reqHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		handleGetSubs(w, r)
	case "POST":
		handlePutSubs(w, r)
	default:
		http.NotFound(w, r)
	}
}


func init() {
	rootCmd.AddCommand(webCmd)
}
