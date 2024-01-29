package model

type BlockHeight struct {
	Id     int64  `json:"id" gorm:"type:int;comment:id"`
	Chain  string `json:"chain"`
	Height int64  `json:"height"`
}

func (BlockHeight) TableName() string {
	return "block_height"
}
