// internal/i18n/i18n.go
package i18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Locale string

const (
	LocaleES Locale = "ar"
	LocaleEN Locale = "en"
)

type Table map[string]string
type Catalog map[Locale]Table

func Load(dir string, fallbacks []Locale) (Catalog, error) {
	cat := Catalog{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		loc := Locale(e.Name()[:len(e.Name())-5]) // “ar.json”
		b, err := os.ReadFile(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		var tbl Table
		if err := json.Unmarshal(b, &tbl); err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		cat[loc] = tbl
	}
	return cat.ensureFallbacks(fallbacks), nil
}
func (c Catalog) ensureFallbacks(order []Locale) Catalog {
	for loc, tbl := range c {
		c[loc] = tbl // shallow copy optional
	}
	c["_order"] = Table{}
	for i, loc := range order {
		c["_order"][fmt.Sprint(i)] = string(loc)
	}
	return c
}

type Translator struct {
	primary  Locale
	catalog  Catalog
	fallback []Locale
}

func NewTranslator(cat Catalog, primary Locale, fallback []Locale) Translator {
	return Translator{primary, cat, append([]Locale{primary}, fallback...)}
}
func (t Translator) T(key string) string {
	for _, loc := range t.fallback {
		if val, ok := t.catalog[loc][key]; ok && val != "" {
			return val
		}
	}
	return key // last-resort debug
}
