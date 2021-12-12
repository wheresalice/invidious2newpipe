package lib

import (
	"encoding/xml"
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

func XmlUrlToChanelUrl(xmlUrl string) string {
	var re = regexp.MustCompile("https://www.youtube.com/feeds/videos.xml\\?channel_id=(.+)")
	s := re.ReplaceAllString(xmlUrl, "https://www.youtube.com/channel/$1")
	return s
}
