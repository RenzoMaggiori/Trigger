package reaction

import (
	"trigger.com/trigger/pkg/action"
)

// MergeRequest represents the structure for the given JSON data
type PullRequest struct {
	Title       string     `json:"title"`
	Source      BranchInfo `json:"source"`
	Destination BranchInfo `json:"destination"`
}

// BranchInfo represents the source and destination fields
type BranchInfo struct {
	Branch Branch `json:"branch"`
}

// Branch represents the branch field in the source and destination
type Branch struct {
	Name string `json:"name"`
}

type Service interface {
	action.Reaction
}

type Handler struct {
	Service
}

type Model struct {
}
