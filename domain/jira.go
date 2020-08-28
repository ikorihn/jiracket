package domain

type Issue struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields Field  `json:"fields"`
}

type Field struct {
	Assignee    Assignee `json:"assignee"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
}

type Assignee struct {
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
}
