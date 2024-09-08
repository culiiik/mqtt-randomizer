package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Funkce pro generování náhodného čísla
func randomMetric() float64 {
	return rand.Float64() * 1000
}

func main() {
	// MQTT nastavení
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://mqttexporter:1883") // adresa MQTT brokera
	opts.SetClientID("go_mqtt_client")

	// Přidání autorizace (uživatelské jméno a heslo)
	opts.SetUsername("culik")
	opts.SetPassword("Pavl94ek2680_")

	// Vytvoření MQTT klienta
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Printf("Error connecting to MQTT broker: %v\n", token.Error())
		os.Exit(1)
	}
	fmt.Println("Connected to MQTT broker")

	// Periodické odesílání metrik každých 10 sekund
	for {
		// Vygenerování náhodných hodnot pro metriky
		temp := randomMetric()
		humidity := randomMetric()
		pressure := randomMetric()

		// Publikování metrik do MQTT brokera
		client.Publish("metrics/temp", 0, false, fmt.Sprintf("%.2f", temp))
		client.Publish("metrics/humidity", 0, false, fmt.Sprintf("%.2f", humidity))
		client.Publish("metrics/pressure", 0, false, fmt.Sprintf("%.2f", pressure))

		fmt.Printf("Published: temp=%.2f, humidity=%.2f, pressure=%.2f\n", temp, humidity, pressure)

		// Čekání 10 sekund
		time.Sleep(10 * time.Second)
	}
}
