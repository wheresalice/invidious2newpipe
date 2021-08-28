package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type Opml struct {
	XMLName xml.Name `xml:"opml"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Body    struct {
		Text    string `xml:",chardata"`
		Outline struct {
			Text     string `xml:",chardata"`
			AttrText string `xml:"text,attr"`
			Title    string `xml:"title,attr"`
			Outline  []struct {
				Text     string `xml:",chardata"`
				AttrText string `xml:"text,attr"`
				Title    string `xml:"title,attr"`
				Type     string `xml:"type,attr"`
				XmlUrl   string `xml:"xmlUrl,attr"`
			} `xml:"outline"`
		} `xml:"outline"`
	} `xml:"body"`
}

type NewPipe struct {
	Subscriptions []Subscriptions `json:"subscriptions"`
}
type Subscriptions struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	ServiceID int    `json:"service_id,omitempty"`
}

func xmlUrlToChanelUrl(xmlUrl string) string {
	var re = regexp.MustCompile("https://www.youtube.com/feeds/videos.xml\\?channel_id=(.+)")
	s := re.ReplaceAllString(xmlUrl, "https://www.youtube.com/channel/$1")
	return s
}

func main() {
	opmlPath := ""

	if len(os.Args) > 1 {
		opmlPath = os.Args[1]
	} else {
		opmlPath = "subscription_manager"
	}

	xmlFile, err := os.Open(opmlPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var opml Opml
	err = xml.Unmarshal(byteValue, &opml)
	if err != nil {
		log.Fatalln(err)
	}

	var newpipe NewPipe

	for _, s := range opml.Body.Outline.Outline {
		newpipe.Subscriptions = append(newpipe.Subscriptions, Subscriptions{
			Name:      s.Title,
			URL:       xmlUrlToChanelUrl(s.XmlUrl),
			ServiceID: 0,
		})
	}

	output, err := json.Marshal(newpipe)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", output)
}
