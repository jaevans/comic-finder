package types

type Publisher struct {
	Aliases         string            `json:"aliases,omitempty" yaml:"aliases,omitempty"`
	APIDetailURL    string            `json:"api_detail_url"`
	Characters      Characters        `json:"characters,omitempty" yaml:"characters,omitempty"`
	DateAdded       string            `json:"date_added,omitempty" yaml:"date_added,omitempty"`
	DateLastUpdated string            `json:"date_last_updated,omitempty" yaml:"date_last_updated,omitempty"`
	Deck            string            `json:"deck,omitempty" yaml:"deck,omitempty"`
	Description     string            `json:"description,omitempty" yaml:"description,omitempty"`
	Id              int               `json:"id"`
	Image           map[string]string `json:"image,omitempty" yaml:"image,omitempty"`
	LocationAddress string            `json:"location_address,omitempty" yaml:"location_address,omitempty"`
	LocationCity    string            `json:"location_city,omitempty" yaml:"location_city,omitempty"`
	LocationState   string            `json:"location_state,omitempty" yaml:"location_state,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	SiteDetailUrl   string            `json:"site_detail_url,omitempty" yaml:"site_detail_url,omitempty"`
	StoryArcs       StoryArcs         `json:"story_arcs,omitempty" yaml:"story_arcs,omitempty"`
	TeamCredits     Teams             `json:"teams,omitempty" yaml:"teams,omitempty"`
	Volumes         Volumes           `json:"volumes,omitempty" yaml:"volumes,omitempty"`
}
type Publishers []*Publisher
