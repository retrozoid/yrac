package yrac

type CommandVisibility struct {
	Type string `json:"$type"`
}

type CommandLimitedVisibility struct {
	CommandVisibility
	PermittedGroups []UserGroup `json:"permittedGroups,omitempty"`
	PermittedUsers  []User      `json:"permittedUsers,omitempty"`
}

type ParsedCommand struct {
	Description string `json:"description,omitempty"`
	Error       bool   `json:"error,omitempty"`
	Delete      bool   `json:"delete,omitempty"`
}

type CommandIssue struct {
	Id         string `json:"id,omitempty"`
	IdReadable string `json:"idReadable,omitempty"`
}

type Command struct {
	Query        string                    `json:"query"`
	Issues       []CommandIssue            `json:"issues"`
	Comment      string                    `json:"comment,omitempty"`
	Visibility   *CommandLimitedVisibility `json:"visibility,omitempty"`
	Caret        int64                     `json:"caret,omitempty"`
	Silent       bool                      `json:"silent,omitempty"`
	UsesMarkdown bool                      `json:"usesMarkdown,omitempty"`
	RunAs        string                    `json:"runAs,omitempty"`
	Commands     []ParsedCommand           `json:"commands,omitempty"`
}
