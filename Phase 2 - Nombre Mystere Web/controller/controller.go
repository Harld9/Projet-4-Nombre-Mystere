package controller

import (
	"net/http"
	"strconv"
	"text/template"
)

type Question struct {
	Enonce  string
	Choix   []string
	Reponse string
}

var QuestionsFaciles []Question
var QuestionsMoyennes []Question
var QuestionsDifficiles []Question

type Donnees struct {
	IsQuizFacile        bool
	IsQuizMoyen         bool
	IsQuizDifficile     bool
	IsFin               bool
	QuestionsFaciles    []Question
	QuestionsMoyennes   []Question
	QuestionsDifficiles []Question
	Message             string
	MessageFinal        string
	Score               int
	NumQuestion         int
	QuestionSuivante    int
	Couleur0            string
	Couleur1            string
	Couleur2            string
	Couleur3            string
}

func Init() {
	QuestionsFaciles = []Question{
		{Enonce: "Quel constructeur est à l'origine de la marque PlayStation ?", Choix: []string{"A - Nintendo", "B - Sega", "C - Sony", "D - Microsoft"}, Reponse: "C"},
		{Enonce: "Dans quel type de jeu la zone de combat se rétrécit-elle au fil du temps ?", Choix: []string{"A - FIFA", "B - Fortnite", "C - Tetris", "D - Sims"}, Reponse: "B"},
		{Enonce: "Comment s'appelle le héros principal que l'on incarne dans la légende de Zelda ?", Choix: []string{"A - Zelda", "B - Ganondorf", "C - Peach", "D - Link"}, Reponse: "D"},
		{Enonce: "De quelle couleur est le fantôme nommé 'Blinky' dans le jeu Pac-Man ?", Choix: []string{"A - Rouge", "B - Noir", "C - Vert", "D - Marron"}, Reponse: "A"},
		{Enonce: "Quel est le premier bloc de ressources que l'on récupère généralement dans Minecraft ?", Choix: []string{"A - Plastique", "B - Bois", "C - Carton", "D - Tissu"}, Reponse: "B"},
		{Enonce: "Quel est le numéro officiel de Pikachu dans le Pokédex national ?", Choix: []string{"A - 1", "B - 150", "C - 25", "D - 42"}, Reponse: "C"},
		{Enonce: "Dans quel jeu doit-on empiler des briques de formes différentes appelées Tetriminos ?", Choix: []string{"A - Tetris", "B - Pong", "C - Snake", "D - Doom"}, Reponse: "A"},
		{Enonce: "Comment s'appelle la ville principale où se déroule l'action de GTA V ?", Choix: []string{"A - Liberty City", "B - Los Santos", "C - Vice City", "D - San Fierro"}, Reponse: "B"},
		{Enonce: "De quelle espèce animale est inspiré le célèbre personnage Sonic ?", Choix: []string{"A - Chat", "B - Chien", "C - Hérisson", "D - Rat"}, Reponse: "C"},
		{Enonce: "Sur quelle console de salon est sorti le tout premier Super Mario Bros ?", Choix: []string{"A - GameCube", "B - NES", "C - Wii", "D - Switch"}, Reponse: "B"},
	}

	QuestionsMoyennes = []Question{
		{Enonce: "Quel studio de développement japonais a créé la série des Dark Souls ?", Choix: []string{"A - FromSoftware", "B - Ubisoft", "C - Bethesda", "D - Capcom"}, Reponse: "A"},
		{Enonce: "En quelle année la version alpha de Minecraft a-t-elle été rendue publique ?", Choix: []string{"A - 2005", "B - 2009", "C - 2011", "D - 2013"}, Reponse: "B"},
		{Enonce: "Comment s'appelle l'intelligence artificielle qui accompagne le Major dans Halo ?", Choix: []string{"A - Alexa", "B - Siri", "C - Cortana", "D - Glados"}, Reponse: "C"},
		{Enonce: "Quel est le prénom de la fille adoptive de Geralt de Riv dans The Witcher ?", Choix: []string{"A - Yennefer", "B - Triss", "C - Ciri", "D - Keira"}, Reponse: "C"},
		{Enonce: "Lequel de ces jeux était le Battle Royale le plus populaire juste avant l'arrivée de Fortnite ?", Choix: []string{"A - H1Z1", "B - PUBG", "C - DayZ", "D - Apex"}, Reponse: "B"},
		{Enonce: "Comment s'appelle l'univers fantastique dans lequel se déroule League of Legends ?", Choix: []string{"A - Azeroth", "B - Runeterra", "C - Sanctuaire", "D - Hyrule"}, Reponse: "B"},
		{Enonce: "Dans la série Metal Gear Solid, quel est le véritable prénom de Solid Snake ?", Choix: []string{"A - John", "B - Jack", "C - David", "D - Liquid"}, Reponse: "C"},
		{Enonce: "Quel constructeur a lancé la console Dreamcast avant de quitter le marché des consoles ?", Choix: []string{"A - Nintendo", "B - Sony", "C - Sega", "D - Atari"}, Reponse: "C"},
		{Enonce: "Dans quel jeu affronte-t-on des ennemis redoutables appelés les 'Claqueurs' (Clickers) ?", Choix: []string{"A - Resident Evil", "B - Dying Light", "C - Last of Us", "D - Days Gone"}, Reponse: "C"},
		{Enonce: "Quel moteur de jeu célèbre a été développé par la société Epic Games ?", Choix: []string{"A - Unity", "B - Frostbite", "C - CryEngine", "D - Unreal"}, Reponse: "D"},
	}

	QuestionsDifficiles = []Question{
		{Enonce: "Quel jeu de borne d'arcade sorti en 1971 est considéré comme le premier jeu vidéo commercial ?", Choix: []string{"A - Pong", "B - Computer Space", "C - Invaders", "D - Pac-Man"}, Reponse: "B"},
		{Enonce: "Lequel de ces créateurs de génie est derrière la saga Metal Gear et Death Stranding ?", Choix: []string{"A - Mikami", "B - Miyazaki", "C - Kojima", "D - Miyamoto"}, Reponse: "C"},
		{Enonce: "Comment se nomme la langue parlée par les dragons dans l'univers de Skyrim ?", Choix: []string{"A - Dovahzul", "B - Draconique", "C - Thalmor", "D - Daedrique"}, Reponse: "A"},
		{Enonce: "Quel était le nom de code de la Nintendo GameCube durant sa phase de développement ?", Choix: []string{"A - Reality", "B - Atlantis", "C - Dolphin", "D - Revolution"}, Reponse: "C"},
		{Enonce: "Dans quelle licence Sora voyage-t-il aux côtés de Donald et Dingo ?", Choix: []string{"A - Final Fantasy", "B - Kingdom Hearts", "C - Dragon Quest", "D - NieR"}, Reponse: "B"},
		{Enonce: "Quel projet de jeu vidéo détient le record du plus gros budget grâce au financement participatif ?", Choix: []string{"A - GTA V", "B - RDR 2", "C - Star Citizen", "D - Cyberpunk"}, Reponse: "C"},
		{Enonce: "En quelle année la console Super Nintendo (SNES) est-elle officiellement sortie en Europe ?", Choix: []string{"A - 1990", "B - 1991", "C - 1992", "D - 1993"}, Reponse: "C"},
		{Enonce: "Quel compositeur légendaire a créé la majorité des musiques de la saga Final Fantasy ?", Choix: []string{"A - Kondo", "B - Uematsu", "C - Yamaoka", "D - Shimomura"}, Reponse: "B"},
		{Enonce: "Comment s'appelle le tout premier boss que l'on rencontre dans le tutoriel de Dark Souls 1 ?", Choix: []string{"A - Le Démon de l'Asile", "B - Les Gargouilles", "C - Queelag", "D - Ornstein"}, Reponse: "A"},
		{Enonce: "Quel virus créé par Umbrella Corp est responsable du désastre de Raccoon City ?", Choix: []string{"A - Le Virus G", "B - Le Virus T", "C - Las Plagas", "D - Le Virus C"}, Reponse: "B"},
	}
}
func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, Donnees{})
}

