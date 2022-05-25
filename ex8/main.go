package main

import (
	"fmt"
	"golang_1/ex8/configuration"
)

func main() {
	configuration.InitiateEnvironment()

	envConf := configuration.Conf{}
	configuration.PopulateByEnv(&envConf)
	fmt.Println(envConf)

	//go run main.go --port=8080 --db_url=postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable --jaeger_url=http://jaeger:16686 --sentry_url=http://sentry:9000 --kafka_broker=kafka:9092 --some_app_id=testid --some_app_key=testkey
	fmt.Println(configuration.PopulateByFlag())
}
