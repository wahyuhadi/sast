package models

type Results struct {
	Errors []any `json:"errors"`
	Paths  struct {
		Comment string   `json:"_comment"`
		Scanned []string `json:"scanned"`
	} `json:"paths"`
	Results []struct {
		CheckID string `json:"check_id"`
		End     struct {
			Col    int `json:"col"`
			Line   int `json:"line"`
			Offset int `json:"offset"`
		} `json:"end"`
		Extra struct {
			EngineKind  string `json:"engine_kind"`
			Fingerprint string `json:"fingerprint"`
			IsIgnored   bool   `json:"is_ignored"`
			Lines       string `json:"lines"`
			Message     string `json:"message"`
			Metadata    struct {
				Category   string   `json:"category"`
				Confidence string   `json:"confidence"`
				Technology []string `json:"technology"`
			} `json:"metadata"`
			Metavars struct {
			} `json:"metavars"`
			Severity        string `json:"severity"`
			ValidationState string `json:"validation_state"`
		} `json:"extra"`
		Path  string `json:"path"`
		Start struct {
			Col    int `json:"col"`
			Line   int `json:"line"`
			Offset int `json:"offset"`
		} `json:"start"`
	} `json:"results"`
	Version string `json:"version"`
}

type Finding struct {
	ScanAt          string  `json:"scan_at"`
	Branch          string  `json:"branch"`
	IsFinding       bool    `json:"is_finding"`
	Key             string  `json:"key"`
	Date            string  `json:"date"`
	CheckID         string  `json:"check_id"`
	EndCol          int     `json:"end_col"`
	EndLine         int     `json:"end_line"`
	EndOffset       int     `json:"end_offset"`
	Project         string  `json:"project"`
	EngineKind      string  `json:"engine_kind"`
	Fingerprint     string  `json:"fingerprint"`
	IsIgnored       bool    `json:"is_ignored"`
	Lines           string  `json:"lines"`
	Message         string  `json:"message"`
	Category        string  `json:"category"`
	Confidence      string  `json:"confidence"`
	Lang            string  `json:"lang"`
	Technology      *string `json:"technology"`
	UniqueID        string  `json:"unique_id"`
	Severity        string  `json:"severity"`
	ValidationState string  `json:"validation_state"`
	Path            string  `json:"path"`
	StartCol        int     `json:"start_col"`
	StartLine       int     `json:"start_line"`
	StartOffset     int     `json:"start_offset"`
}
