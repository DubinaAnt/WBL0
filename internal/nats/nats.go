package nats

import (
	"WBL0/internal/cache"

	"WBL0/internal/model"
	"WBL0/internal/repository"
	"encoding/json"
	"log"

	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
)

func NewConnection() (stan.Conn, error) {
	natsConnection, err := stan.Connect(viper.GetString("nats.stanclusterid"), viper.GetString("nats.clientid"))
	if err != nil {
		return nil, err
	}
	return natsConnection, err
}

func NewNatsSubscriber(conn stan.Conn, repos *repository.Repository, cache *cache.Cache) stan.Subscription {

	handler := func(msg *stan.Msg) {
		var order model.Order

		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("unmarshal error: %s", err.Error())
			err = msg.Ack()
			if err != nil {
				log.Printf("nats handler: %s", err.Error())
			}
			return
		}

		_, err = repos.CreateOrder(order)
		if err != nil {
			log.Printf("create order db error: %s", err.Error())
			return
		}
		cache.Set(order.UID, order)
		err = msg.Ack()
		if err != nil {
			log.Printf("nats handler ask: %s", err.Error())
		}
	}

	sub, err := conn.Subscribe(
		"order",
		handler,
		stan.DurableName("stand1"),
		stan.SetManualAckMode(),
	)
	if err != nil {
		log.Fatalf("create subscribe error: %s", err.Error())
	}

	return sub
}
