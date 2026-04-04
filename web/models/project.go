package models

import (
	"html/template"

	"golang.org/x/text/language"
)

type Project struct {
	Preview     *Image
	Name        template.HTML
	Description template.HTML
}

func GetProjects(lang language.Tag) []Project {
	switch lang {
	case language.English:
		return []Project{
			{
				Preview:     &Image{URL: "assets/bereginja-b3c3psg.jpg"},
				Name:        template.HTML("First Ukrainian School in Trier &mdash; Bereginja"),
				Description: template.HTML("An educational and cultural center for Ukrainian children and youth in Germany aged 5 to 17. <br /><br /> Currently, <strong>36 children</strong> study in 2 preparatory groups and 6 classes. <strong>8 experienced volunteer teachers</strong> with pedagogical education teach the children and teenagers."),
			},
			{
				Preview:     &Image{URL: "assets/camp-vpqiwoj.jpg"},
				Name:        template.HTML("Youth Camp"),
				Description: template.HTML("We invite children aged 9–16 to spend unforgettable holidays at our camp! <br /><br /> A place filled with friendship, warmth, and love for native culture, designed for relaxation, personal growth, and new connections."),
			},
			{
				Preview:     &Image{URL: "assets/art-ryeagqz.jpg"},
				Name:        template.HTML("Art Workshop"),
				Description: template.HTML("A safe space where Ukrainian youth in Germany can explore their emotions, find inner peace, and reconnect with themselves through creativity."),
			},
			{
				Preview:     &Image{URL: "assets/band-60xx7c3.jpg"},
				Name:        template.HTML("Music Band"),
				Description: template.HTML("Founded in 2024, an ongoing project supporting young talents."),
			},
		}

	case language.German:
		return []Project{
			{
				Preview:     &Image{URL: "assets/bereginja-b3c3psg.jpg"},
				Name:        template.HTML("Erste Ukrainische Schule in Trier &mdash; Bereginja"),
				Description: template.HTML("Ein Bildungs- und Kulturzentrum für ukrainische Kinder und Jugendliche in Deutschland im Alter von 5 bis 17 Jahren. <br /><br /> Derzeit besuchen <strong>36 Kinder</strong> 2 Vorbereitungsgruppen und 6 Klassen. <strong>8 erfahrene ehrenamtliche Lehrkräfte</strong> mit pädagogischer Ausbildung unterrichten die Kinder und Jugendlichen."),
			},
			{
				Preview:     &Image{URL: "assets/camp-vpqiwoj.jpg"},
				Name:        template.HTML("Kindercamp"),
				Description: template.HTML("Wir laden Kinder im Alter von 9–16 Jahren ein, unvergessliche Ferien in unserem Lager zu verbringen! <br /><br /> Ein Ort voller Freundschaft, Wärme und Liebe zur eigenen Kultur, geschaffen für Erholung, persönliche Entwicklung und neue Bekanntschaften."),
			},
			{
				Preview:     &Image{URL: "assets/art-ryeagqz.jpg"},
				Name:        template.HTML("Kunstwerkstatt"),
				Description: template.HTML("Ein sicherer Raum, in dem ukrainische Jugendliche in Deutschland ihre Emotionen erkunden, innere Ruhe finden und über Kreativität wieder zu sich selbst finden können."),
			},
			{
				Preview:     &Image{URL: "assets/band-60xx7c3.jpg"},
				Name:        template.HTML("Musikband"),
				Description: template.HTML("Gegründet 2024, ein laufendes Projekt zur Unterstützung junger Talente."),
			},
		}

	case language.Ukrainian:
		return []Project{
			{
				Preview:     &Image{URL: "assets/bereginja-b3c3psg.jpg"},
				Name:        template.HTML("Перша українська школа у Трірі &mdash; Берегиня"),
				Description: template.HTML("Освітньо-культурний центр для українських дітей та молоді в Німеччині віком від 5 до 17 років. <br /> <br /> Наразі у школі навчаються <strong>36 дітей</strong> у 2 підготовчих групах та 6 класаx. <strong>8 досвідчених волонтерів-викладачів</strong> із педагогічною освітою навчають дітей та підлітків."),
			},
			{
				Preview:     &Image{URL: "assets/camp-vpqiwoj.jpg"},
				Name:        template.HTML("Дитячі Табори"),
				Description: template.HTML("Запрошуємо дітей віком 9–16 років провести незабутні канікули в нашому таборі! <br /> <br /> Це місце, де панує дружба, тепло та любов до рідної культури, створене для відпочинку, розвитку та нових знайомств."),
			},
			{
				Preview:     &Image{URL: "assets/art-ryeagqz.jpg"},
				Name:        template.HTML("Арт Майстерня"),
				Description: template.HTML("безпечний простір, де українська молодь у Німеччині може досліджувати свої емоції, знаходити внутрішній спокій та відновлювати зв&#39;язок з собою через творчість."),
			},
			{
				Preview:     &Image{URL: "assets/band-60xx7c3.jpg"},
				Name:        template.HTML("Музичний Гурт"),
				Description: template.HTML("Заснований у 2024 році, триваючий проєкт для підтримки молодих талантів"),
			},
		}

	default:
		panic("unreachable")
	}
}
