package model

type ListElemReqBody struct {
	LevelStack []string `json:"level_stack"`
	PageSize   int64    `validate:"gte=0,max=1000"`
	Page       int64    `validate:"gte=0"`
	Key        string
}

type ListedElem struct {
	LevelStack    []string `json:"level_stack"`
	Key           string
	IsBucket      bool
	NoOfChildBkts int
	NoOfPairs     int
	ChildBkts     []string
	ChildKeys     []string
	Value         string
}
