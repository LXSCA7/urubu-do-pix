package routine

import (
	"log"
	"time"

	// "math/rand"
	"os"
	"urubu-do-pix/services"

	"math/rand"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()
	randomNumber := randomize()
	percentage := (randomNumber - 1.0) * 100
	//                   * * * * *
	_, err := c.AddFunc("0 0 * * *", func() {
		log.Printf("Executando rendimento diario de %f%%", percentage)
		err := services.DailyInvestment(randomNumber)
		if err != nil {
			log.Printf("Erro ao renderizar juros: %v\n", err)
		} else {
			log.Printf("Juros renderizados com sucesso\n")
		}
	})
	if err != nil {
		log.Printf("erro ao agendar tarefa: %v\n", err)
		os.Exit(1)
	}

	c.Start()
}

func randomize() float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	min := 1.04
	max := 1.08

	return min + rand.Float64()*(max-min)
}
