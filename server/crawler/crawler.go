package crawler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	c.AddFunc("0 */1 * * *", func() { scrapJobs() })
	c.Start()
}

func formatJobTimes(timeLeftStr, publishedAtStr *string) (a, b int64) {
	timeLeft, _ := strconv.ParseInt(*timeLeftStr, 10, 64)
	publishedAt, _ := strconv.ParseInt(*publishedAtStr, 10, 64)

	return timeLeft, publishedAt
}

func handleError(err error) {
	var evalErr *rod.ErrEval
	if errors.Is(err, context.DeadlineExceeded) { // timeout error
		fmt.Println("timeout err")
	} else if errors.As(err, &evalErr) { // eval error
		fmt.Println(evalErr.LineNumber)
	} else if err != nil {
		fmt.Println("can't handle", err)
	}
}

func scrapJobs() {
	var page *rod.Page
	err := rod.Try(func() {
		l := launcher.MustNewManaged("ws://crawler:7317")
		page = rod.
			New().
			Client(l.MustClient()).
			CancelTimeout().
			MustConnect().
			MustPage("https://www.99freelas.com.br/projects?order=mais-recentes&categoria=web-mobile-e-software")
	})
	handleError(err)

	defer page.MustClose()

	page.MustWaitLoad()

	els, err := page.Elements("li.result-item")
	if err != nil {
		log.Fatal(err.Error())
	}

	sliceOfJobs := []models.Job{}

	for _, el := range els {
		tags := []string{}

		var skillsEl rod.Elements
		err = rod.Try(func() {
			skillsEl = el.MustElements("a.habilidade")
		})
		handleError(err)

		for _, v := range skillsEl {
			tags = append(tags, v.MustText())
		}
		tagsJson, _ := json.Marshal(tags)

		var timeLeftStr *string
		var publishedAtStr *string
		err = rod.Try(func() {
			timeLeftStr = el.MustElement("b.datetime").MustAttribute("cp-datetime")
			publishedAtStr = el.MustElement("b.datetime-restante").MustAttribute("cp-datetime")
		})
		handleError(err)

		timeLeft, publishedAt := formatJobTimes(timeLeftStr, publishedAtStr)

		var information string
		err = rod.Try(func() {
			information, err = el.MustElement("p.information").Text()
			if err != nil {
				panic(err)
			}
		})
		handleError(err)

		offersResult := strings.Split(information, "Propostas: ")
		offersResult = strings.Split(offersResult[1], " | Interessados: ")

		offers := offersResult[0]
		interested := offersResult[1]

		url := fmt.Sprint("https://www.99freelas.com.br" + *el.MustElement("h1.title a").MustAttribute("href"))

		job := models.Job{
			Title:       el.MustElement("h1.title").MustText(),
			Description: el.MustElement("div.description").MustText(),
			Tags:        datatypes.JSON(tagsJson),
			Url:         url,
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

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%v] Job runned\n", time.Now().In(loc).Format("02-Jan-2006 15:04:05"))
}
