package main

import (
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
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

	// Канал для обработки сигналов завершения программы
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Чтение сообщений
	done := make(chan struct{})
	go func() {
		defer close(done)
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
			fmt.Println(len(update))
			for _, el := range update {
				if el.Ticker == "BTCUSDT" || el.Ticker == "ETHUSDT" || el.Ticker == "EOSUSDT" {
					fmt.Println(el.Ticker, " : ", el.PriceChange)
				}
			}
		}
	}()

	// Обработка завершения программы
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Получен сигнал завершения, закрытие соединения...")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Ошибка при отправке сообщения закрытия:", err)
				return
			}
			return
		}
	}
}
