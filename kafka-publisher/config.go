package kafkapublisher
type Config struct {
	Brokers []string `json:"brokers" env:"PUBLISHER_BROKERS"`
	Topic   string   `json:"topic" env:"PUBLISHER_TOPIC"`
	GroupID string   `json:"group_id" env:"PUBLISHER_GROUP_ID"`
}
