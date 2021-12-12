
package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/wheresalice/invidious2newpipe/lib"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Convert an invidious OPML export to a Newpipe subscription json file locally",
	Run: func(cmd *cobra.Command, args []string) {

		opmlPath := ""

		if len(os.Args) > 2 {
			opmlPath = os.Args[2]
		} else {
			opmlPath = "subscription_manager"
		}

		xmlFile, err := os.Open(opmlPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer xmlFile.Close()

		byteValue, _ := ioutil.ReadAll(xmlFile)
		var opml lib.Opml
		err = xml.Unmarshal(byteValue, &opml)
		if err != nil {
			log.Fatalln(err)
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
			log.Fatalln(err)
		}
		fmt.Printf("%s\n", output)	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
