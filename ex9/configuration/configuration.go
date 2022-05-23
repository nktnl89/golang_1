package configuration

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/kelseyhightower/envconfig"
)

var (
	port        = flag.Int("port", 0, "")
	dbUrl       = flag.String("db_url", "", "")
	jaegerUrl   = flag.String("jaeger_url", "", "")
	sentryUrl   = flag.String("sentry_url", "", "")
	kafkaBroker = flag.String("kafka_broker", "", "")
	someAppId   = flag.String("some_app_id", "", "")
	someAppKey  = flag.String("some_app_key", "", "")
)

type Conf struct {
	Port        int    `yaml:"port" json:"port" envconfig:"port"`
	DbUrl       string `yaml:"db_url" json:"db_url" envconfig:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url" json:"jaeger_url" envconfig:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url" json:"sentry_url" envconfig:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker" json:"kafka_broker" envconfig:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id" json:"some_app_id" envconfig:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key" json:"some_app_key" envconfig:"some_app_key"`
}

func InitiateEnvironment() {
	os.Setenv("port", "8080")
	os.Setenv("db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	os.Setenv("jaeger_url", "http://jaeger:16686")
	os.Setenv("sentry_url", "http://sentry:9000")
	os.Setenv("kafka_broker", "kafka:9092")
	os.Setenv("some_app_id", "testid")
	os.Setenv("some_app_key", "testkey")
}

func configUrlsValidate(conf *Conf) {
	_, err := url.Parse(conf.SentryUrl)
	if err != nil {
		fmt.Println("Sentry url isn't valid")
	}
	_, err = url.Parse(conf.JaegerUrl)
	if err != nil {
		fmt.Println("Jaeger url isn't valid")
	}
	_, err = url.Parse(conf.DbUrl)
	if err != nil {
		fmt.Println("Db url isn't valid")
	}

}

func PopulateByFlag() Conf {

	flag.Parse()

	var conf = Conf{
		Port:        *port,
		DbUrl:       *dbUrl,
		JaegerUrl:   *jaegerUrl,
		KafkaBroker: *kafkaBroker,
		SentryUrl:   *sentryUrl,
		SomeAppId:   *someAppId,
		SomeAppKey:  *someAppKey,
	}

	configUrlsValidate(&conf)

	return conf

}

func PopulateByEnv(conf *Conf) {

	err := envconfig.Process("", conf)
	if err != nil {
		_ = fmt.Errorf("can't process the config: %w", err)
	}

	configUrlsValidate(conf)
}

func PopulateFromJson(filePath string) Conf {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	conf := Conf{}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		log.Printf("Не могу декодировать json-файл в структуру: %v", err)
	}

	configUrlsValidate(&conf)

	return conf

}

func PopulateFromYaml(filePath string) Conf {

	if _, err := os.Stat(filePath); err != nil {
		log.Fatalf("Не могу проверить существование файла: %v", err)
	}

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	conf := Conf{}

	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	configUrlsValidate(&conf)

	return conf
}
