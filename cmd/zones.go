package cmd

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/emgag/cloudflare-terraform-import/internal/lib/dns"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(importZonesCmd)
}

var importZonesCmd = &cobra.Command{
	Use:   "zones",
	Short: "Generate import zones files for zone",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))

		if err != nil {
			log.Fatal(err)
		}

		e := dns.NewZonesExporter(api)

		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			os.Mkdir("zones", 0755)
		}

		if err != nil {
			log.Fatal(err)
		}

		tf, err := os.Create(fmt.Sprintf("zones/zones.tf"))

		if err != nil {
			log.Fatal(err)
		}

		sh, err := os.Create(fmt.Sprintf("zones/zones.sh"))

		if err != nil {
			log.Fatal(err)
		}

		e.DumpZones(tf, sh)
	},
}
