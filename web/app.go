package web

import (
	"log"
	"net/http"

	"html/template"

	models "dach-trier.com/web/models"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

// Global Application State
type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) Router() http.Handler {
	router := chi.NewRouter()
	router.Use(chi_middleware.Logger)

	// assets

	fs := http.FileServer(http.Dir("web/assets/"))
	router.Get("/assets/*", http.StripPrefix("/assets/", fs).ServeHTTP)

	// pages

	router.Get("/", app.serveIndexPage)
	router.Get("/projects", app.serveProjectPage)
	router.Get("/events", app.serveEventPage)
	router.NotFound(app.serveNotFoundPage)

	return router
}

func (app *App) serveIndexPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseFiles("web/templates/index.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/header.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/footer.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/templates/icons/*.html"))

	type props struct {
		Title  string
		Events []models.Event
	}

	tmpl.ExecuteTemplate(w, "page", props{
		Title: "DACH e.V. Trier",
		Events: []models.Event{
			models.Event{Preview: &models.Image{URL: "assets/walzer-y0ldi4q.jpg"}, Name: "Walzer"},
			models.Event{Preview: &models.Image{URL: "assets/art-1wl3o4x.jpg"}, Name: "Art Workshop"},
			models.Event{Preview: &models.Image{URL: "assets/theater-qssoth3.jpg"}, Name: "Theater"},
			models.Event{Preview: &models.Image{URL: "assets/jugendaustausch-sqji6l4.jpg"}, Name: "Jugendaustausch"},
			models.Event{Preview: &models.Image{URL: "assets/band-60xx7c3.jpg"}, Name: "Band"},
			models.Event{Preview: &models.Image{URL: "assets/camp-6bqndr3.jpg"}, Name: "Camp"},
			models.Event{Preview: &models.Image{URL: "assets/bereginja-3avd0zj.jpg"}, Name: "Bereginja"},
		},
	})
}

func (app *App) serveProjectPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseFiles("web/templates/projects.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/header.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/footer.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/templates/icons/*.html"))

	type props struct {
		Title    string
		Projects []models.Project
	}

	tmpl.ExecuteTemplate(w, "page", props{
		Title: "DACH e.V. Trier",
		Projects: []models.Project{
			models.Project{
				Preview:     &models.Image{URL: "assets/bereginja-b3c3psg.jpg"},
				Name:        template.HTML("Перша українська школа у Трірі &mdash; Берегиня"),
				Description: template.HTML("Освітньо-культурний центр для українських дітей та молоді в Німеччині віком від 5 до 17 років. <br /> <br /> Наразі у школі навчаються <strong>36 дітей</strong> у 2 підготовчих групах та 6 класаx. <strong>8 досвідчених волонтерів-викладачів</strong> із педагогічною освітою навчають дітей та підлітків."),
			},
			models.Project{
				Preview:     &models.Image{URL: "assets/camp-vpqiwoj.jpg"},
				Name:        template.HTML("Дитячі Табори"),
				Description: template.HTML("Запрошуємо дітей віком 9–16 років провести незабутні канікули в нашому таборі! <br /> <br /> Це місце, де панує дружба, тепло та любов до рідної культури, створене для відпочинку, розвитку та нових знайомств."),
			},
			models.Project{
				Preview:     &models.Image{URL: "assets/art-ryeagqz.jpg"},
				Name:        template.HTML("Арт Майстерня"),
				Description: template.HTML("безпечний простір, де українська молодь у Німеччині може досліджувати свої емоції, знаходити внутрішній спокій та відновлювати зв&#39;язок з собою через творчість."),
			},
			models.Project{
				Preview:     &models.Image{URL: "assets/band-60xx7c3.jpg"},
				Name:        template.HTML("Музичний Гурт"),
				Description: template.HTML("Заснований у 2024 році, триваючий проєкт для підтримки молодих талантів"),
			},
		},
	})
}

func (app *App) serveEventPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	tmpl = template.New("")
	tmpl = template.Must(tmpl.ParseFiles("web/templates/event-page.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/header.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/footer.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/templates/icons/*.html"))

	type props struct {
		Title string
	}

	tmpl.ExecuteTemplate(w, "page", props{
		Title: "DACH e.V. Trier",
	})
}

func (app *App) serveNotFoundPage(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("404 not found"))

	if err != nil {
		log.Fatalf("failed to write 404 response: %v", err)
	}
}
