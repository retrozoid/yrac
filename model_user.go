package yrac

// Represents a user in YouTrack. Please note that the read-only properties of a user account, like              credentials, or email and so on, you can only change in               <a href=\"https://www.jetbrains.com/help/youtrack/devportal/?Hub-REST-API\">Hub REST API</a>.
type User struct {
	Login             string        `json:"login,omitempty"`
	FullName          string        `json:"fullName,omitempty"`
	Email             string        `json:"email,omitempty"`
	JabberAccountName string        `json:"jabberAccountName,omitempty"`
	RingId            string        `json:"ringId,omitempty"`
	Guest             bool          `json:"guest,omitempty"`
	Online            bool          `json:"online,omitempty"`
	Banned            bool          `json:"banned,omitempty"`
	Tags              []IssueTag    `json:"tags,omitempty"`
	SavedQueries      []SavedQuery  `json:"savedQueries,omitempty"`
	AvatarUrl         string        `json:"avatarUrl,omitempty"`
	Profiles          *UserProfiles `json:"profiles,omitempty"`
	Id                string        `json:"id,omitempty"`
	Type_             string        `json:"$type,omitempty"`
}

// Represents a saved search.
type SavedQuery struct {
	Owner        *User      `json:"owner,omitempty"`
	VisibleFor   *UserGroup `json:"visibleFor,omitempty"`
	UpdateableBy *UserGroup `json:"updateableBy,omitempty"`
	Query        string     `json:"query,omitempty"`
	Issues       []Issue    `json:"issues,omitempty"`
}

// Represents a set of values that contains users. You can add to the set both individual user accounts and groups of users.
type UserBundle struct {
	IsUpdateable    bool        `json:"isUpdateable,omitempty"`
	Id              string      `json:"id,omitempty"`
	Type_           string      `json:"$type,omitempty"`
	Groups          []UserGroup `json:"groups,omitempty"`
	Individuals     []User      `json:"individuals,omitempty"`
	AggregatedUsers []User      `json:"aggregatedUsers,omitempty"`
}

// Represents default settings for the user-type field.
type UserCustomFieldDefaults struct {
	CanBeEmpty     bool         `json:"canBeEmpty,omitempty"`
	EmptyFieldText string       `json:"emptyFieldText,omitempty"`
	IsPublic       bool         `json:"isPublic,omitempty"`
	Parent         *CustomField `json:"parent,omitempty"`
	Id             string       `json:"id,omitempty"`
	Type_          string       `json:"$type,omitempty"`
	Bundle         *UserBundle  `json:"bundle,omitempty"`
	DefaultValues  []User       `json:"defaultValues,omitempty"`
}

// Represents a group of users.
type UserGroup struct {
	Id             string   `json:"id,omitempty"`
	Name           string   `json:"name,omitempty"`
	RingId         string   `json:"ringId,omitempty"`
	UsersCount     int64    `json:"usersCount,omitempty"`
	Icon           string   `json:"icon,omitempty"`
	AllUsersGroup  bool     `json:"allUsersGroup,omitempty"`
	TeamForProject *Project `json:"teamForProject,omitempty"`
}

type UserProfiles struct {
	General       *GeneralUserProfile       `json:"general,omitempty"`
	Notifications *NotificationsUserProfile `json:"notifications,omitempty"`
	TimeTracking  *TimeTrackingUserProfile  `json:"timetracking,omitempty"`
	Id            string                    `json:"id,omitempty"`
	Type_         string                    `json:"$type,omitempty"`
}

// Represents time tracking settings in the user's profile.
type TimeTrackingUserProfile struct {
	PeriodFormat *PeriodFieldFormat `json:"periodFormat,omitempty"`
	Id           string             `json:"id,omitempty"`
	Type_        string             `json:"$type,omitempty"`
}

// Represents the format, in which the period values are displayed to the user.
type PeriodFieldFormat struct {
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Represents the user profile in YouTrack.
type GeneralUserProfile struct {
	DateFieldFormat *DateFormatDescriptor `json:"dateFieldFormat,omitempty"`
	Timezone        *TimeZoneDescriptor   `json:"timezone,omitempty"`
	Locale          *LocaleDescriptor     `json:"locale,omitempty"`
	Id              string                `json:"id,omitempty"`
	Type_           string                `json:"$type,omitempty"`
}

// Represents a language locale that is used in UI.
type LocaleDescriptor struct {
	Locale    string `json:"locale,omitempty"`
	Language  string `json:"language,omitempty"`
	Community bool   `json:"community,omitempty"`
	Name      string `json:"name,omitempty"`
	Id        string `json:"id,omitempty"`
	Type_     string `json:"$type,omitempty"`
}

// Represents a time zone.
type TimeZoneDescriptor struct {
	Presentation string `json:"presentation,omitempty"`
	Offset       int32  `json:"offset,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}

// Represents date format.
type DateFormatDescriptor struct {
	Presentation string `json:"presentation,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
	DatePattern  string `json:"datePattern,omitempty"`
	Id           string `json:"id,omitempty"`
	Type_        string `json:"$type,omitempty"`
}

// Represents the notification settings in the profile of the specific user.
type NotificationsUserProfile struct {
	NotifyOnOwnChanges                     bool   `json:"notifyOnOwnChanges,omitempty"`
	JabberNotificationsEnabled             bool   `json:"jabberNotificationsEnabled,omitempty"`
	EmailNotificationsEnabled              bool   `json:"emailNotificationsEnabled,omitempty"`
	MentionNotificationsEnabled            bool   `json:"mentionNotificationsEnabled,omitempty"`
	DuplicateClusterNotificationsEnabled   bool   `json:"duplicateClusterNotificationsEnabled,omitempty"`
	MailboxIntegrationNotificationsEnabled bool   `json:"mailboxIntegrationNotificationsEnabled,omitempty"`
	UsePlainTextEmails                     bool   `json:"usePlainTextEmails,omitempty"`
	AutoWatchOnComment                     bool   `json:"autoWatchOnComment,omitempty"`
	AutoWatchOnCreate                      bool   `json:"autoWatchOnCreate,omitempty"`
	AutoWatchOnVote                        bool   `json:"autoWatchOnVote,omitempty"`
	AutoWatchOnUpdate                      bool   `json:"autoWatchOnUpdate,omitempty"`
	Id                                     string `json:"id,omitempty"`
	Type_                                  string `json:"$type,omitempty"`
}

// Represents project settings for the user field.
type UserProjectCustomField struct {
	Bundle        *UserBundle `json:"bundle,omitempty"`
	DefaultValues []User      `json:"defaultValues,omitempty"`
}

// Represents the change of the boolean flag that indicates whether YouTrack Wiki or Markdown is used in the target entity as a markup language. If \"true\", then the markdown is used. Otherwise, YouTrack Wiki markup.
type UsesMarkupActivityItem struct {
	Removed bool   `json:"removed,omitempty"`
	Added   bool   `json:"added,omitempty"`
	Markup  string `json:"markup,omitempty"`
}
