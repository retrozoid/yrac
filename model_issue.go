package yrac

type Issue struct {
	IdReadable          string             `json:"idReadable,omitempty"`
	Created             int64              `json:"created,omitempty"`
	Updated             int64              `json:"updated,omitempty"`
	Resolved            int64              `json:"resolved,omitempty"`
	NumberInProject     int64              `json:"numberInProject,omitempty"`
	Project             *Project           `json:"project,omitempty"`
	Summary             string             `json:"summary,omitempty"`
	Description         string             `json:"description,omitempty"`
	UsesMarkdown        bool               `json:"usesMarkdown,omitempty"`
	WikifiedDescription string             `json:"wikifiedDescription,omitempty"`
	Reporter            *User              `json:"reporter,omitempty"`
	Updater             *User              `json:"updater,omitempty"`
	DraftOwner          *User              `json:"draftOwner,omitempty"`
	IsDraft             bool               `json:"isDraft,omitempty"`
	Visibility          *Visibility        `json:"visibility,omitempty"`
	Votes               int32              `json:"votes,omitempty"`
	Comments            []IssueComment     `json:"comments,omitempty"`
	CommentsCount       int32              `json:"commentsCount,omitempty"`
	Tags                []IssueTag         `json:"tags,omitempty"`
	Links               []IssueLink        `json:"links,omitempty"`
	ExternalIssue       *ExternalIssue     `json:"externalIssue,omitempty"`
	CustomFields        []IssueCustomField `json:"customFields,omitempty"`
	Voters              *IssueVoters       `json:"voters,omitempty"`
	Watchers            *IssueWatchers     `json:"watchers,omitempty"`
	Attachments         []IssueAttachment  `json:"attachments,omitempty"`
	Subtasks            *IssueLink         `json:"subtasks,omitempty"`
	Parent              *IssueLink         `json:"parent,omitempty"`
	Id                  string             `json:"id,omitempty"`
	Type_               string             `json:"$type,omitempty"`
}

