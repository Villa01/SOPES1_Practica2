package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resul struct {
	Game_id   int32  `json:"game_id"`
	Players   int32  `json:"players"`
	Game_name string `json:"game_name"`
	Winner    int32  `json:"winner"`
	Queue     string `json:"queue"`
	Date_game string `json:"date_game"`
}

type GameJson struct {
	Gameid   int32  `json:gameid`
	Players  int32  `json:players`
	Gamename string `json:gamename`
	Winner   int32  `json:winner`
	Queue    string `json:queue`
	Datetime string
}

func GetCollection(collection string) *mongo.Collection {

	host := os.Getenv("MONGO_HOST")

	// userName := "root"
	// pass := "password"
	// if userName == "" || pass == "" || host == "" {
	// 	panic("Error: No se pudo conectar con mongodb, credeciales invalidas en el entorno")
	// }
	// URI := "mongodb://" + userName + ":" + pass + "@" + host + ":27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database("ProyectoF2").Collection(collection)
}



func InsertMongo(R string) (e error) {
	var r Resul
	err := json.Unmarshal([]byte(R), &r)
	if err != nil {
		return err
	}

	data := GameJson{
		Gameid:   r.Game_id,
		Players:  r.Players,
		Gamename: r.Game_name,
		Winner:   r.Winner,
		Queue:    r.Queue,
		Datetime: r.Date_game,
	}

	// data := struct{data string}{R}
	var coleccion = GetCollection("Logs")
	coleccion.InsertOne(context.Background(), data)

	fmt.Print("Se ha Guardado en Mongo :" + R + "\n\n")
	return nil
}
