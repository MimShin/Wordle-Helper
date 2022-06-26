package wordle

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

type FilteredDict struct {
	Words       []string
	Filter, Key string
}

type Wordle struct {
	history []FilteredDict
}

func NewWordle(path string) *Wordle {

	words, err := os.ReadFile(path)
	if err != nil || len(words) == 0 {
		log.Errorf("unable to read form file: %s", path)
		return nil
	}

	d := FilteredDict{}
	d.Words = strings.Split(string(words), "\n")
	n := len(d.Words) - 1
	if d.Words[n] == "" {
		d.Words = d.Words[:n]
	}

	w := Wordle{
		history: []FilteredDict{}}
	w.history = append(w.history, d)
	return &w
}

func (w Wordle) String() string {
	str := ""
	for _, h := range w.history {
		str += fmt.Sprintf("count: %5d, filter: %s, key: %s\n", len(h.Words), h.Filter, h.Key)
	}
	return str
}

func (w Wordle) Current() FilteredDict {
	if len(w.history) == 0 {
		return FilteredDict{}
	}
	return w.history[len(w.history)-1]
}

func (w *Wordle) Filter(filter, key string) {
	badChars, goodChars := "", ""

	for i := 0; i < len(key); i++ {
		if key[i] == '-' {
			badChars += string(filter[i])
			continue
		}
		if key[i] == '+' {
			goodChars += string(filter[i])
			continue
		}
	}

	for i := 0; i < len(goodChars); i++ {
		badChars = strings.Replace(badChars, string(goodChars[i]), "", 1)
	}

	d := w.Current()
	f := []string{}
	for _, wrd := range d.Words {
		noStarWrd := wrd
		for i, k := range key {
			if k == '*' {
				noStarWrd = noStarWrd[:i] + "." + noStarWrd[i+1:]
			}
		}
		if strings.ContainsAny(noStarWrd, badChars) {
			continue
		}
		addWord := true
		for i := 0; i < len(filter); i++ {
			k, c := key[i], filter[i]
			if k == '*' && wrd[i] == c {
				continue
			}
			if k == '*' {
				addWord = false
				break
			}
			if wrd[i] == c {
				addWord = false
				break
			}
			if k == '+' && !strings.Contains(wrd, string(c)) {
				addWord = false
				break
			}
		}
		if addWord {
			f = append(f, wrd)
		}
	}

	fd := FilteredDict{
		Key:    key,
		Filter: filter,
		Words:  f,
	}

	w.history = append(w.history, fd)
}

func (w *Wordle) Undo() {
	if len(w.history) <= 1 {
		return
	}

	w.history = w.history[0 : len(w.history)-1]
}
