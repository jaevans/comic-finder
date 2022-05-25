package types

import (
	"encoding/json"
	"fmt"
	"log"
)

type Issue struct {
	Aliases                   string            `json:"aliases,omitempty" yaml:"aliases,omitempty"`
	APIDetailURL              string            `json:"api_detail_url"`
	CharacterCredits          Characters        `json:"character_credits,omitempty" yaml:"character_credits,omitempty"`
	CharactersDiedIn          Characters        `json:"characters_died_in,omitempty" yaml:"characters_died_in,omitempty"`
	ConceptCredits            Concepts          `json:"concept_credits,omitempty" yaml:"concept_credits,omitempty"`
	CoverDate                 string            `json:"cover_date,omitempty" yaml:"cover_date,omitempty"`
	DateAdded                 string            `json:"date_added,omitempty" yaml:"date_added,omitempty"`
	DateLastUpdated           string            `json:"date_last_updated,omitempty" yaml:"date_last_updated,omitempty"`
	Deck                      string            `json:"deck,omitempty" yaml:"deck,omitempty"`
	Description               string            `json:"description,omitempty" yaml:"description,omitempty"`
	DisbandedTeams            Teams             `json:"disbanded_teams,omitempty" yaml:"disbanded_teams,omitempty"`
	FirstAppearanceCharacters Characters        `json:"first_appearance_characters,omitempty" yaml:"first_appearance_characters,omitempty"`
	FirstAppearanceConcepts   Concepts          `json:"first_appearance_concepts,omitempty" yaml:"first_appearance_concepts,omitempty"`
	FirstAppearanceLocations  Locations         `json:"first_appearance_locations,omitempty" yaml:"first_appearance_locations,omitempty"`
	FirstAppearanceObjects    Objects           `json:"first_appearance_objects,omitempty" yaml:"first_appearance_objects,omitempty"`
	FirstAppearanceStoryArcs  StoryArcs         `json:"first_appearance_story_arcs,omitempty" yaml:"first_appearance_story_arcs,omitempty"`
	FirstAppearanceTeams      Teams             `json:"first_appearance_teams,omitempty" yaml:"first_appearance_teams,omitempty"`
	HasStaffReview            bool              `json:"has_staff_review,omitempty" yaml:"has_staff_review,omitempty"`
	Id                        int               `json:"id"`
	Image                     map[string]string `json:"image,omitempty" yaml:"image,omitempty"`
	IssueNumber               string            `json:"issue_number,omitempty" yaml:"issue_number,omitempty"`
	LocationCredits           Locations         `json:"location_credits,omitempty" yaml:"location_credits,omitempty"`
	PersonCredits             People            `json:"person_credits,omitempty" yaml:"person_credits,omitempty"`
	SiteDetailUrl             string            `json:"site_detail_url,omitempty" yaml:"site_detail_url,omitempty"`
	StoreDate                 string            `json:"store_date,omitempty" yaml:"store_date,omitempty"`
	StoryArcCredits           StoryArcs         `json:"story_arc_credits,omitempty" yaml:"story_arc_credits,omitempty"`
	TeamCredits               Teams             `json:"team_credits,omitempty" yaml:"team_credits,omitempty"`
	TeamsDisabndedIn          Teams             `json:"teams_disbanded_in,omitempty" yaml:"teams_disbanded_in,omitempty"`
	Volume                    *Volume           `json:"volume,omitempty" yaml:"volume,omitempty"`
}
type Issues []*Issue

func (a *ComicVineClient) GetIssues(opts GetOptions) (Issues, error) {
	var resp Issues

	responseObject, err := a.genericGet(ResourceTypeIssue, opts)
	if err != nil {
		return resp, err
	}
	resp = make(Issues, len(responseObject.Results))
	for i, data := range responseObject.Results {
		var issue Issue
		err = json.Unmarshal(data, &issue)
		if err != nil {
			log.Fatal(err)
		}
		resp[i] = &issue
	}
	return resp, nil
}

func (a *ComicVineClient) GetIssueById(id int, opts GetOptions) (*Issue, error) {
	opts.Filters = append(opts.Filters, fmt.Sprintf("id:%d", id))
	issues, err := a.GetIssues(opts)
	if err != nil {
		return nil, err
	}
	if len(issues) != 1 {
		return nil, fmt.Errorf("wrong number of issues returned, expected 1, got %d", len(issues))
	}
	return issues[0], nil
}
