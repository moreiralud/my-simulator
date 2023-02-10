package main

import (
	"github.com/joho/godotenv"
	"log"
	"my-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loaddin .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	for {
		_ = 1
	}

	//route := route2.Route{
	//	ID:       "1",
	//	ClientID: "1",
	//}
	//route.LoadPositions()
	//stringjson, _ := route.ExportJsonPositions()
	//fmt.Println(stringjson[1])
}
