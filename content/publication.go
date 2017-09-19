package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Publication ->
type Publication struct {
	item.Item

	Title          string   `json:"title"`
	Format         string   `json:"format"`
	ResearchTopics []string `json:"research-topics"`
	Authors        []string `json:"authors"`
	Labs           []string `json:"labs"`
	Year           string   `json:"year"`
	Citation       string   `json:"citation"`
	Abstract       string   `json:"abstract"`
	Pdf            string   `json:"pdf"`
	Url            string   `json:"url"`
}

// MarshalEditor writes a buffer of html to edit a Publication within the CMS
// and implements editor.Editable
func (p *Publication) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Publication field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Select("Format", p, map[string]string{
				"label": "Format",
			}, map[string]string{
				// "value": "Display Name",
				"journal-article":  "journal article",
				"conference-paper": "conference paper",
				"poster":           "poster",
				"technical-report": "technical report",
				"book":             "book",
				"thesis":           "thesis",
			}),
		},
		editor.Field{
			View: editor.Tags("ResearchTopics", p, map[string]string{
				"label":       "ResearchTopics",
				"placeholder": "+ResearchTopics",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Authors", p, map[string]string{
				"label": "Authors",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Labs", p, map[string]string{
				"label": "Labs",
			},
				"Lab",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Input("Year", p, map[string]string{
				"label":       "Year",
				"type":        "text",
				"placeholder": "Enter the Year here",
			}),
		},
		editor.Field{
			View: editor.Input("Citation", p, map[string]string{
				"label":       "Citation",
				"type":        "text",
				"placeholder": "Enter the Citation here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Abstract", p, map[string]string{
				"label":       "Abstract",
				"placeholder": "Enter the Abstract here",
			}),
		},
		editor.Field{
			View: editor.File("Pdf", p, map[string]string{
				"label":       "Pdf",
				"placeholder": "Upload the Pdf here",
			}),
		},
		editor.Field{
			View: editor.Input("Url", p, map[string]string{
				"label":       "Url",
				"type":        "text",
				"placeholder": "Enter the Url here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Publication editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Publication"] = func() interface{} { return new(Publication) }
}

func (p *Publication) String() string {
	return p.Title
}
