package localization

import (
	"github.com/eduardolat/goeasyi18n"
)

var _ Localization = (*localizationImpl)(nil)

type Localization interface {
	IsSupported(langName Language) bool
	T(key string, options *TranslateParams) string
	AddLanguages(langs map[Language]string) error
}

type localizationImpl struct {
	driver *goeasyi18n.I18n
}

type TranslateParams struct {
	LangName Language
	Data     any
	Count    *int
	Gender   *string
}

func NewLocalization() Localization {
	driver := goeasyi18n.NewI18n()

	return &localizationImpl{
		driver: driver,
	}
}

func (s *localizationImpl) T(key string, options *TranslateParams) string {
	if options == nil {
		options = &TranslateParams{}
	}

	if !s.IsSupported(options.LangName) {
		options.LangName = EN_US
	}

	return s.driver.T(string(options.LangName), key, goeasyi18n.Options{
		Data:   options.Data,
		Count:  options.Count,
		Gender: options.Gender,
	})
}

func (s *localizationImpl) IsSupported(langName Language) bool {
	return s.driver.HasLanguage(string(langName))
}

func (s *localizationImpl) AddLanguages(langs map[Language]string) error {
	for langName, path := range langs {
		lang, err := goeasyi18n.LoadFromYamlFiles(path)

		if err != nil {
			return err
		}

		s.driver.AddLanguage(string(langName), lang)
	}

	return nil
}