// For an issue that was imported from another service, represents the reference to the issue in the external system.
type ExternalIssue struct {
	Name  string `json:"name,omitempty"`
	Url   string `json:"url,omitempty"`
	Key   string `json:"key,omitempty"`
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents a file that is attached to an issue or a comment.
type IssueAttachment struct {
	Name          string        `json:"name,omitempty"`
	Author        *User         `json:"author,omitempty"`
	Created       int64         `json:"created,omitempty"`
	Updated       int64         `json:"updated,omitempty"`
	Size          int64         `json:"size,omitempty"`
	Extension     string        `json:"extension,omitempty"`
	Charset       string        `json:"charset,omitempty"`
	MimeType      string        `json:"mimeType,omitempty"`
	MetaData      string        `json:"metaData,omitempty"`
	Draft         bool          `json:"draft,omitempty"`
	Removed       bool          `json:"removed,omitempty"`
	Base64Content string        `json:"base64Content,omitempty"`
	Url           string        `json:"url,omitempty"`
	Visibility    *Visibility   `json:"visibility,omitempty"`
	Issue         *Issue        `json:"issue,omitempty"`
	Comment       *IssueComment `json:"comment,omitempty"`
	ThumbnailURL  string        `json:"thumbnailURL,omitempty"`
	Id            string        `json:"id,omitempty"`
	Type_         string        `json:"$type,omitempty"`
}

// Swimlane settings for the case, when each swimlane is represented by an issue and contains this issue's subtasks.
type IssueBasedSwimlaneSettings struct {
	Enabled         bool            `json:"enabled,omitempty"`
	Id              string          `json:"id,omitempty"`
	Type_           string          `json:"$type,omitempty"`
	Field           *FilterField    `json:"field,omitempty"`
	DefaultCardType *SwimlaneValue  `json:"defaultCardType,omitempty"`
	Values          []SwimlaneValue `json:"values,omitempty"`
}

// Represents single swimlane in case of IssueBasedSwimlaneSettings.
type SwimlaneValue struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents an existing issue comment.
type IssueComment struct {
	Text         string            `json:"text,omitempty"`
	UsesMarkdown bool              `json:"usesMarkdown,omitempty"`
	TextPreview  string            `json:"textPreview,omitempty"`
	Created      int64             `json:"created,omitempty"`
	Updated      int64             `json:"updated,omitempty"`
	Author       *User             `json:"author,omitempty"`
	Issue        *Issue            `json:"issue,omitempty"`
	Attachments  []IssueAttachment `json:"attachments,omitempty"`
	Visibility   *Visibility       `json:"visibility,omitempty"`
	Deleted      bool              `json:"deleted,omitempty"`
	Id           string            `json:"id,omitempty"`
	Type_        string            `json:"$type,omitempty"`
	Reactions    []Reaction        `json:"reactions,omitempty"`
}

type Reaction struct {
	Id       string `json:"id,omitempty"`
	Type_    string `json:"$type,omitempty"`
	Reaction string `json:"reaction,omitempty"`
	Author   *User  `json:"author,omitempty"`
}

// Represents an event when a user creates a new issue.
type IssueCreatedActivityItem struct {
	Target  *Issue  `json:"target,omitempty"`
	Removed []Issue `json:"removed,omitempty"`
	Added   []Issue `json:"added,omitempty"`
}

// Represents an issue folder, such as a project, a saved search, or a tag.
type IssueFolder struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents issue links of a particular link type (for example, 'relates to').
type IssueLink struct {
	Direction     string         `json:"direction,omitempty"`
	LinkType      *IssueLinkType `json:"linkType,omitempty"`
	Issues        []Issue        `json:"issues,omitempty"`
	TrimmedIssues []Issue        `json:"trimmedIssues,omitempty"`
	Id            string         `json:"id,omitempty"`
	Type_         string         `json:"$type,omitempty"`
}

// Represents the settings of a link type in YouTrack.
type IssueLinkType struct {
	Name                    string `json:"name,omitempty"`
	LocalizedName           string `json:"localizedName,omitempty"`
	SourceToTarget          string `json:"sourceToTarget,omitempty"`
	LocalizedSourceToTarget string `json:"localizedSourceToTarget,omitempty"`
	TargetToSource          string `json:"targetToSource,omitempty"`
	LocalizedTargetToSource string `json:"localizedTargetToSource,omitempty"`
	Directed                bool   `json:"directed,omitempty"`
	Aggregation             bool   `json:"aggregation,omitempty"`
	ReadOnly                bool   `json:"readOnly,omitempty"`
	Id                      string `json:"id,omitempty"`
	Type_                   string `json:"$type,omitempty"`
}

// Represents an event when a user resolves an issue.
type IssueResolvedActivityItem struct {
	Removed int64 `json:"removed,omitempty"`
	Added   int64 `json:"added,omitempty"`
}

// Represents an issue tag.
type IssueTag struct {
	Id             string      `json:"id,omitempty"`
	Owner          *User       `json:"owner,omitempty"`
	VisibleFor     *UserGroup  `json:"visibleFor,omitempty"`
	UpdateableBy   *UserGroup  `json:"updateableBy,omitempty"`
	Issues         []Issue     `json:"issues,omitempty"`
	Color          *FieldStyle `json:"color,omitempty"`
	UntagOnResolve bool        `json:"untagOnResolve,omitempty"`
	Name           *string     `json:"name,omitempty"`
}

// Represents time tracking settings in the issue.
type IssueTimeTracker struct {
	WorkItems []IssueWorkItem `json:"workItems,omitempty"`
	Enabled   bool            `json:"enabled,omitempty"`
	Id        string          `json:"id,omitempty"`
	Type_     string          `json:"$type,omitempty"`
}

// Represents users that have voted for the issue or its duplicates.
type IssueVoters struct {
	HasVote   bool            `json:"hasVote,omitempty"`
	Original  []User          `json:"original,omitempty"`
	Duplicate []DuplicateVote `json:"duplicate,omitempty"`
	Id        string          `json:"id,omitempty"`
	Type_     string          `json:"$type,omitempty"`
}

// Represents a vote for duplicates of the issue.
type DuplicateVote struct {
	Issue *Issue `json:"issue,omitempty"`
	User  *User  `json:"user,omitempty"`
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents a user who subscribed for notifications about an issue.
type IssueWatcher struct {
	User      *User  `json:"user,omitempty"`
	Issue     *Issue `json:"issue,omitempty"`
	IsStarred bool   `json:"isStarred,omitempty"`
	Id        string `json:"id,omitempty"`
	Type_     string `json:"$type,omitempty"`
}

// Represents users that are subscribed to notifications about the issue.
type IssueWatchers struct {
	HasStar           bool           `json:"hasStar,omitempty"`
	IssueWatchers     []IssueWatcher `json:"issueWatchers,omitempty"`
	DuplicateWatchers []IssueWatcher `json:"duplicateWatchers,omitempty"`
	Id                string         `json:"id,omitempty"`
	Type_             string         `json:"$type,omitempty"`
}

// Represents a work item in an issue.
type IssueWorkItem struct {
	Author       *User          `json:"author,omitempty"`
	Creator      *User          `json:"creator,omitempty"`
	Text         string         `json:"text,omitempty"`
	TextPreview  string         `json:"textPreview,omitempty"`
	WorkType     *WorkItemType  `json:"type,omitempty"`
	Created      int64          `json:"created,omitempty"`
	Updated      int64          `json:"updated,omitempty"`
	Duration     *DurationValue `json:"duration,omitempty"`
	Date         int64          `json:"date,omitempty"`
	Issue        *Issue         `json:"issue,omitempty"`
	UsesMarkdown bool           `json:"usesMarkdown,omitempty"`
	Id           string         `json:"id,omitempty"`
	Type         string         `json:"$type,omitempty"`
}

// Describes visibility settings for an entity (issue, comment, etc.)
type Visibility struct {
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents an issue property, which can be a predefined field, a custom field, a link, and so on.
type FilterField struct {
	Name  string `json:"name,omitempty"`
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents the duration value and its visual presentation.
type DurationValue struct {
	Minutes      int32  `json:"minutes,omitempty"`
	Presentation string `json:"presentation,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}

type VcsChange struct {
	Id      string `json:"id,omitempty"`
	Date    int64  `json:"date,omitempty"`
	Fetched int64  `json:"fetched,omitempty"`
	Files   int    `json:"files,omitempty"`
	Author  *User  `json:"author,omitempty"`
	Text    string `json:"text,omitempty"`
	Version string `json:"version,omitempty"`
	Issue   *Issue `json:"issue,omitempty"`
	Type_   string `json:"$type,omitempty"`
}
