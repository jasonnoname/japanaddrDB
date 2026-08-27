package model

// ValidationStatus describes how completely an input address was found in ABR.
type ValidationStatus string

const (
	ValidationStatusExact      ValidationStatus = "exact"
	ValidationStatusBaseOnly   ValidationStatus = "base_address_only"
	ValidationStatusIncomplete ValidationStatus = "insufficient_detail"
	ValidationStatusNotFound   ValidationStatus = "not_found"
)

// ValidateResponse is a stable, application-oriented view of the best ABR
// match. Exists means that an official residence number or parcel was found.
// Exact additionally means that the matcher consumed the complete input.
type ValidateResponse struct {
	Exists            bool              `json:"exists"`
	Exact             bool              `json:"exact"`
	Status            ValidationStatus  `json:"status"`
	InputAddress      string            `json:"input_address"`
	MatchedAddress    string            `json:"matched_address"`
	UnmatchedAddress  []string          `json:"unmatched_address"`
	MatchLevel        MatchLevel        `json:"match_level"`
	IDs               IDs               `json:"ids"`
	StructuredAddress StructuredAddress `json:"structured_address"`
	ResultInfo        ResultInfo        `json:"result_info"`
}
