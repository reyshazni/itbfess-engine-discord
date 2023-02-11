package Entity

type Match struct {
	ChannelIdFirst  string `gorm:"primarykey"`
	ChannelIdSecond string `gorm:"primarykey"`
}
