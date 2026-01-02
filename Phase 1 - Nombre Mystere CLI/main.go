package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	var menuchoice int
	var videur string
	var nbmin int
	var nbmax int
	var nbchoisi int
	sleep := time.Duration(2) * time.Second

	for {
		ClearScreen()

		tentative := 0

		fmt.Println("Bonjour et bienvenue au jeu du nombre mystère")
		fmt.Println("Voulez-vous jouez ?")
		fmt.Println("1 - Oui")
		fmt.Println("2 - Non")

		_, err := fmt.Scan(&menuchoice)

		if err != nil {
			fmt.Println("❌ Choix impossible, réessayez.")
			fmt.Scanln(&videur)
			time.Sleep(sleep)
			continue
		}

		switch menuchoice {

		case 1:
			ClearScreen()

			fmt.Println("Définissez la tranche dans laquelle vous voulez jouer :")
			fmt.Println("Nombre minimum :")

			_, errmin := fmt.Scan(&nbmin)

			if errmin != nil {
				fmt.Println("❌ Choix impossible, réessayez.")
				fmt.Scanln(&videur)
				time.Sleep(sleep)
				continue
			}

			fmt.Println("Nombre maximum :")

			_, errmax := fmt.Scan(&nbmax)

			if errmax != nil || nbmin == nbmax || nbmax < nbmin {
				fmt.Println("❌ Choix impossible, réessayez. (Le max ne peut pas être égal au min ni lui être inférieur.)")
				fmt.Scanln(&videur)
				time.Sleep(sleep)
				continue
			}
			plage := nbmax - nbmin + 1
			nbmystere := rand.IntN(plage) + nbmin
		jeu:

			for {

				fmt.Println("Devinez le nombre mystère dans la nbmax que vous avez choisie :")
				fmt.Println("Quel nombre pensez vous que le nombre mystère est ?")
				_, errchoix := fmt.Scan(&nbchoisi)
				if errchoix != nil || nbchoisi > nbmax {
					fmt.Println("❌ Choix impossible, réessayez.")
					fmt.Scanln(&videur)
					time.Sleep(sleep)
					continue
				}
				tentative++

				ecart := nbchoisi - nbmystere

				if ecart < 0 {
					ecart = ecart * -1
				}

				switch {

				case nbmystere > nbchoisi && float64(ecart) > (0.75*float64(nbmax)):
					fmt.Println("🥶 Frérot t'es à l'ouest complet... C'est beaucoup plus grand !")
					Separator()
				case nbmystere < nbchoisi && float64(ecart) > (0.75*float64(nbmax)):
					fmt.Println("🥶 T'as séché les cours de maths ? C'est beaucoup plus petit !")
					Separator()

				case nbmystere > nbchoisi && float64(ecart) <= (0.75*float64(nbmax)) && float64(ecart) > (0.50*float64(nbmax)):
					fmt.Println("❄️ Ça caille ici, t'es loin. C'est plus grand.")
					Separator()
				case nbmystere < nbchoisi && float64(ecart) <= (0.75*float64(nbmax)) && float64(ecart) > (0.50*float64(nbmax)):
					fmt.Println("❄️ Mets un pull, t'es loin. C'est plus petit.")
					Separator()

				case nbmystere > nbchoisi && float64(ecart) <= (0.50*float64(nbmax)) && float64(ecart) > (0.25*float64(nbmax)):
					fmt.Println("🌡️ Mouais, ça passe, mais t'es pas encore dessus. C'est plus grand.")
					Separator()
				case nbmystere < nbchoisi && float64(ecart) <= (0.50*float64(nbmax)) && float64(ecart) > (0.25*float64(nbmax)):
					fmt.Println("🌡️ On se rapproche doucement... mais c'est plus petit.")
					Separator()

				case nbmystere > nbchoisi && float64(ecart) <= (0.25*float64(nbmax)) && float64(ecart) > (0.10*float64(nbmax)):
					fmt.Println("🔥 Là on commence à discuter ! C'est plus grand !")
					Separator()
				case nbmystere < nbchoisi && float64(ecart) <= (0.25*float64(nbmax)) && float64(ecart) > (0.10*float64(nbmax)):
					fmt.Println("🔥 Chaud cacao ! C'est plus petit !")
					Separator()

				case nbmystere > nbchoisi:
					fmt.Println("🥵 T'es littéralement dessus (ou presque) ! Un poil plus grand !")
					Separator()
				case nbmystere < nbchoisi:
					fmt.Println("🥵 A deux doigts de la gloire ! Un poil plus petit !")
					Separator()

				default:
					switch {
					case tentative == 1:
						fmt.Println("🏆 Wesh ?! T'as triché ou t'es devin ? GG du premier coup !")
						Separator()
					case tentative <= 5:
						fmt.Println("🚀 Propre, efficace, carré. T'as géré !")
						Separator()
					case tentative <= 10:
						fmt.Println("👏 Pas mal, mais peut mieux faire. T'as gagné quand même.")
						Separator()
					default:
						fmt.Println("😅 Enfin ! J'ai failli m'endormir devant mon écran... Mais bravo.")
						Separator()
					}
					fmt.Println("🎉 T'as plié le game en", tentative, "coups !")
					Separator()
					break jeu
				}
			}

			time.Sleep(sleep)
			ClearScreen()
			fmt.Println("Veux tu continuez à jouer ?")
			fmt.Println("1 - Oui")
			fmt.Println("2 - Non")

			_, err := fmt.Scan(&menuchoice)
			if err != nil {
				fmt.Println("❌ Choix impossible, réessayez.")
				fmt.Scanln(&videur)
				time.Sleep(sleep)
				continue

			}
		}
	}
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Separator() {
	fmt.Println("\n==================================================")
	fmt.Println("")
}
