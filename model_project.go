package yrac

// Represents a YouTrack project.
type Project struct {
	Name           string      `json:"name,omitempty"`
	Id             string      `json:"id,omitempty"`
	Type_          string      `json:"$type,omitempty"`
	StartingNumber int64       `json:"startingNumber,omitempty"`
	ShortName      string      `json:"shortName,omitempty"`
	Description    string      `json:"description,omitempty"`
	Leader         *User       `json:"leader,omitempty"`
	CreatedBy      *User       `json:"createdBy,omitempty"`
	Issues         []Issue     `json:"issues,omitempty"`
	CustomFields   interface{} `json:"customFields,omitempty"`
	Archived       bool        `json:"archived,omitempty"`
	FromEmail      string      `json:"fromEmail,omitempty"`
	ReplyToEmail   string      `json:"replyToEmail,omitempty"`
	Template       bool        `json:"template,omitempty"`
	IconUrl        string      `json:"iconUrl,omitempty"`
	Team           *UserGroup  `json:"team,omitempty"`
}

// Represents the change in the project attribute on an Issue.
type ProjectActivityItem struct {
	Target  *Issue      `json:"target,omitempty"`
	Removed interface{} `json:"removed,omitempty"`
	Added   interface{} `json:"added,omitempty"`
}

// Lets you set a color for a card based on its project.
type ProjectBasedColorCoding struct {
	Id            string         `json:"id,omitempty"`
	Type_         string         `json:"$type,omitempty"`
	ProjectColors []ProjectColor `json:"projectColors,omitempty"`
}

// Represents color setting for one project on the board.
type ProjectColor struct {
	Project *Project    `json:"project,omitempty"`
	Color   *FieldStyle `json:"color,omitempty"`
	Id      string      `json:"id,omitempty"`
	Type_   string      `json:"$type,omitempty"`
}

// Represents custom field settings for the particular project.
type ProjectCustomField struct {
	Field          *CustomField `json:"field,omitempty"`
	Project        *Project     `json:"project,omitempty"`
	CanBeEmpty     bool         `json:"canBeEmpty,omitempty"`
	EmptyFieldText string       `json:"emptyFieldText,omitempty"`
	Ordinal        int32        `json:"ordinal,omitempty"`
	IsPublic       bool         `json:"isPublic,omitempty"`
	HasRunningJob  bool         `json:"hasRunningJob,omitempty"`
	Id             string       `json:"id,omitempty"`
	Type_          string       `json:"$type,omitempty"`
}

// Represents time tracking settings of the project.
type ProjectTimeTrackingSettings struct {
	Enabled       bool                `json:"enabled,omitempty"`
	Estimate      *ProjectCustomField `json:"estimate,omitempty"`
	TimeSpent     *ProjectCustomField `json:"timeSpent,omitempty"`
	WorkItemTypes []WorkItemType      `json:"workItemTypes,omitempty"`
	Project       *Project            `json:"project,omitempty"`
	Id            string              `json:"id,omitempty"`
	Type_         string              `json:"$type,omitempty"`
}

// Represents a work type that can be assigned to a work item.
type WorkItemType struct {
	Name         string `json:"name,omitempty"`
	AutoAttached bool   `json:"autoAttached,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}
