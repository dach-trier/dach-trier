package models

import "golang.org/x/text/language"

type Event struct {
	Preview *Image
	Name    string
}

func GetEvents(lang language.Tag) []Event {
	switch lang {
	case language.English:
		return []Event{
			{Preview: &Image{URL: "assets/walzer-y0ldi4q.jpg"}, Name: "Waltz"},
			{Preview: &Image{URL: "assets/art-1wl3o4x.jpg"}, Name: "Art Workshop"},
			{Preview: &Image{URL: "assets/theater-qssoth3.jpg"}, Name: "Theater"},
			{Preview: &Image{URL: "assets/jugendaustausch-sqji6l4.jpg"}, Name: "Youth Exchange"},
			{Preview: &Image{URL: "assets/band-60xx7c3.jpg"}, Name: "Music Band"},
			{Preview: &Image{URL: "assets/camp-6bqndr3.jpg"}, Name: "Youth Camp"},
			{Preview: &Image{URL: "assets/bereginja-3avd0zj.jpg"}, Name: "Ukrainian School in Trier - Bereginja"},
		}
	case language.German:
		return []Event{
			{Preview: &Image{URL: "assets/walzer-y0ldi4q.jpg"}, Name: "Walzer"},
			{Preview: &Image{URL: "assets/art-1wl3o4x.jpg"}, Name: "Kunstworkshop"},
			{Preview: &Image{URL: "assets/theater-qssoth3.jpg"}, Name: "Theater"},
			{Preview: &Image{URL: "assets/jugendaustausch-sqji6l4.jpg"}, Name: "Jugendaustausch"},
			{Preview: &Image{URL: "assets/band-60xx7c3.jpg"}, Name: "Musikband"},
			{Preview: &Image{URL: "assets/camp-6bqndr3.jpg"}, Name: "Jugendcamp"},
			{Preview: &Image{URL: "assets/bereginja-3avd0zj.jpg"}, Name: "Ukrainische Schule in Trier - Bereginja"},
		}
	case language.Ukrainian:
		return []Event{
			{Preview: &Image{URL: "assets/walzer-y0ldi4q.jpg"}, Name: "Вальс"},
			{Preview: &Image{URL: "assets/art-1wl3o4x.jpg"}, Name: "Арт Майстерня"},
			{Preview: &Image{URL: "assets/theater-qssoth3.jpg"}, Name: "Театр"},
			{Preview: &Image{URL: "assets/jugendaustausch-sqji6l4.jpg"}, Name: "Молодіжний обмін"},
			{Preview: &Image{URL: "assets/band-60xx7c3.jpg"}, Name: "Музичний гурт"},
			{Preview: &Image{URL: "assets/camp-6bqndr3.jpg"}, Name: "Молодіжний табір"},
			{Preview: &Image{URL: "assets/bereginja-3avd0zj.jpg"}, Name: "Українська школа в Трірі - Берегиня"},
		}
	default:
		panic("unreachable")
	}
}
