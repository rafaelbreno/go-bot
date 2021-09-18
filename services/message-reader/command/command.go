package commands

// Command stores all data related
// to one command
type Command struct {
	Trigger string  `json:"trigger"`
	Answer  string  `json:"answer"`
	Fields  []Field `json:"fields"`
	SentBy  string  `json:"sent_by"`
}

type Field struct {
	Key       string   `json:"key"`
	Values    []string `json:"values"`
	Blacklist []string `json:"blacklist"`
}
