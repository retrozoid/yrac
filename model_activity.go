package yrac

import (
	"encoding/json"
)

type ActivityItem struct {
	Id        string             `json:"id,omitempty"`
	Author    *User              `json:"author,omitempty"`
	Timestamp int64              `json:"timestamp,omitempty"`
	Field     *CustomFilterField `json:"field"`
	Type      string             `json:"$type,omitempty"`
	Removed   interface{}        `json:"removed"`
	Added     interface{}        `json:"added"`
	Target    interface{}        `json:"target"`
}

func (activityItem *ActivityItem) UnmarshalJSON(data []byte) error {
	type alias ActivityItem
	var (
		a   alias
		err error
	)
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}
	switch a.Type {
	case "TextMarkupActivityItem", "IssueResolvedActivityItem", "UsesMarkupActivityItem", "TextCustomFieldActivityItem", "SimpleValueActivityItem":
	// 		it's primitive type interface{}
	case "CustomFieldActivityItem":
		switch a.Field.CustomField.FieldType.Id {
		case "string", "integer", "float", "period", "date", "date and time", "text":
		// 		it's primitive type interface{}
		case "build[1]", "build[*]":
			a.Added = &[]BuildBundleElement{}
			a.Removed = &[]BuildBundleElement{}
		case "enum[1]", "enum[*]":
			a.Added = &[]EnumBundleElement{}
			a.Removed = &[]EnumBundleElement{}
		case "version[1]", "version[*]":
			a.Added = &[]VersionBundleElement{}
			a.Removed = &[]VersionBundleElement{}
		case "user[1]", "user[*]":
			a.Added = &[]User{}
			a.Removed = &[]User{}
		case "ownedField[1]", "ownedField[*]":
			a.Added = &[]OwnedBundleElement{}
			a.Removed = &[]OwnedBundleElement{}
		case "group[1]", "group[*]":
			a.Added = &[]UserGroup{}
			a.Removed = &[]UserGroup{}
		case "state[1]":
			a.Added = &[]StateBundleElement{}
			a.Removed = &[]StateBundleElement{}
		default:
			panic("unimplemented  type of CustomFieldActivityItem")
		}
	case "LinksActivityItem", "IssueCreatedActivityItem":
		a.Added = &[]Issue{}
		a.Removed = &[]Issue{}
	case "TagsActivityItem":
		a.Added = &[]IssueTag{}
		a.Removed = &[]IssueTag{}
	case "CommentActivityItem":
		a.Added = &[]IssueComment{}
		a.Removed = &[]IssueComment{}
	case "AttachmentActivityItem", "CommentAttachmentsActivityItem":
		a.Added = &[]IssueAttachment{}
		a.Removed = &[]IssueAttachment{}
	case "WorkItemAuthorActivityItem", "VisibilityUserActivityItem", "VotersActivityItem":
		a.Added = &[]User{}
		a.Removed = &[]User{}
	case "WorkItemActivityItem":
		a.Added = &[]IssueWorkItem{}
		a.Removed = &[]IssueWorkItem{}
	case "WorkItemDurationActivityItem":
		a.Added = &PeriodValue{}
		a.Removed = &PeriodValue{}
	case "VisibilityGroupActivityItem":
		a.Added = &[]UserGroup{}
		a.Removed = &[]UserGroup{}
	case "VcsChangeActivityItem":
		a.Added = &[]VcsChange{}
		a.Removed = &[]VcsChange{}
	case "CommentReactionActivityItem":
		a.Added = &[]Reaction{}
		a.Removed = &[]Reaction{}
		a.Target = &IssueComment{}

	default:
		panic("unimplemented activity item type")

	}
	if err = json.Unmarshal(data, &a); err != nil {
		return err
	}
	*activityItem = ActivityItem(a)
	return nil
}
