package slug

import (
	"bytes"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/mahbubzulkarnain/gn/internal/pkg/unidecode"
)

var (
	CustomSub     map[string]string
	CustomRuneSub map[rune]string

	MaxLength int

	EnableSmartTruncate = true

	Lowercase = true

	AppendTimestamp = false

	regexpNonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-_]")
	regexpMultipleDashes     = regexp.MustCompile("-+")
)

func Make(s string) (slug string) {
	return MakeLang(s, "en")
}

func MakeLang(s string, lang string) (slug string) {
	slug = strings.TrimSpace(s)

	slug = SubstituteRune(slug, CustomRuneSub)
	slug = Substitute(slug, CustomSub)

	switch strings.ToLower(lang) {
	case "bg", "bgr":
		slug = SubstituteRune(slug, bgSub)
	case "cs", "ces":
		slug = SubstituteRune(slug, csSub)
	case "de", "deu":
		slug = SubstituteRune(slug, deSub)
	case "en", "eng":
		slug = SubstituteRune(slug, enSub)
	case "es", "spa":
		slug = SubstituteRune(slug, esSub)
	case "fi", "fin":
		slug = SubstituteRune(slug, fiSub)
	case "fr", "fra":
		slug = SubstituteRune(slug, frSub)
	case "gr", "el", "ell":
		slug = SubstituteRune(slug, grSub)
	case "hu", "hun":
		slug = SubstituteRune(slug, huSub)
	case "id", "idn", "ind":
		slug = SubstituteRune(slug, idSub)
	case "it", "ita":
		slug = SubstituteRune(slug, itSub)
	case "kz", "kk", "kaz":
		slug = SubstituteRune(slug, kkSub)
	case "nb", "nob":
		slug = SubstituteRune(slug, nbSub)
	case "nl", "nld":
		slug = SubstituteRune(slug, nlSub)
	case "nn", "nno":
		slug = SubstituteRune(slug, nnSub)
	case "pl", "pol":
		slug = SubstituteRune(slug, plSub)
	case "pt", "prt", "pt-br", "br", "bra", "por":
		slug = SubstituteRune(slug, ptSub)
	case "ro", "rou":
		slug = SubstituteRune(slug, roSub)
	case "sl", "slv":
		slug = SubstituteRune(slug, slSub)
	case "sv", "swe":
		slug = SubstituteRune(slug, svSub)
	case "tr", "tur":
		slug = SubstituteRune(slug, trSub)
	default:
		slug = SubstituteRune(slug, enSub)
	}

	slug = unidecode.Unidecode(slug)

	if Lowercase {
		slug = strings.ToLower(slug)
	}

	if !EnableSmartTruncate && len(slug) >= MaxLength {
		slug = slug[:MaxLength]
	}

	// Process all remaining symbols
	slug = regexpNonAuthorizedChars.ReplaceAllString(slug, "-")
	slug = regexpMultipleDashes.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-_")

	if MaxLength > 0 && EnableSmartTruncate {
		slug = smartTruncate(slug)
	}

	if AppendTimestamp {
		slug = slug + "-" + timestamp()
	}

	return slug
}

func Substitute(s string, sub map[string]string) (buf string) {
	buf = s
	var keys []string
	for k := range sub {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		buf = strings.Replace(buf, key, sub[key], -1)
	}
	return
}

func SubstituteRune(s string, sub map[rune]string) string {
	var buf bytes.Buffer
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf.WriteString(d)
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func smartTruncate(text string) string {
	if len(text) <= MaxLength {
		return text
	}

	for i := MaxLength; i >= 0; i-- {
		if text[i] == '-' {
			return text[:i]
		}
	}
	return text[:MaxLength]
}

func timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func IsSlug(text string) bool {
	if text == "" ||
		(MaxLength > 0 && len(text) > MaxLength) ||
		text[0] == '-' || text[0] == '_' ||
		text[len(text)-1] == '-' || text[len(text)-1] == '_' {
		return false
	}
	for _, c := range text {
		if (c < 'a' || c > 'z') && c != '-' && c != '_' && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}
