package web

import (
	"embed"
	"golang.org/x/text/language"
	"log"
	"net/http"

	"html/template"

	i18n "dach-trier.com/i18n"
	acceptlanguage "dach-trier.com/i18n/acceptlanguage"
	models "dach-trier.com/web/models"

	chi "github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

//go:embed i18n/*.json
var i18nFS embed.FS
var i18nSupportedLanguages = []language.Tag{language.AmericanEnglish, language.BritishEnglish, language.English, language.German, language.Ukrainian, language.Russian}

// Global Application State
type App struct {
	bundle *i18n.Bundle
}

func NewApp() (*App, error) {
	app := &App{}

	app.bundle = i18n.NewBundle()
	app.bundle.MustLoadJsonTranslationsFS(language.English, i18nFS, "i18n/en.json")
	app.bundle.MustLoadJsonTranslationsFS(language.German, i18nFS, "i18n/de.json")
	app.bundle.MustLoadJsonTranslationsFS(language.Ukrainian, i18nFS, "i18n/uk.json")

	return app, nil
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

	lang := acceptlanguage.MustSelect(r.Header.Get("Accept-Language"), i18nSupportedLanguages)
	switch lang {
	case language.AmericanEnglish, language.BritishEnglish:
		lang = language.English
	case language.Russian:
		lang = language.Ukrainian
	}
	localizer := i18n.NewLocalizer(app.bundle, lang)

	tmpl = template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{
		"t": func(id string, a ...any) template.HTML {
			return template.HTML(localizer.MustLocalize(id, a...))
		},
	})

	tmpl = template.Must(tmpl.ParseFiles("web/templates/index.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/header.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/footer.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/templates/icons/*.html"))

	type props struct {
		Title  string
		Events []models.Event
	}

	tmpl.ExecuteTemplate(w, "page", props{
		Title:  "DACH e.V. Trier",
		Events: models.GetEvents(lang),
	})
}

func (app *App) serveProjectPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	lang := acceptlanguage.MustSelect(r.Header.Get("Accept-Language"), i18nSupportedLanguages)
	switch lang {
	case language.AmericanEnglish, language.BritishEnglish:
		lang = language.English
	case language.Russian:
		lang = language.Ukrainian
	}
	localizer := i18n.NewLocalizer(app.bundle, lang)

	tmpl = template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{
		"t": func(id string, a ...any) template.HTML {
			return template.HTML(localizer.MustLocalize(id, a...))
		},
	})

	tmpl = template.Must(tmpl.ParseFiles("web/templates/projects.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/header.html"))
	tmpl = template.Must(tmpl.ParseFiles("web/templates/footer.html"))
	tmpl = template.Must(tmpl.ParseGlob("web/templates/icons/*.html"))

	type props struct {
		Title    string
		Projects []models.Project
	}

	tmpl.ExecuteTemplate(w, "page", props{
		Title:    "DACH e.V. Trier",
		Projects: models.GetProjects(lang),
	})
}

func (app *App) serveEventPage(w http.ResponseWriter, r *http.Request) {
	var tmpl *template.Template

	lang := acceptlanguage.MustSelect(r.Header.Get("Accept-Language"), i18nSupportedLanguages)
	switch lang {
	case language.AmericanEnglish, language.BritishEnglish:
		lang = language.English
	case language.Russian:
		lang = language.Ukrainian
	}
	localizer := i18n.NewLocalizer(app.bundle, lang)

	tmpl = template.New("")
	tmpl = tmpl.Funcs(template.FuncMap{
		"t": func(id string, a ...any) template.HTML {
			return template.HTML(localizer.MustLocalize(id, a...))
		},
	})

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
