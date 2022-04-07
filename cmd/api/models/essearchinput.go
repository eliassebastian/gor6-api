package model

type SearchMatch struct {
	AliasesName string `json:"aliases.name"`
}

type SearchQuery struct {
	Match SearchMatch `json:"match"`
}

type SearchInput struct {
	Query  SearchQuery `json:"query"`
	Fields []string    `json:"fields"`
	Source bool        `json:"_source"`
	Size   int         `json:"size"`
}
