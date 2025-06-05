package yrac

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func toValues(fields string) url.Values {
	v := url.Values{}
	if fields != "" {
		v["fields"] = []string{fields}
	}
	return v
}

func (s Service) QueryIssues(ctx context.Context, query string, top int, skip int, fields string) ([]Issue, error) {
	var issues []Issue
	var p = url.Values{}
	p.Add("query", query)
	p.Add("$top", fmt.Sprintf("%d", top))
	p.Add("$skip", fmt.Sprintf("%d", skip))
	p.Add("fields", fields)
	err := s.call(ctx, http.MethodGet, "/issues", p, nil, &issues)
	return issues, err
}

func (s Service) GetIssue(ctx context.Context, id string, fields string) (Issue, error) {
	var issue Issue
	var err = s.call(ctx, http.MethodGet, fmt.Sprintf("/issues/%s", id), toValues(fields), nil, &issue)
	return issue, err
}

func (s Service) CreateIssue(ctx context.Context, issue Issue, fields string) (Issue, error) {
	var val Issue
	var err = s.call(ctx, http.MethodPost, "/issues/", toValues(fields), issue, &val)
	return val, err
}

func (s Service) CreateIssueSilently(ctx context.Context, issue Issue, fields string) (Issue, error) {
	var val Issue
	values := toValues(fields)
	values.Add("muteUpdateNotifications", "true")
	var err = s.call(ctx, http.MethodPost, "/issues/", values, issue, &val)
	return val, err
}

func (s Service) UpdateIssue(ctx context.Context, id string, issue Issue, fields string) (Issue, error) {
	var val Issue
	var err = s.call(ctx, http.MethodPost, fmt.Sprintf("/issues/%s", id), toValues(fields), issue, &val)
	return val, err
}

func (s Service) GetIssueComments(ctx context.Context, id string, top int, skip int, fields string) ([]IssueComment, error) {
	var val []IssueComment
	var p = url.Values{}
	p.Add("$top", fmt.Sprintf("%d", top))
	p.Add("$skip", fmt.Sprintf("%d", skip))
	p.Add("fields", fields)
	var err = s.call(ctx, http.MethodGet, fmt.Sprintf("/issues/%s/comments", id), p, nil, &val)
	return val, err
}

func (s Service) CreateIssueComment(ctx context.Context, id string, c IssueComment, fields string) (IssueComment, error) {
	var val IssueComment
	var err = s.call(ctx, http.MethodPost, fmt.Sprintf("/issues/%s/comments", id), toValues(fields), c, &val)
	return val, err
}

func (s Service) AddIssueCommentReaction(ctx context.Context, id string, commentId string, r Reaction, fields string) (Reaction, error) {
	var val Reaction
	var err = s.call(ctx, http.MethodPost, fmt.Sprintf("/issues/%s/comments/%s/reactions", id, commentId), toValues(fields), r, &val)
	return val, err
}

func (s Service) UpdateIssueComment(ctx context.Context, id string, commentId string, c IssueComment, fields string) (IssueComment, error) {
	var val IssueComment
	var err = s.call(ctx, http.MethodPost, fmt.Sprintf("/issues/%s/comments/%s", id, commentId), toValues(fields), c, &val)
	return val, err
}

func (s Service) CreateIssueCommentAttachment(ctx context.Context, id string, commentId string, a IssueAttachment, fields string) (IssueAttachment, error) {
	var val IssueAttachment
	var err = s.call(ctx, http.MethodPost, fmt.Sprintf("/issues/%s/comments/%s/attachments", id, commentId), toValues(fields), a, &val)
	return val, err
}

func (s Service) GetIssueCustomFields(ctx context.Context, id string, fields string) (IssueCustomFields, error) {
	var val IssueCustomFields
	err := s.call(ctx, http.MethodGet, fmt.Sprintf("/issues/%s/customFields", id), toValues(fields), nil, &val)
	return val, err
}

func (s Service) ApplyCommand(ctx context.Context, command Command) (Command, error) {
	var val Command
	err := s.call(ctx, http.MethodPost, "/commands", toValues(""), command, &val)
	return val, err
}

func (s Service) GetIssueActivities(ctx context.Context,
	id string, top, skip int,
	fields, categories, author, issueQuery string,
	start, end int64,
	reverse bool) ([]ActivityItem, error) {

	var val []ActivityItem
	var p = url.Values{
		"$top":       {fmt.Sprintf("%d", top)},
		"$skip":      {fmt.Sprintf("%d", skip)},
		"fields":     {fields},
		"categories": {categories},
		"reverse":    {strconv.FormatBool(reverse)},
	}
	if author != "" {
		p.Add("author", author)
	}
	if issueQuery != "" {
		p.Add("issueQuery", issueQuery)
	}
	if start > 0 {
		p.Add("start", fmt.Sprintf("%d", start))
	}
	if end > 0 {
		p.Add("end", fmt.Sprintf("%d", end))
	}
	var err = s.call(ctx, http.MethodGet, fmt.Sprintf("/issues/%s/activities", id), p, nil, &val)
	return val, err
}

func (s Service) GetUser(ctx context.Context, id string, fields string) (User, error) {
	var val User
	var p = url.Values{}
	p.Add("fields", fields)
	var err = s.call(ctx, http.MethodGet, fmt.Sprintf("/users/%s", id), p, nil, &val)
	return val, err
}

func (s Service) GetUsers(ctx context.Context, query string, top int, skip int, fields string) ([]User, error) {
	var val []User
	var p = url.Values{}
	p.Add("query", query)
	p.Add("$top", fmt.Sprintf("%d", top))
	p.Add("$skip", fmt.Sprintf("%d", skip))
	p.Add("fields", fields)
	var err = s.call(ctx, http.MethodGet, "/users", p, nil, &val)
	return val, err
}

func (s Service) GetGroupUsers(ctx context.Context, id string, top int, skip int, fields string) ([]User, error) {
	var val []User
	var p = url.Values{}
	p.Add("$top", fmt.Sprintf("%d", top))
	p.Add("$skip", fmt.Sprintf("%d", skip))
	p.Add("fields", fields)
	var err = s.call(ctx, http.MethodGet, fmt.Sprintf("/groups/%s/users", id), p, nil, &val)
	return val, err
}
