package kafkaconsumer

type Config struct {
	Brokers   []string `mapstructure:"broker" env:"CONSUMER_BROKERS"`
	Topic     string   `mapstructure:"topic" env:"CONSUMER_TOPIC"`
	Partition int      `mapstructure:"partition" env:"CONSUMER_PARTITION"`
	GroupID     string   `mapstructure:"group" env:"CONSUMER_GROUP"`
}
