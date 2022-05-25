package types

import (
	"encoding/json"
	"fmt"
	"log"
)

type Volume struct {
	Aliases          string            `json:"aliases,omitempty" yaml:"aliases,omitempty"`
	APIDetailURL     string            `json:"api_detail_url"`
	CharacterCredits Characters        `json:"character_credits,omitempty" yaml:"character_credits,omitempty"`
	ConceptCredits   Concepts          `json:"concept_credits,omitempty" yaml:"concept_credits,omitempty"`
	CountOfIssues    int               `json:"count_of_issues,omitempty" yaml:"count_of_issues,omitempty"`
	DateAdded        string            `json:"date_added,omitempty" yaml:"date_added,omitempty"`
	Deck             string            `json:"deck,omitempty" yaml:"deck,omitempty"`
	Description      string            `json:"description,omitempty" yaml:"description,omitempty"`
	FirstIssue       Issue             `json:"first_issue,omitempty" yaml:"first_issue,omitempty"`
	Id               int               `json:"id"`
	Image            map[string]string `json:"image,omitempty" yaml:"image,omitempty"`
	LastIssue        Issue             `json:"last_issue,omitempty" yaml:"last_issue,omitempty"`
	LocationCredits  Locations         `json:"location_credits,omitempty" yaml:"location_credits,omitempty"`
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`
	ObjectCredits    Objects           `json:"object_credits,omitempty" yaml:"object_credits,omitempty"`
	PersonCredits    People            `json:"person_credits,omitempty" yaml:"person_credits,omitempty"`
	Publisher        Publisher         `json:"publisher,omitempty" yaml:"publisher,omitempty"`
	SiteDetailUrl    string            `json:"site_detail_url,omitempty" yaml:"site_detail_url,omitempty"`
	StartYear        string            `json:"start_year,omitempty" yaml:"start_year,omitempty"`
	TeamCredits      Teams             `json:"team_credits,omitempty" yaml:"team_credits,omitempty"`
}
type Volumes []*Volume

func (a *ComicVineClient) GetVolumes(opts GetOptions) (Volumes, error) {
	var resp Volumes

	responseObject, err := a.genericGet(ResourceTypeVolume, opts)
	if err != nil {
		return resp, err
	}
	resp = make(Volumes, len(responseObject.Results))
	for i, data := range responseObject.Results {
		var volume Volume
		err = json.Unmarshal(data, &volume)
		if err != nil {
			log.Fatal(err)
		}
		resp[i] = &volume
	}
	return resp, nil
}

func (a *ComicVineClient) GetVolumeById(id int, opts GetOptions) (*Volume, error) {
	opts.Filters = append(opts.Filters, fmt.Sprintf("id:%d", id))
	volumes, err := a.GetVolumes(opts)
	if err != nil {
		return nil, err
	}
	if len(volumes) != 1 {
		return nil, fmt.Errorf("wrong number of volumes returned, expected 1, got %d", len(volumes))
	}
	return volumes[0], nil
}
