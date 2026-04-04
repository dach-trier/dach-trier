package i18n

import "fmt"
import "golang.org/x/text/language"

type Localizer struct {
	lang   language.Tag
	bundle *Bundle
}

func NewLocalizer(bundle *Bundle, lang language.Tag) *Localizer {
	return &Localizer{bundle: bundle, lang: lang}
}

func (l *Localizer) Localize(id string, a ...any) (string, error) {
	translations, ok := l.bundle.translations[l.lang]
	if !ok {
		return "", ErrNoTranslations{Lang: l.lang}
	}

	translation, ok := translations[id]
	if !ok {
		return "", ErrTranslationMissing{Lang: l.lang, Id: id}
	}

	return fmt.Sprintf(translation, a...), nil
}

func (l *Localizer) MustLocalize(id string, a ...any) string {
	res, err := l.Localize(id, a...)
	if err != nil {
		panic(err)
	}
	return res
}
