package main

import (
	"log"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/mattn/go-gimei"
)

type Player struct {
	ID   int    `msg:"id"`
	Name string `msg:"name"`
}

func main() {
	logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "127.0.0.1"})
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer logger.Close()

	tag := "debug.player"

	for i := 1; i < 100; i++ {
		p := Player{
			ID:   i,
			Name: gimei.NewName().Kanji(),
		}
		logObj := make(map[string]interface{})
		logObj["player.name"] = p.Name
		logObj["player.id"] = p.ID

		err := logger.Post(tag, logObj)
		if err != nil {
			log.Println(err.Error())
		}
	}

	log.Println("done")
}
