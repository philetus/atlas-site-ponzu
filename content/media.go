package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Media ->
type Media struct {
	item.Item

	Title          string   `json:"title"`
	Url            string   `json:"url"`
	Flavor         string   `json:"flavor"`
	ResearchTopics []string `json:"research-topics"`
	SplashImage    string   `json:"splash-image"`
	Summary        string   `json:"summary"`
	Owner          string   `json:"owner"`
	Authors        []string `json:"authors"`
	Featured       []string `json:"featured"`
	Publications   []string `json:"publications"`
	Projects       []string `json:"projects"`
	Labs           []string `json:"labs"`
	Events         []string `json:"events"`
}

// MarshalEditor writes a buffer of html to edit a Media within the CMS
// and implements editor.Editable
func (m *Media) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(m,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Media field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", m, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Input("Url", m, map[string]string{
				"label":       "Url",
				"type":        "text",
				"placeholder": "Enter the Url here",
			}),
		},
		editor.Field{
			View: editor.Select("Flavor", m, map[string]string{
				"label": "Flavor",
			}, map[string]string{
				// "value": "Display Name",
				"website":       "website",
				"youtube-video": "youtube video",
				"github":        "github",
				"news-story":    "news story",
				"instagram":     "instagram (stream)",
				"flickr":        "flickr (stream)",
			}),
		},
		editor.Field{
			View: editor.Tags("ResearchTopics", m, map[string]string{
				"label":       "ResearchTopics",
				"placeholder": "+ResearchTopics",
			}),
		},
		editor.Field{
			View: editor.File("SplashImage", m, map[string]string{
				"label":       "SplashImage",
				"placeholder": "Upload the SplashImage here",
			}),
		},
		editor.Field{
			View: editor.Input("Summary", m, map[string]string{
				"label":       "Summary",
				"type":        "text",
				"placeholder": "Enter the Summary here",
			}),
		},
		editor.Field{
			View: reference.Select("Owner", m, map[string]string{
				"label": "Owner",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Authors", m, map[string]string{
				"label": "Authors",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Featured", m, map[string]string{
				"label": "Featured",
			},
				"Person",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Publications", m, map[string]string{
				"label": "Publications",
			},
				"Publication",
				`{{ .title }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Projects", m, map[string]string{
				"label": "Projects",
			},
				"Project",
				`{{ .title }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Labs", m, map[string]string{
				"label": "Labs",
			},
				"Lab",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Events", m, map[string]string{
				"label": "Events",
			},
				"Event",
				`{{ .title }} {{ .start }} `,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Media editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Media"] = func() interface{} { return new(Media) }
}

func (m *Media) String() string {
	return m.Title + " @ " + m.Url
}
