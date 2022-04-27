package producers

import (
	"encoding/json"
	"fmt"
	"log"
    "math/rand"
    "time"
	"os"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	// "strconv"
)

type Result struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
	Date_game string `json:"date_game"`
}

func SaveToKafka(gameId, players int32, date string) string {
	result := getResult(gameId, players, date)
	fmt.Println(string(result))

	// Produce messages to topic (asynchronously)
	host := os.Getenv("KAFKA_HOST")
	nvoProducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	if err != nil {
		log.Fatalf("Error:no se pudo instanciar el producer :(. Error: %s", err.Error())
		panic(err)
	}

	topic := "games-topic"
	nvoProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(result),
	}, nil)

	return string(result)
}

func getResult(gameId, players int32, date string) string {

	var winner int

	//aca iria el if de los 5 juegos

	// rand.Seed(time.Now().UnixNano())
    // min := 1
    // max := 5
    // game := rand.Intn(max - min + 1) + min

	rand.Seed(time.Now().Unix())
	winner = rand.Intn(int(players))

	jsonBytes, err := json.Marshal(Result{
		Game_id:   gameId,
		Players:   players,
		Game_name: "Juego",
		Winner:    int32(winner),
		Queue:     "kafka",
		Date_game: date,
	})

	if err != nil {
		log.Fatalf("Error: No se puede Convertir struct a json. Error: %s", err.Error())
		panic(err)
	}
	return string(jsonBytes)
}
