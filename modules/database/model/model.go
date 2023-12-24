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

type ItemToDelete struct {
	LevelStack []string `json:"level_stack"`
	Key        string   `json:"key"`
}

type PairsToAdd struct {
	LevelStack []string `json:"level_stack"`
	Pairs      []Pair   `json:"pairs"`
}

type BucketsToAdd struct {
	LevelStack []string `json:"level_stack"`
	Buckets    []string `json:"buckets"`
}

type Pair struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type ItemToUpdate struct {
	LevelStack []string    `json:"level_stack"`
	Key        string      `json:"key"`
	NewValue   interface{} `json:"new_value,omitempty"`
}

type ItemToRename struct {
	LevelStack []string `json:"level_stack"`
	Key        string   `json:"key"`
	NewKey     string   `json:"new_key,omitempty"`
}
