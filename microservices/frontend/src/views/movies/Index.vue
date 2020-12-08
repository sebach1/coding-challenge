<template>
  <div>

    <md-table v-if="films" v-model="films"  md-sort="title"  md-card>
      <md-table-toolbar>
        <h1 class="md-title">Films</h1>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="ID" md-sort-by="film.id" md-numeric>{{ item.film.id }}</md-table-cell>
        <md-table-cell md-label="Title" md-sort-by="film.title">{{ item.film.title }}</md-table-cell>
        <md-table-cell md-label="Description" md-sort-by="film.description">{{ item.film.description }}</md-table-cell>
        <md-table-cell md-label="Director" md-sort-by="film.directorName">{{ item.film.directorName }}</md-table-cell>
        <md-table-cell md-label="Producer" md-sort-by="film.producerName">{{ item.film.producerName }}</md-table-cell>
        <md-table-cell md-label="Release year" md-sort-by="film.releaseYear">{{ item.film.releaseYear }}</md-table-cell>
        <md-table-cell md-label="Rating" md-sort-by="film.rating">{{ item.film.rating }}</md-table-cell>
        <md-table-cell md-label="People">
          {{ displayPeople(item.people) }}
        </md-table-cell>
      </md-table-row>
    </md-table>

  </div>
</template>
<script>
import axios from "axios"
export default {
  data() {
    return {
      films: undefined,
    };
  },
mounted() {
    this.fetchFilms();
  },
  methods: {
    fetchFilms() {
      fetch('http://localhost:10000/films')
        .then(response => response.json())
        .then((films) =>{
         this.films = films.films;
        }
      )},
    displayPeople(ppl){
      if (!ppl){
        return ""
      }
      return Object.values(ppl).map(x => x.name).join(", ");
    },
   }
  }

        // <md-icon>movie</md-icon>

</script>
