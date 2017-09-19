package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Event ->
type Event struct {
	item.Item

	Title          string   `json:"title"`
	Start          string   `json:"start"`
	Finish         string   `json:"finish"`
	Location       string   `json:"location"`
	Labs           []string `json:"labs"`
	Hosts          []string `json:"hosts"`
	ResearchTopics []string `json:"research-topics"`
	SplashImage    string   `json:"splash-image"`
	Summary        string   `json:"summary"`
	Description    string   `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Event within the CMS
// and implements editor.Editable
func (e *Event) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(e,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Event field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", e, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Start", e, map[string]string{
				"label":       "Start",
				"type":        "text",
				"placeholder": "Enter the Start here",
			}),
		},
		editor.Field{
			View: editor.Input("Finish", e, map[string]string{
				"label":       "Finish",
				"type":        "text",
				"placeholder": "Enter the Finish here",
			}),
		},
		editor.Field{
			View: editor.Input("Location", e, map[string]string{
				"label":       "Location",
				"type":        "text",
				"placeholder": "Enter the Location here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Labs", e, map[string]string{
				"label": "Labs",
			},
				"Lab",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Hosts", e, map[string]string{
				"label": "Hosts",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Tags("ResearchTopics", e, map[string]string{
				"label":       "ResearchTopics",
				"placeholder": "+ResearchTopics",
			}),
		},
		editor.Field{
			View: editor.File("SplashImage", e, map[string]string{
				"label":       "SplashImage",
				"placeholder": "Upload the SplashImage here",
			}),
		},
		editor.Field{
			View: editor.Input("Summary", e, map[string]string{
				"label":       "Summary",
				"type":        "text",
				"placeholder": "Enter the Summary here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", e, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Event editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Event"] = func() interface{} { return new(Event) }
}

func (e *Event) String() string {
	return e.Title + " @ " + e.Start
}
