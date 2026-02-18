package ui

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
	"portfolio_website/internal/views"
)

type Renderer struct {
	bundle *i18n.Bundle
}

func NewRenderer(bundle *i18n.Bundle) Renderer {
	return Renderer{bundle: bundle}
}
func (r Renderer) RenderHome(w http.ResponseWriter, lang string) error {
	localizer := i18n.NewLocalizer(r.bundle, lang, "en")
	hero := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "hero"})
	heroSub := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "hero.sub"})
	profileTitle := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.title"})
	profileBody0 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.body.0"})
	profileBody1 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.bbody.1"})
	profileBody2 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.body.2"})
	profileBody3 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.bbody.3"})
	profileBody4 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.body.4"})
	profileBody5 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.bbody.5"})
	profileBody6 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.body.6"})
	profileBody7 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.bbody.7"})
	profileBody8 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.body.8"})
	profileBody9 := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "profile.bbody.9"})
	data := map[string]interface{}{
		"Hero":         hero,
		"HeroSub":      heroSub,
		"ProfileTitle": profileTitle,
		"ProfileBody0": profileBody0,
		"ProfileBody1": profileBody1,
		"ProfileBody2": profileBody2,
		"ProfileBody3": profileBody3,
		"ProfileBody4": profileBody4,
		"ProfileBody5": profileBody5,
		"ProfileBody6": profileBody6,
		"ProfileBody7": profileBody7,
		"ProfileBody8": profileBody8,
		"ProfileBody9": profileBody9,
	}
	return views.Page.Execute(w, data)
}
