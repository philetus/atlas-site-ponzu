package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Person -> stores info about a person at atlas!
type Person struct {
	item.Item

	Name  string `json:"name"`
	Photo string `json:"photo"`
	Bio   string `json:"bio"`
}

// MarshalEditor writes a buffer of html to edit a Person within the CMS
// and implements editor.Editable
func (p *Person) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Person field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", p, map[string]string{
				"label":       "Photo",
				"placeholder": "Upload the Photo here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Bio", p, map[string]string{
				"label":       "Bio",
				"placeholder": "Enter the Bio here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Person editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Person"] = func() interface{} { return new(Person) }
}

func (p *Person) String() string {
	return p.Name
}
