package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Person -> describes atlas member or peer
type Person struct {
	item.Item

	Name              string   `json:"name"`
	Role              string   `json:"role"`
	ResearchInterests []string `json:"research-interests"`
	Portrait          string   `json:"portrait"`
	Quip              string   `json:"quip"`
	Bio               string   `json:"bio"`
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
			View: editor.Select("Role", p, map[string]string{
				"label": "Role",
			}, map[string]string{
				// "value": "Display Name",
				"peer":       "peer (not atlas member)",
				"faculty":    "faculty",
				"researcher": "researcher",
				"student":    "student",
				"alumnus":    "alumnus",
			}),
		},
		editor.Field{
			View: editor.Tags("ResearchInterests", p, map[string]string{
				"label":       "ResearchInterests",
				"placeholder": "+ResearchInterests",
			}),
		},
		editor.Field{
			View: editor.File("Portrait", p, map[string]string{
				"label":       "Portrait",
				"placeholder": "Upload the Portrait here",
			}),
		},
		editor.Field{
			View: editor.Input("Quip", p, map[string]string{
				"label":       "Quip",
				"type":        "text",
				"placeholder": "Enter the Quip here",
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
