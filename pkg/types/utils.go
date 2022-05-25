package types

import (
	"encoding/json"
)

type ResourceType int

const (
	ResourceTypeCharacter ResourceType = iota
	ResourceTypeChat
	ResourceTypeConcept
	ResourceTypeEpisode
	ResourceTypeIssue
	ResourceTypeLocation
	ResourceTypeMovie
	ResourceTypeObject
	ResourceTypeOrigin
	ResourceTypePerson
	ResourceTypePower
	ResourceTypePromo
	ResourceTypePublisher
	ResourceTypeSeries
	ResourceTypeStoryArc
	ResourceTypeTeam
	ResourceTypeTypes
	ResourceTypeVideo
	ResourceTypeVolume
)

type Response struct {
	StatusCode   int               `json:"status_code"`
	Error        string            `json:"error"`
	TotalResults int               `json:"number_of_total_results,omitempty"`
	PageResults  int               `json:"number_of_page_results,omitempty"`
	Limit        int               `json:"limit,omitempty"`
	Offset       int               `json:"offset,omitempty"`
	Results      []json.RawMessage `json:"results,omitempty"` // Collect the results as just bytes that'll be unmarheled later
}

type GetOptions struct {
	Fields  []string
	Limit   int
	Offset  int
	Sort    string
	Filters []string
}
