package internal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	Host         string `mapstructure:"DB_HOST"`
	Port         string `mapstructure:"DB_PORT"`
	User         string `mapstructure:"DB_USER"`
	Password     string `mapstructure:"DB_PASSWORD"`
	DB           string `mapstructure:"DB_DB"`
	CommandTable string `mapstructure:"COMMAND_TABLE"`
	HeaderTable  string `mapstructure:"HEADER_TABLE"`
}

type Header struct {
	Method      string `json:"method" bson:"method"`
	Account     string `json:"account,omitempty" bson:"account,omitempty"`
	Password    string `json:"password,omitempty" bson:"password,omitempty"`
	Token       string `json:"token,omitempty" bson:"token,omitempty"`
	ContentType string `json:"content_type,omitempty" bson:"content_type,omitempty"`
}

type Protocol struct {
	Protocol string   `json:"protocol" bson:"protocol"`
	Method   string   `json:"method" bson:"method"`
	URL      string   `json:"url" bson:"url"`
	Headers  Header   `json:"headers" bson:"headers"`
	SendData SendData `json:"send_data,omitempty" bson:"send_data,omitempty"`
}

type SendData struct {
	DataType string      `json:"data_type,omitempty" bson:"data_type,omitempty"`
	Data     interface{} `json:"data,omitempty" bson:"data,omitempty"`
}

type Insert struct {
	Check         bool `json:"check" bson:"check"`
	ParseTemplate int  `json:"parse_template,omitempty" bson:"parse_template,omitempty"`
}

type ReadCommandTemplate struct {
	ID          primitive.ObjectID `json:"_id" json:"id,omitempty" bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Protocol    Protocol           `json:"protocol" bson:"protocol"`
	Insert      Insert             `json:"insert" bson:"insert"`
}
type WriteCommandTemplate struct {
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Protocol    Protocol `json:"protocol" bson:"protocol"`
	Insert      Insert   `json:"insert" bson:"insert"`
}
