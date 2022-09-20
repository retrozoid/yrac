package yrac

import (
	"encoding/json"
)

type IssueCustomField struct {
	Id                 string              `json:"id,omitempty"`
	Type               string              `json:"$type,omitempty"`
	Name               string              `json:"name,omitempty"`
	ProjectCustomField *ProjectCustomField `json:"projectCustomField,omitempty"`
	Value              interface{}         `json:"value"`
}

func NewIssueCustomField(id string, value interface{}) IssueCustomField {
	return IssueCustomField{
		Id:    id,
		Value: value,
	}
}

type IssueCustomFields []IssueCustomField

func (fields IssueCustomFields) Field(name string) *IssueCustomField {
	for i := 0; i < len(fields); i++ {
		if fields[i].Name == name {
			return &fields[i]
		}
	}
	return nil
}

func (field IssueCustomField) MarshalJSON() ([]byte, error) {
	switch field.Value.(type) {
	case string, bool, int, int8, int16, int32, int64, float32, float64,
		*string, *bool, *int, *int8, *int16, *int32, *int64, *float32, *float64:
		field.Type = "SimpleIssueCustomField"
	case *TextFieldValue, TextFieldValue:
		field.Type = "TextIssueCustomField"
	case *PeriodValue, PeriodValue:
		field.Type = "PeriodIssueCustomField"
	case *BuildBundleElement, BuildBundleElement:
		field.Type = "SingleBuildIssueCustomField"
	case *EnumBundleElement, EnumBundleElement:
		field.Type = "SingleEnumIssueCustomField"
	case *UserGroup, UserGroup:
		field.Type = "SingleGroupIssueCustomField"
	case *OwnedBundleElement, OwnedBundleElement:
		field.Type = "SingleOwnedIssueCustomField"
	case *User, User:
		field.Type = "SingleUserIssueCustomField"
	case *VersionBundleElement, VersionBundleElement:
		field.Type = "SingleVersionIssueCustomField"
	case *StateBundleElement, StateBundleElement:
		field.Type = "StateIssueCustomField"
	case *[]BuildBundleElement, []BuildBundleElement:
		field.Type = "MultiBuildIssueCustomField"
	case *[]EnumBundleElement, []EnumBundleElement:
		field.Type = "MultiEnumIssueCustomField"
	case *[]UserGroup, []UserGroup:
		field.Type = "MultiGroupIssueCustomField"
	case *[]OwnedBundleElement, []OwnedBundleElement:
		field.Type = "MultiOwnedIssueCustomField"
	case *[]User, []User:
		field.Type = "MultiUserIssueCustomField"
	case *[]VersionBundleElement, []VersionBundleElement:
		field.Type = "MultiVersionIssueCustomField"
	}
	type alias IssueCustomField
	return json.Marshal(alias(field))
}

func (field *IssueCustomField) UnmarshalJSON(data []byte) error {
	type alias IssueCustomField
	var (
		a   alias
		err error
	)
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}
	switch a.Type {
	case "SimpleIssueCustomField", "DateIssueCustomField":
		/* it's primitive type interface{} */
	case "TextIssueCustomField":
		a.Value = &TextFieldValue{}
	case "PeriodIssueCustomField":
		a.Value = &PeriodValue{}

	case "SingleBuildIssueCustomField":
		a.Value = &BuildBundleElement{}
	case "SingleEnumIssueCustomField":
		a.Value = &EnumBundleElement{}
	case "SingleGroupIssueCustomField":
		a.Value = &UserGroup{}
	case "SingleOwnedIssueCustomField":
		a.Value = &OwnedBundleElement{}
	case "SingleUserIssueCustomField":
		a.Value = &User{}
	case "SingleVersionIssueCustomField":
		a.Value = &VersionBundleElement{}
	case "StateIssueCustomField":
		a.Value = &StateBundleElement{}

	case "MultiBuildIssueCustomField":
		a.Value = &[]BuildBundleElement{}
	case "MultiEnumIssueCustomField":
		a.Value = &[]EnumBundleElement{}
	case "MultiGroupIssueCustomField":
		a.Value = &[]UserGroup{}
	case "MultiOwnedIssueCustomField":
		a.Value = &[]OwnedBundleElement{}
	case "MultiUserIssueCustomField":
		a.Value = &[]User{}
	case "MultiVersionIssueCustomField":
		a.Value = &[]VersionBundleElement{}
	default:
		panic("unimplemented custom field type " + a.Type)

	}
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}
	*field = IssueCustomField(a)
	return nil
}

type CustomField struct {
	Name                   string               `json:"name,omitempty"`
	LocalizedName          string               `json:"localizedName,omitempty"`
	FieldType              *FieldType           `json:"fieldType,omitempty"`
	IsAutoAttached         bool                 `json:"isAutoAttached,omitempty"`
	IsDisplayedInIssueList bool                 `json:"isDisplayedInIssueList,omitempty"`
	Ordinal                int32                `json:"ordinal,omitempty"`
	Aliases                string               `json:"aliases,omitempty"`
	FieldDefaults          *CustomFieldDefaults `json:"fieldDefaults,omitempty"`
	HasRunningJob          bool                 `json:"hasRunningJob,omitempty"`
	IsUpdateable           bool                 `json:"isUpdateable,omitempty"`
	Instances              []ProjectCustomField `json:"instances,omitempty"`
	Id                     string               `json:"id,omitempty"`
	Type_                  string               `json:"$type,omitempty"`
}

// Represents the style settings of the field in YouTrack.
type FieldStyle struct {
	Background string `json:"background,omitempty"`
	Foreground string `json:"foreground,omitempty"`
	Id         string `json:"id,omitempty"`
	Type_      string `json:"$type,omitempty"`
}

// Represents the type of a custom field in YouTrack.
type FieldType struct {
	Id    string `json:"id,omitempty"`
	Type_ string `json:"$type,omitempty"`
}

// Sets a card color based on a value of some custom field.
type FieldBasedColorCoding struct {
	Id        string       `json:"id,omitempty"`
	Type_     string       `json:"$type,omitempty"`
	Prototype *CustomField `json:"prototype,omitempty"`
}

// Represents an activity that affects a custom field of an issue.
type CustomFieldActivityItem struct {
	Id      string      `json:"id,omitempty"`
	Type_   string      `json:"$type,omitempty"`
	Target  *Issue      `json:"target,omitempty"`
	Removed interface{} `json:"removed,omitempty"`
	Added   interface{} `json:"added,omitempty"`
}

// Represents default project-related settings of the custom field. These settings are applied at the moment when the custom field is attached to a project. After that, any changes in default settings do not affect the field settings for this project.
type CustomFieldDefaults struct {
	CanBeEmpty     bool         `json:"canBeEmpty,omitempty"`
	EmptyFieldText string       `json:"emptyFieldText,omitempty"`
	IsPublic       bool         `json:"isPublic,omitempty"`
	Parent         *CustomField `json:"parent,omitempty"`
	Id             string       `json:"id,omitempty"`
	Type_          string       `json:"$type,omitempty"`
}

// Represents a custom field of the issue.
type CustomFilterField struct {
	Name        string       `json:"name,omitempty"`
	Id          string       `json:"id,omitempty"`
	Type_       string       `json:"$type,omitempty"`
	CustomField *CustomField `json:"customField,omitempty"`
}
