package lang

import (
	"os"
)

var (
	LocalLanguage *localLanguage
)

type localLanguage struct {
	local    string
	language map[string]*Language
}

func Load() {
	LocalLanguage = &localLanguage{}
	LocalLanguage.language = map[string]*Language{}
	local := os.Getenv("SLS_LANG")
	if local != "" {
		LocalLanguage.setLocal(local)
	}
	LocalLanguage.setLanguage(ZHCN, zh)
	LocalLanguage.setLanguage(EN, en)
}

func (i *localLanguage) setLocal(local string) {
	i.local = local
}

func (i *localLanguage) setLanguage(local string, language *Language) {
	i.language[local] = language
}

func (i *localLanguage) Get(k string) string {
	if language, ok := i.language[i.local]; ok {
		return language.Get(k)
	}
	return k
}

type Language map[string]string

func (v Language) Get(key string) string {
	if v == nil {
		return key
	}
	vs := v[key]
	if len(vs) == 0 {
		return key
	}
	return vs
}

func (v Language) Set(key, value string) {
	v[key] = value
}

func (v Language) Has(key string) bool {
	_, ok := v[key]
	return ok
}
