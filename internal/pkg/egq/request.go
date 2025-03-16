package egq

type RequestRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Request ...
type Request struct {
	Page      int                     `json:"page" query:"page"`
	PageSize  int                     `json:"page_size" query:"page_size"`
	Sort      string                  `query:"sort"`
	Search    string                  `query:"search"`
	Filters   map[string]any          `query:"filters"`
	DateRange map[string]RequestRange `query:"dateRange"`
	Fields    []string                `query:"fields"`
	Includes  []string                `query:"includes"`
}
