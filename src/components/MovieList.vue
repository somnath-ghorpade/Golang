<template>
  <div id="movie-list">
    <div v-if="filteredMovies.length" class="movie-list-holder">
      <movie-item
        v-for="movie in filteredMovies"
        :key="movie.id"
        v-bind:movie="movie"
        v-bind:day="day"
        v-bind:times="times"
      >
        <!-- Slot-->
        <movie-sessions
          v-bind:movie="movie"
          v-bind:day="day"
          v-bind:times="times"
        ></movie-sessions>
      </movie-item>
    </div>
    <div v-else-if="movies.length" class="no-results">
      No results
    </div>
    <div v-else class="no-results">
      Loading...
    </div>
  </div>
</template>

<script>
import MovieItem from "@/components/MovieItem.vue";
import dayjs from "dayjs";
import times from "@/util/times";
import MovieSessions from "@/components/MovieSessions.vue";
export default {
  components: { MovieSessions, MovieItem },
  data: () => ({
    name: process.env.VUE_APP_NAME
  }),
  props: {
    genres: Array,
    movies: Array,
    times: Array,
    day: Object
  },
  methods: {
    genreFilter(movie) {
      if (!this.genres.length) {
        return true;
      } else {
        const movieGenres = movie.movie.Genre.split(", ");
        return movieGenres.some(genre => this.genres.includes(genre));
      }
    },
    dayFilter(movie) {
      return movie.sessions.some(session =>
        dayjs(session.time).isSame(this.day, "day")
      );
    },
    timeFilter(movie) {
      if (!this.times.length || this.times.length === 2) {
        return true;
      } else {
        return movie.sessions.some(session => {
          if (!dayjs(session.time).isSame(this.day, "day")) {
            return false;
          } else if (this.times[0] === times.BEFORE_6PM) {
            // check for before 6 pm
            return dayjs(session.time).hour() < 18;
          } else {
            // after 6 pm
            return dayjs(session.time).hour() >= 18;
          }
        });
      }
    }
  },
  computed: {
    filteredMovies() {
      return (
        this.movies
          // .filter(this.dayFilter)
          .filter(this.genreFilter)
          .filter(this.timeFilter)
      );
    }
  }
};
</script>

<style scoped></style>
