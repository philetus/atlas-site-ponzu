package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Lab ->
type Lab struct {
	item.Item

	Name           string   `json:"name"`
	ResearchTopics []string `json:"research-topics"`
	Directors      []string `json:"directors"`
	Members        []string `json:"members"`
	Logo           string   `json:"logo"`
	SplashImage    string   `json:"splash-image"`
	Mission        string   `json:"mission"`
	Description    string   `json:"description"`
}

// MarshalEditor writes a buffer of html to edit a Lab within the CMS
// and implements editor.Editable
func (l *Lab) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(l,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Lab field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", l, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Tags("ResearchTopics", l, map[string]string{
				"label":       "ResearchTopics",
				"placeholder": "+ResearchTopics",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Directors", l, map[string]string{
				"label": "Directors",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Members", l, map[string]string{
				"label": "Members",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.File("Logo", l, map[string]string{
				"label":       "Logo",
				"placeholder": "Upload the Logo here",
			}),
		},
		editor.Field{
			View: editor.File("SplashImage", l, map[string]string{
				"label":       "SplashImage",
				"placeholder": "Upload the SplashImage here",
			}),
		},
		editor.Field{
			View: editor.Input("Mission", l, map[string]string{
				"label":       "Mission",
				"type":        "text",
				"placeholder": "Enter the Mission here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", l, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Lab editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Lab"] = func() interface{} { return new(Lab) }
}

func (l *Lab) String() string {
	return l.Name
}
