package yrac

type Bundle struct {
	IsUpdateable bool `json:"isUpdateable,omitempty"`
}

type BundleElement struct {
	Name          string      `json:"name,omitempty"`
	Bundle        *Bundle     `json:"bundle,omitempty"`
	Description   string      `json:"description,omitempty"`
	Ordinal       int         `json:"ordinal,omitempty"`
	Color         *FieldStyle `json:"color,omitempty"`
	HasRunningJob bool        `json:"hasRunningJob,omitempty"`
}

type LocalizableBundleElement struct {
	BundleElement
	LocalizedName string `json:"localizedName,omitempty"`
}

type StateBundleElement struct {
	LocalizableBundleElement
	IsResolved bool `json:"isResolved,omitempty"`
}

func NewStateBundleElement(name string) *StateBundleElement {
	return &StateBundleElement{
		LocalizableBundleElement: LocalizableBundleElement{
			BundleElement: BundleElement{
				Name: name,
			},
		},
	}
}

type BuildBundleElement struct {
	BundleElement
	AssembleDate int64 `json:"assembleDate,omitempty"`
}

type VersionBundleElement struct {
	BundleElement
	Archived    bool  `json:"archived,omitempty"`
	ReleaseDate int64 `json:"releaseDate,omitempty"`
	Released    bool  `json:"released,omitempty"`
}

type OwnedBundleElement struct {
	BundleElement
	Owner *User `json:"owner,omitempty"`
}

type EnumBundleElement LocalizableBundleElement

func NewEnumBundleElement(name string) *EnumBundleElement {
	return &EnumBundleElement{
		BundleElement: BundleElement{
			Name: name,
		},
	}
}

type TextFieldValue struct {
	Text         *string `json:"text,omitempty"`
	MarkdownText string  `json:"markdownText"`
}

type PeriodValue struct {
	Minutes      int    `json:"minutes"`
	Presentation string `json:"presentation"`
}
