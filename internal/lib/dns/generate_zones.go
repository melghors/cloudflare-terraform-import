package dns

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/emgag/cloudflare-terraform-import/internal/lib/util"
	"io"
	"text/template"
)

var zonesTemplate = `
resource "cloudflare_zone" "{{ .ResourceName }}" {
	zone = "{{ .ZoneName }}"
  plan = "free"
}
`

type Exporter_zone struct {
	API *cloudflare.API
}

func (e *Exporter_zone) DumpZones(tfWriter io.Writer, shWriter io.Writer) error {

	recs, err := e.API.ListZones()

	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(zonesTemplate)

	rrCounter := make(map[string]int)

	for _, r := range recs {
		fmt.Println(r.Name)
		t := fmt.Sprintf(
			"%s",
			r.Name,
		)

		rrCounter[t]++

		if rrCounter[t] > 1 {
			t = fmt.Sprintf("%d", rrCounter[t])
		}

		t = util.ToZoneName(t)

		err := tmpl.Execute(tfWriter, struct {
			ZoneName	string
			ResourceName	string
		}{r.Name, t})

		if err != nil {
			return err
		}

		fmt.Fprintf(shWriter, "terraform import cloudflare_zone.%s %s\n", t, r.ID)
	}

	return nil
}

func NewZonesExporter(api *cloudflare.API) *Exporter_zone {
	return &Exporter_zone{API: api}
}
