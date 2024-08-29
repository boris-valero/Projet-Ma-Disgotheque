# Ma Disgotheque


## Description

"ma-disgotheque" est une application web permettant de gérer une collection d'albums de musique. Elle utilise le framework Gin pour le routage HTTP et gère les requêtes CORS.

## Fonctionnalités

- **Lister les albums** : Récupérer la liste complète des albums.
- **Ajouter un album** : Ajouter un nouvel album à la collection.
- **Récupérer un album par ID** : Récupérer les détails d'un album spécifique en utilisant son ID.

## Structure du projet

```
go.mod
go.sum
main
main.go
main.go.back
main2.go.back
README.md
```

## Installation

1. Clonez le dépôt :
    ```sh
    git clone <URL_DU_DEPOT>
    cd Projet-Ma-Disgotheque
    ```

2. Installez les dépendances :
    ```sh
    go mod tidy
    ```

## Utilisation

1. Lancez l'application :
    ```sh
    go run main.go
    ```

2. L'application sera accessible à l'adresse suivante :
    ```
    http://localhost:8080
    ```

## Points de terminaison API

### GET /albums

Récupère la liste de tous les albums.

**Réponse :**
```json
[
    {
        "id": "1",
        "title": "Album Title",
        "artist": "Artist Name",
        "year": 2021,
    },
    ...
]
```

### POST /albums

Ajoute un nouvel album.

**Corps de la requête :**
```json
{
    "id": "2",
    "title": "New Album",
    "artist": "New Artist",
    "year": 2022,
}
```

### GET /albums/:id

Récupère les détails d'un album spécifique par son ID.

**Réponse :**
```json
{
    "id": "1",
    "title": "Album Title",
    "artist": "Artist Name",
    "year": 2021,
}
