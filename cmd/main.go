package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gabrielmq/kafka-consumer-go/infra/kafka"
	"github.com/gabrielmq/kafka-consumer-go/infra/repository"
	"github.com/gabrielmq/kafka-consumer-go/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file")
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/fullcycle")
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewCourseMySQLRepository(db)
	ucase := usecase.NewCreateCourse(repository)

	var msgChan = make(chan *ckafka.Message)

	consumer := kafka.NewConsumer(
		&ckafka.ConfigMap{
			"bootstrap.servers": os.Getenv("BOOTSTRAP_SERVERS"),
			"group.id":          os.Getenv("GROUP_ID"),
		},
		strings.Split(os.Getenv("TOPICS"), ","),
	)

	go consumer.Consume(msgChan)
	for msg := range msgChan {
		var input usecase.CreateCourseInputDto
		json.Unmarshal(msg.Value, &input)

		output, err := ucase.Execute(input)
		if err != nil {
			fmt.Println("error:", err)
		} else {
			fmt.Println(output)
		}
	}
}
