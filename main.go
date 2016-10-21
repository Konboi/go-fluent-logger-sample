//go:generate msgp
//msgp:ignore Logger

package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/mattn/go-gimei"
	"github.com/serenize/snaker"
)

type Player struct {
	ID   int    `msg:"id"`
	Name string `msg:"name"`
}

type PlayerComment struct {
	ID        int    `msg:"id"`
	Comment   string `msg:"comment"`
	PlayerID  int    `msg:"player_id"`
	CreatedAt int64  `msg:"created_at"`
	UpdatedAt int64  `msg:"updated_at"`
}

type Logger struct {
	Fluent *fluent.Fluent
}

func main() {
	flogger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "127.0.0.1"})

	if err != nil {
		log.Fatalln(err.Error())
	}
	defer flogger.Close()

	tag := "debug.player"
	logger := Logger{
		Fluent: flogger,
	}

	for i := 1; i < 10; i++ {
		p := Player{
			ID:   i,
			Name: gimei.NewName().Kanji(),
		}
		pc := PlayerComment{
			ID:        i + 10,
			Comment:   "dummmy",
			PlayerID:  p.ID,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		err := logger.Post(tag, &p, &pc)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("done")
}

func (logger Logger) Post(tag string, vs ...interface{}) error {
	postData := make(map[string]interface{}, len(vs))

	for _, v := range vs {
		rt := reflect.TypeOf(v)

		if rv := reflect.ValueOf(v); rv.Kind() != reflect.Ptr {
			return fmt.Errorf("v is not pointer")
		}

		postData[snaker.CamelToSnake(fmt.Sprintf("%s", rt.Elem().Name()))] = v
	}
	err := logger.Fluent.Post(tag, postData)
	if err != nil {
		log.Println(tag, "post data  error", err)
		return err
	}

	return nil
}
