package routine

import (
	"log"
	"os"
	"urubu-do-pix/services"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("Executanto rendimento di'ario de 7%")
		err := services.DailyInvestment()
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
