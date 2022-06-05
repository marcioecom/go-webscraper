package crawler

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
	"webscraper/server/models"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/robfig/cron/v3"
	"gorm.io/datatypes"
)

func SchedulerJob() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { scrapJobs() })
	c.Start()
}

func formatJobTimes(timeLeftStr, publishedAtStr *string) (a, b int64) {
	timeLeft, _ := strconv.ParseInt(*timeLeftStr, 10, 64)
	publishedAt, _ := strconv.ParseInt(*publishedAtStr, 10, 64)

	return timeLeft, publishedAt
}

func scrapJobs() {
	l := launcher.MustNewManaged("ws://crawler:7317")
	page := rod.New().Client(l.MustClient()).MustConnect().MustPage("https://www.99freelas.com.br/projects?order=mais-recentes&categoria=web-mobile-e-software")

	defer page.MustClose()

	page.MustWaitLoad()

	els, err := page.Elements("li.result-item")
	if err != nil {
		log.Fatal(err.Error())
	}

	sliceOfJobs := []models.Job{}

	for _, el := range els {
		tags := []string{}

		skillsEl := el.MustElements("a.habilidade")
		for _, v := range skillsEl {
			tags = append(tags, v.MustText())
		}
		tagsJson, _ := json.Marshal(tags)

		timeLeftStr := el.MustElement("b.datetime").MustAttribute("cp-datetime")
		publishedAtStr := el.MustElement("b.datetime-restante").MustAttribute("cp-datetime")
		timeLeft, publishedAt := formatJobTimes(timeLeftStr, publishedAtStr)

		information, _ := el.MustElement("p.information").Text()
		offersResult := strings.Split(information, "Propostas: ")
		offersResult = strings.Split(offersResult[1], " | Interessados: ")

		offers := offersResult[0]
		interested := offersResult[1]

		job := models.Job{
			Title:       el.MustElement("h1.title").MustText(),
			Description: el.MustElement("div.description").MustText(),
			Tags:        datatypes.JSON(tagsJson),
			Offers:      offers,
			Interested:  interested,
			SeenAt:      time.Now(),
			TimeLeft:    timeLeft,
			PublishedAt: publishedAt,
		}

		sliceOfJobs = append(sliceOfJobs, job)
	}

	for _, v := range sliceOfJobs {
		models.UpdateJob(v)
	}
}
