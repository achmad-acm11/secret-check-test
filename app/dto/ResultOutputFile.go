package dto

import "time"

type ResultOutputFile struct {
	SchemaVersion int    `json:"SchemaVersion"`
	ArtifactName  string `json:"ArtifactName"`
	ArtifactType  string `json:"ArtifactType"`
	Metadata      struct {
		ImageConfig struct {
			Architecture string    `json:"architecture"`
			Created      time.Time `json:"created"`
			Os           string    `json:"os"`
			Rootfs       struct {
				Type    string      `json:"type"`
				DiffIds interface{} `json:"diff_ids"`
			} `json:"rootfs"`
			Config struct {
			} `json:"config"`
		} `json:"ImageConfig"`
	} `json:"Metadata"`
	Results []ScanSecretResult `json:"Results"`
}

type ScanSecretResult struct {
	Target  string       `json:"Target"`
	Class   string       `json:"Class"`
	Secrets []SecretData `json:"Secrets"`
}

type SecretData struct {
	RuleID    string     `json:"RuleID"`
	Category  string     `json:"Category"`
	Severity  string     `json:"Severity"`
	Title     string     `json:"Title"`
	StartLine int        `json:"StartLine"`
	EndLine   int        `json:"EndLine"`
	Code      SecretCode `json:"Code"`
	Match     string     `json:"Match"`
	Layer     struct {
	} `json:"Layer"`
}

type SecretCode struct {
	Lines []struct {
		Number      int    `json:"Number"`
		Content     string `json:"Content"`
		IsCause     bool   `json:"IsCause"`
		Annotation  string `json:"Annotation"`
		Truncated   bool   `json:"Truncated"`
		Highlighted string `json:"Highlighted,omitempty"`
		FirstCause  bool   `json:"FirstCause"`
		LastCause   bool   `json:"LastCause"`
	} `json:"Lines"`
}
