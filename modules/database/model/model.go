package model

type ListElemReqBody struct {
	LevelStack []string `json:"level_stack"`
	PageSize   int64    `validate:"gte=0,max=1000"`
	Page       int64    `validate:"gte=0"`
	SearchKey  string
}

type ListedElem struct {
	LevelStack []string `json:"level_stack"`
	SearchKey  string   `json:"search_key,omitempty"`
	Results    []Result `json:"results"`
}

type Result struct {
	Name          string   `json:"name"`
	IsBucket      bool     `json:"is_bucket"`
	Value         string   `json:"value,omitempty"`
	NoOfChildBkts int      `json:"no_of_child_bkts,omitempty"`
	NoOfPairs     int      `json:"no_of_pairs,omitempty"`
	ChildBkts     []string `json:"child_bkts,omitempty"`
	ChildKeys     []string `json:"child_keys,omitempty"`
}
