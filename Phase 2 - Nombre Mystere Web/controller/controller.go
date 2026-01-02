package controller

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
)

var (
	MinInt        int
	MaxInt        int
	NombreMystere int
	Tentative     int
)

type Donnees struct {
	IsPlage        bool
	IsGame         bool
	IsEnd          bool
	ErreurMin      string
	ErreurMax      string
	ErreurPlage    string
	ErreurNbChoisi string
	MessageStatut  string
	Plage          string
	MessageFin     string
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := Donnees{
		IsPlage:        true,
		IsGame:         false,
		IsEnd:          false,
		ErreurMin:      "",
		ErreurMax:      "",
		ErreurPlage:    "",
		ErreurNbChoisi: "",
		MessageStatut:  "",
		MessageFin:     "",
	}

	if r.Method == http.MethodPost {

		if r.FormValue("reset") == "recommencer" {
			Tentative = 0

			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, data)
			return
		}

		Tentative = 0
		Min := r.FormValue("min")
		Max := r.FormValue("max")

		var errmin error
		MinInt, errmin = strconv.Atoi(Min)

		if errmin != nil || MinInt < 0 {
			data.ErreurMin = "Veuillez entrer un minimum valide"
		} else {
			data.ErreurMin = ""
		}

		var errmax error
		MaxInt, errmax = strconv.Atoi(Max)

		if errmax != nil || MaxInt <= MinInt {
			data.ErreurMax = "Le maximum doit être strictement supérieur au minimum"
		} else {
			data.ErreurMax = ""
		}

		if data.ErreurMin == "" && data.ErreurMax == "" {
			NombreMystere = rand.Intn(MaxInt-MinInt+1) + MinInt

			data.IsPlage = false

			http.Redirect(w, r, "/game", http.StatusSeeOther)
			return
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Game(w http.ResponseWriter, r *http.Request) {
	data := Donnees{
		IsPlage:        false,
		IsGame:         true,
		IsEnd:          false,
		ErreurNbChoisi: "",
		MessageStatut:  "",
		Plage:          "",
	}

	data.Plage = fmt.Sprintf("%d et %d", MinInt, MaxInt)

	if r.Method == http.MethodPost {
		NbChoisi := r.FormValue("nbchoisi")
		NbChoisiInt, err := strconv.Atoi(NbChoisi)

		if err != nil {
			data.MessageStatut = "Ce n'est pas un nombre valide."
		} else if NbChoisiInt < MinInt || NbChoisiInt > MaxInt {
			data.MessageStatut = "Le nombre doit être entre " + strconv.Itoa(MinInt) + " et " + strconv.Itoa(MaxInt)
		} else if NbChoisiInt < NombreMystere {
			data.MessageStatut = "C'est plus !"
			Tentative++
		} else if NbChoisiInt > NombreMystere {
			data.MessageStatut = "C'est moins !"
			Tentative++
		} else {
			Tentative++
			http.Redirect(w, r, "/end", http.StatusSeeOther)
			return
		}
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func End(w http.ResponseWriter, r *http.Request) {
	data := Donnees{
		IsPlage:    false,
		IsGame:     false,
		IsEnd:      true,
		MessageFin: "",
	}
	switch {
	case Tentative == 1:
		data.MessageFin = "🏆 Wesh ?! T'as triché ou t'es devin ? GG du premier coup !"
	case Tentative <= 5:
		data.MessageFin = fmt.Sprintf("🚀 Propre, efficace, carré. T'as géré en %d coups!", Tentative)
	case Tentative <= 10:
		data.MessageFin = fmt.Sprintf("👏 Pas mal, mais peut mieux faire. T'as gagné quand même. En %d coup", Tentative)
	case Tentative <= 10:
		data.MessageFin = fmt.Sprintf("Peut mieux faire, tu as réussi en %d coups", Tentative)
	default:
		return
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
