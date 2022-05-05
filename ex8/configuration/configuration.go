package configuration

import (
	"flag"
	"fmt"
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
	Port        int    `envconfig:"port"`
	DbUrl       string `envconfig:"db_url"`
	JaegerUrl   string `envconfig:"jaeger_url"`
	SentryUrl   string `envconfig:"sentry_url"`
	KafkaBroker string `envconfig:"kafka_broker"`
	SomeAppId   string `envconfig:"some_app_id"`
	SomeAppKey  string `envconfig:"some_app_key"`
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

func PopulateByFlag() Conf {

	flag.Parse()

	return Conf{
		Port:        *port,
		DbUrl:       *dbUrl,
		JaegerUrl:   *jaegerUrl,
		KafkaBroker: *kafkaBroker,
		SentryUrl:   *sentryUrl,
		SomeAppId:   *someAppId,
		SomeAppKey:  *someAppKey,
	}

}

func PopulateByEnv(conf *Conf) {

	err := envconfig.Process("", conf)
	if err != nil {
		fmt.Errorf("can't process the config: %w", err)
	}
}
