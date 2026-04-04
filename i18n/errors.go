package i18n

import "fmt"
import "golang.org/x/text/language"

type ErrNoTranslations struct {
	Lang language.Tag
}

func (e ErrNoTranslations) Error() string {
	return fmt.Sprintf("no translations found for %q", e.Lang)
}

type ErrTranslationMissing struct {
	Lang language.Tag
	Id   string
}

func (e ErrTranslationMissing) Error() string {
	return fmt.Sprintf("no translation found with id %q in language %q", e.Id, e.Lang)
}

type ErrTranslationDuplicate struct {
	Lang language.Tag
	Id   string
}

func (e ErrTranslationDuplicate) Error() string {
	return fmt.Sprintf("duplicate translation with id %q in language %q", e.Id, e.Lang)
}
