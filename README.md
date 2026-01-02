# Projet 4 : Nombre Mystère

Ce projet consiste en la création d'un jeu de devinettes "Plus ou Moins" développé en Go. Il se divise en deux étapes : une version console (CLI) et une version web complète.

---

## Phase 1 : Version Terminal (CLI)

La première phase met en œuvre la génération de nombres aléatoires et la logique conditionnelle avancée.

### Objectif
Maîtriser les boucles `for`, les conditions multiples (`if/else`, `switch`) et la gestion des entrées utilisateur pour créer une boucle de jeu interactive.

### Fonctionnalités
* **Configuration personnalisée** : Le joueur définit lui-même la plage de nombres (Minimum et Maximum).
* **Indices contextuels** : Le jeu indique si le nombre proposé est "plus grand" ou "plus petit", avec des nuances de température ("Ça caille", "T'es loin", etc.) selon l'écart.
* **Compteur d'essais** : Suivi du nombre de tentatives nécessaires pour trouver la solution.
* **Gestion d'erreurs** : Vérification que le Min est inférieur au Max et que les saisies sont bien numériques.

### Installation et Lancement
1. Accéder au répertoire :
```bash
cd "Phase 1 - Nombre Mystere CLI"
```
2. Exécuter le programme :
```bash
go run main.go
```

### Utilisation
1. Choisir de jouer dans le menu principal.
2. Définir le nombre minimum et maximum.
3. Proposer un nombre.
4. Lire l'indice et ajuster la proposition jusqu'à la victoire.

---

## Phase 2 : Extension Web (HTML + Go)

Cette phase adapte le jeu pour le web, nécessitant une gestion de l'état du jeu (nombre à trouver, plage, essais) au travers des requêtes HTTP.

### Objectif
Gérer des formulaires de configuration et de jeu, valider les données côté serveur et afficher des retours visuels dynamiques à l'utilisateur.

### Fonctionnalités
* **Configuration de partie** : Formulaire d'accueil pour définir l'intervalle de jeu.
* **Validation robuste** : Messages d'erreur explicites si l'intervalle est invalide (ex: Min > Max).
* **Interface de jeu** : Feedback immédiat sur les propositions (trop grand/petit) et affichage du nombre de tentatives.
* **Redirection et État** : Transition fluide entre la configuration et la phase de jeu.

### Installation et Lancement
1. Accéder au répertoire :
```bash
cd "Phase 2 - Nombre Mystere Web"
```
2. Démarrer le serveur :
```bash
go run main.go
```
3. Accéder à l'application :
Ouvrir un navigateur à l'adresse : `http://localhost:8080`

### Structure du Projet Web
* **main.go** : Point d'entrée du serveur.
* **controller/** : Gestion de la logique du jeu (génération nombre, comparaison, erreurs).
* **router/** : Définition des routes (`/`, `/game`).
* **template/** : Interface HTML dynamique.
* **static/** : Styles CSS.

---

## Annexes : Logique de jeu

Le programme utilise un générateur pseudo-aléatoire pour choisir un nombre entier `N` tel que :
* `Min <= N <= Max`
* Les indices sont basés sur la différence entre la proposition et le nombre mystère.
