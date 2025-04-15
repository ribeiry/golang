package service

import (
	"WeatherProject/config"
	"WeatherProject/models"
	db "WeatherProject/repository"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func CreateWeatherEntry(locale, msg string) error {
	w := models.WeatherResponse{
		Country: locale,
		Date:    time.Now().Format("2006-01-02"),
		Text:    msg,
	}
	return db.InsertDb(w)
}

func GetTodayWeather() (*models.WeatherResponse, error) {
	today := time.Now().Format("2006-01-02")
	weather, err := db.FindByDate(today)

	if err != nil || weather == nil || weather.Country == "" || weather.Text == "" {
		log.Println("Error ao buscar ao banco")

		weather, err = GetTemperature()

		if err != nil {
			log.Println("Error ao buscar na API")
		}

		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			saveErrChan := make(chan error)

			go func() {
				saveErrChan <- CreateWeatherEntry(weather.Country, weather.Text)

			}()

			select {
			case err := <-saveErrChan:
				if err != nil {
					log.Println("Error ao salvar no banco: ", err)
				} else {
					log.Println("Salvamento feito com sucesso")
				}
			case <-ctx.Done():
				log.Println(" timeout para esperar o Banco de dados")
			}
		}()

		return weather, nil
	} else {
		return weather, nil
	}
}

func GetAllWeather() ([]models.WeatherResponse, error) {
	return db.FindAllDb()
}

func GetTemperature() (*models.WeatherResponse, error) {

	url := fmt.Sprintf("%s%s?token=%s", config.BaseURL, config.SynopticPath, config.APIKey)
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, err

	}
	var weather []models.WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		log.Println(err)
		return nil, err

	}
	if len(weather) == 0 {
		log.Println(err)

		return nil, err
	}
	return &weather[0], nil
}
