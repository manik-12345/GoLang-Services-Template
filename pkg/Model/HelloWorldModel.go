package Model

type HelloWorldModel struct {
	HelloTitle string `gorm:"column:hello_title;type:varchar(300);" json:"hello-title"`

	WorldText string `gorm:"column:world_text;type:varchar(300)" json:"world-text"`
}
