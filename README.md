# Music Recommender - Medium Story Example

This is a simple music recommender system example that I've created for my Medium story. You can find the story [here](https://9ssi7.medium.com/building-a-music-recommendation-app-with-golang-neo4j-and-graphql-697f842ea688).

## Running the project

Run the following command to start the project:

```bash
make dev
```

and then you can access the GraphQL playground from [http://localhost:4000](http://localhost:4000).

## Mutations

Make sure to replace `<user-id>`, `<genre-id>` and `<song-id>` with the actual ids.

### Create a user

```graphql
mutation Mutation {
  createUser(input: {userName: "9ssi7", email: "info@ssibrahimbas.com", favoriteGenres: []}) {
    id
    userName
  }
}
```

### Create a genre

```graphql
mutation Mutation {
  createGenre(input: {name: "rap"}) {
    id
    name
  }
}
```

### Create a song

```graphql
mutation Mutation {
  createSong(input: {
    title:"Sarı Saçlım Mavi Gözlüm",
    artist:"Barış Manço",
    genreId:"<genre-id>"
  }) {
    id
    title
    artist
  }
}
```

### Mark Song as Listened

```graphql
mutation Mutation {
  markSongAsListened(userId: "<user-id>", songId: "<song-id>")
}
```

## Queries

### Get User

```graphql
query Query {
  user(email: "info@ssibrahimbas.com") {
    id
    email
    userName
  }
}
```

### Get Songs

```graphql
query Query {
   songs{
    id
    title
    artist
  }
}
```

### Get Genres

```graphql
query Query {
  genres {
    id
    name
  }
}
```

### Get Recommendations

```graphql
query Query {
   songsRecommendation(userId: "<user-id>"){
    id
    title
    artist
  }
}
```