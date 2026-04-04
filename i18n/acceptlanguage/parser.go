package acceptlanguage

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"golang.org/x/text/language"
)

func Select(acceptlanguage string, supportedLanguages []language.Tag) (language.Tag, error) {
	var err error

	tokens := strings.Split(acceptlanguage, ",")
	preferredLang := language.Und
	preferredLangQuality := math.SmallestNonzeroFloat64

	for _, token := range tokens {
		var lang language.Tag
		var quality float64 = 1.0

		token = strings.TrimSpace(token)
		parts := strings.SplitN(token, ";", 2)
		if len(parts) <= 0 {
			panic("unreachable")
		}
		lhs := parts[0]
		lhs = strings.ToLower(lhs)
		lhs = strings.Replace(lhs, "_", "-", 0)

		lang, err = language.Parse(lhs)
		if err != nil {
			return language.Und, err
		}

		supported := false
		for _, supportedLanguage := range supportedLanguages {
			if supported = supportedLanguage == lang; supported {
				break
			}
		}
		if !supported {
			continue
		}

		if len(parts) == 2 {
			rhs := parts[1]
			var ok bool
			if quality, ok = parseQualityValue(rhs); !ok {
				return language.Und, fmt.Errorf("invalid quality value (%s)\n", rhs)
			}
		}

		if quality > preferredLangQuality {
			preferredLang = lang
			preferredLangQuality = quality
		}
	}

	return preferredLang, nil
}

func MustSelect(acceptlanguage string, supportedLanguages []language.Tag) language.Tag {
	lang, err := Select(acceptlanguage, supportedLanguages)
	if err != nil {
		if len(supportedLanguages) == 0 {
			return language.Und
		}
		return supportedLanguages[0]
	}
	return lang
}

func parseQualityValue(raw string) (float64, bool) {
	if !strings.HasPrefix(raw, "q=") {
		return 0, false
	}

	quality := strings.SplitN(raw, "=", 2)[1]

	if len(quality) == 0 {
		return 0, false
	}

	result, err := strconv.ParseFloat(quality, 64)

	if err != nil {
		return 0, false
	}

	return result, true
}
