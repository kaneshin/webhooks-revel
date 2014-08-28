package entities

type SlackData struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	IconUrl   string `json:"icon_url"`
	IconEmoji string `json:"icon_emoji"`
}

func NewSlackData() *SlackData {
	data := &SlackData{}
	return data
}

func (self *SlackData) InitializeSlackData(ch string, text string) {
	self.Channel = "#" + ch
	self.Text = text
}
