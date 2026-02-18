// internal/web/handlers.go
package handlers

type IndexData struct {
	Translator i18n.Translator
	// other page data...
}

func (d IndexData) T(key string) string {
	return d.Translator.T(key)
}
func IndexHandler(cat i18n.Catalog) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loc := detectLocale(r) // e.g., URL /en/..., Accept-Language, cookie, default ES
		t := i18n.NewTranslator(cat, loc, []i18n.Locale{i18n.LocaleEN})
		data := IndexData{Translator: t}
		templates.Index(data).Render(r.Context(), w)
	}
}
