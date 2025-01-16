package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"log"
)

type MarkPriceUpdate struct {
	Ticker      string `json:"s"` // Символ (ticker)
	PriceChange string `json:"p"`
}

func main() {
	// URL для подключения к WebSocket Binance (замените на нужный поток)
	socketURL := "wss://fstream.binance.com/ws/!markPrice@arr"

	// Установка соединения
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, nil)
	if err != nil {
		log.Fatalf("Ошибка подключения к WebSocket: %v", err)
	}
	defer conn.Close()

	myData := []string{"BTCUSDT", "ETHUSDT", "EOSUSDT"}

	// Чтение сообщений
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Ошибка чтения сообщения:", err)
				return
			}
			if len(message) == 0 {
				log.Println("Пустое сообщение, пропускаем...")
				continue
			}
			//Распарсить сообщение в структуру
			var update []MarkPriceUpdate
			if err := json.Unmarshal(message, &update); err != nil {
				log.Println("Ошибка парсинга JSON:", err)
				continue
			}
			i := 0
			for _, el := range update {
				if i < len(myData) && el.Ticker == myData[i] {
					fmt.Printf("Тикер: %s, Цена: %s\n", el.Ticker, el.PriceChange)
					i++
				}
			}
		}
	}()

	select {}

	// Обработка завершения программы
}