func QuizFacile(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("next_question"))
	score, _ := strconv.Atoi(r.FormValue("score"))

	if page >= len(QuestionsFaciles) {
		var messageFinal string
		if score <= 4 {
			messageFinal = "Vous êtes nul revoyez vos basiques, vous avez eu un score de"
		} else if score <= 7 {
			messageFinal = "Vous êtes Mid, vous avez eu un score de"
		} else if score <= 9 {
			messageFinal = "Vous êtes fort,bien joué à vous, vous avez eu un score de"
		} else {
			messageFinal = "Vous êtes le GOAT, bien joué à vous, vous avez eu un score de"
		}
		data := Donnees{IsFin: true, Score: score, MessageFinal: messageFinal}
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}

	data := Donnees{
		IsQuizFacile:     true,
		QuestionsFaciles: QuestionsFaciles,
		NumQuestion:      page,
		Score:            score,
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func QuizMoyen(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("next_question"))
	score, _ := strconv.Atoi(r.FormValue("score"))

	if page >= len(QuestionsMoyennes) {
		var messageFinal string
		if score <= 4 {
			messageFinal = "Vous êtes nul revoyez vos basiques, vous avez eu un score de"
		} else if score <= 7 {
			messageFinal = "Vous êtes Mid, vous avez eu un score de"
		} else if score <= 9 {
			messageFinal = "Vous êtes fort,bien joué à vous, vous avez eu un score de"
		} else {
			messageFinal = "Vous êtes le GOAT, bien joué à vous, vous avez eu un score de"
		}
		data := Donnees{IsFin: true, Score: score, MessageFinal: messageFinal}
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}

	data := Donnees{
		IsQuizMoyen:       true,
		QuestionsMoyennes: QuestionsMoyennes,
		NumQuestion:       page,
		Score:             score,
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func QuizDifficile(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.FormValue("next_question"))
	score, _ := strconv.Atoi(r.FormValue("score"))

	if page >= len(QuestionsDifficiles) {
		var messageFinal string
		if score <= 4 {
			messageFinal = "Vous êtes nul revoyez vos basiques, vous avez eu un score de"
		} else if score <= 7 {
			messageFinal = "Vous êtes Mid, vous avez eu un score de"
		} else if score <= 9 {
			messageFinal = "Vous êtes fort,bien joué à vous, vous avez eu un score de"
		} else {
			messageFinal = "Vous êtes le GOAT, bien joué à vous, vous avez eu un score de"
		}
		data := Donnees{IsFin: true, Score: score, MessageFinal: messageFinal}
		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
		return
	}

	data := Donnees{
		IsQuizDifficile:     true,
		QuestionsDifficiles: QuestionsDifficiles,
		NumQuestion:         page,
		Score:               score,
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Reponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		choix := r.FormValue("choix_utilisateur")
		idStr := r.FormValue("question_id")
		scoreStr := r.FormValue("current_score")
		typeQuiz := r.FormValue("quiz_type")

		i, _ := strconv.Atoi(idStr)
		score, _ := strconv.Atoi(scoreStr)

		data := Donnees{}

		if typeQuiz == "facile" {
			data.IsQuizFacile = true
			data.QuestionsFaciles = QuestionsFaciles

			bonneReponse := QuestionsFaciles[i].Reponse
			lettreChoisie := string(choix[0])

			if lettreChoisie == bonneReponse {
				data.Message = "Bien joué !"
				score = score + 1
			} else {
				data.Message = "Loupé ! La réponse était : " + bonneReponse
			}

			data.NumQuestion = i
			data.QuestionSuivante = i + 1
			data.Score = score

			if string(QuestionsFaciles[i].Choix[0][0]) == bonneReponse {
				data.Couleur0 = "boutonvrai"
			} else {
				data.Couleur0 = "boutonfaux"
			}
			if string(QuestionsFaciles[i].Choix[1][0]) == bonneReponse {
				data.Couleur1 = "boutonvrai"
			} else {
				data.Couleur1 = "boutonfaux"
			}
			if string(QuestionsFaciles[i].Choix[2][0]) == bonneReponse {
				data.Couleur2 = "boutonvrai"
			} else {
				data.Couleur2 = "boutonfaux"
			}
			if string(QuestionsFaciles[i].Choix[3][0]) == bonneReponse {
				data.Couleur3 = "boutonvrai"
			} else {
				data.Couleur3 = "boutonfaux"
			}
		}

		if typeQuiz == "moyen" {
			data.IsQuizMoyen = true
			data.QuestionsMoyennes = QuestionsMoyennes

			bonneReponse := QuestionsMoyennes[i].Reponse
			lettreChoisie := string(choix[0])

			if lettreChoisie == bonneReponse {
				data.Message = "Bien joué !"
				score = score + 1
			} else {
				data.Message = "Loupé ! La réponse était : " + bonneReponse
			}

			data.NumQuestion = i
			data.QuestionSuivante = i + 1
			data.Score = score

			if string(QuestionsMoyennes[i].Choix[0][0]) == bonneReponse {
				data.Couleur0 = "boutonvrai"
			} else {
				data.Couleur0 = "boutonfaux"
			}
			if string(QuestionsMoyennes[i].Choix[1][0]) == bonneReponse {
				data.Couleur1 = "boutonvrai"
			} else {
				data.Couleur1 = "boutonfaux"
			}
			if string(QuestionsMoyennes[i].Choix[2][0]) == bonneReponse {
				data.Couleur2 = "boutonvrai"
			} else {
				data.Couleur2 = "boutonfaux"
			}
			if string(QuestionsMoyennes[i].Choix[3][0]) == bonneReponse {
				data.Couleur3 = "boutonvrai"
			} else {
				data.Couleur3 = "boutonfaux"
			}
		}

		if typeQuiz == "difficile" {
			data.IsQuizDifficile = true
			data.QuestionsDifficiles = QuestionsDifficiles

			bonneReponse := QuestionsDifficiles[i].Reponse
			lettreChoisie := string(choix[0])

			if lettreChoisie == bonneReponse {
				data.Message = "Bien joué !"
				score = score + 1
			} else {
				data.Message = "Loupé ! La réponse était : " + bonneReponse
			}

			data.NumQuestion = i
			data.QuestionSuivante = i + 1
			data.Score = score

			if string(QuestionsDifficiles[i].Choix[0][0]) == bonneReponse {
				data.Couleur0 = "boutonvrai"
			} else {
				data.Couleur0 = "boutonfaux"
			}
			if string(QuestionsDifficiles[i].Choix[1][0]) == bonneReponse {
				data.Couleur1 = "boutonvrai"
			} else {
				data.Couleur1 = "boutonfaux"
			}
			if string(QuestionsDifficiles[i].Choix[2][0]) == bonneReponse {
				data.Couleur2 = "boutonvrai"
			} else {
				data.Couleur2 = "boutonfaux"
			}
			if string(QuestionsDifficiles[i].Choix[3][0]) == bonneReponse {
				data.Couleur3 = "boutonvrai"
			} else {
				data.Couleur3 = "boutonfaux"
			}
		}

		tmpl := template.Must(template.ParseFiles("template/index.html"))
		tmpl.Execute(w, data)
	}
}
