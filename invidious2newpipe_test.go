package main

import "testing"

func TestChannelUrl(t *testing.T) {
	channelUrl := xmlUrlToChanelUrl(`https://www.youtube.com/feeds/videos.xml?channel_id=12345`)
	expectedUrl := `https://www.youtube.com/channel/12345`
	if channelUrl != expectedUrl {
		t.Errorf("returned incorrect channel url: got %v want %v", channelUrl, expectedUrl)
	}
}
