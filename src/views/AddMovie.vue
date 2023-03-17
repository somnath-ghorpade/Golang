<template>
  <div id="addMovie">
    <div class="home">
      <router-link v-bind:to="{ name: 'home' }">Back to home</router-link>
    </div>
    <h2>
      Add New Movie
    </h2>
    <div class="row">
      <div class="col-6">
        <div class="input-group">
          <label>
           ImdbId
          </label>
          <input v-model="movie.imdbId" type="text" placeholder="Enter imdbId" />
        </div>
        <div class="input-group">
          <label>
            Poster
          </label>
          <input v-model="movie.poster" type="text" placeholder="Enter poster url" />
        </div>
        <div class="input-group">
          <label>
            Title
          </label>
          <input type="text" v-model="movie.title" placeholder="Enter movie title" />
        </div>
        <div class="input-group">
          <label>
            Actors
          </label>
          <v-select
            v-model="movie.cast"
            :options="actors"
            :value="selected"
            label="name"
            :taggable="true"
            placeholder="Enter the Cast"
          ></v-select>
        </div>
        <div class="input-group">
          <label>
            Rating
          </label>
          <v-select
            v-model="movie.rating"
            :options="ratings"
            :value="selected"
            label="Source"
            :taggable="true"
            placeholder="Enter the Cast"
          ></v-select>
        </div>
        <div class="input-group">
          <label>
            Plot
          </label>
          <textarea name="" v-model="movie.plot" rows="2" placeholder="Enter the movie plot"></textarea>
        </div>
      </div>
      <div class="col-6">
        <div class="input-group">
          <label>Genre</label>
          <div>
            <v-select
              :options="genres"
              :value="selected"
              label="name"
              :taggable="true"
              v-model="movie.genre"
              placeholder="Enter the Genre"
            ></v-select>
          </div>
        </div>
        <div class="col-6">
          <div class="input-group">
            <label>Sessions</label>
            <div>
              <v-select
                :options="sessions"
                :value="selected"
                label="id"
                :multiple="true"
                :taggable="true"
                v-model="movie.sessions"
                placeholder="Select the sessions"
              ></v-select>
            </div>
          </div>
        </div>
        <div class="input-group">
          <label>
            Director
          </label>
          <input type="input" v-model="movie.director" placeholder="Enter Release Date" />
        </div>
        <div class="input-group">
          <label>
            Release Date
          </label>
          <input type="date" v-model="movie.releaseDate" placeholder="Enter Release Date" />
        </div>
        <div class="input-group">
          <label>
            Running Time
          </label>
          <input type="text" v-model="movie.running" placeholder="Enter Running Time" />
        </div>
      </div>
    </div>
    <div class="btn-holder">
      <button class="action-btn" @click="AddMovie">Add movie</button>
    </div>
  </div>
</template>

<script>
  /* eslint-disable */
  import axios from "axios"
  import vSelect from "vue-select"
  import "vue-select/dist/vue-select.css"
  export default {
    components: {
      vSelect,
    },
    data() {
      return {
        actors: ["Mikela Jay", "Rob Beckwermert", "Christopher Gora", "Nina Jones", "Ritesh, Genelia"],
        value: [],
        // https://i0.wp.com/boxofficeworldwide.com/wp-content/uploads/2022/12/img_4027.jpg?fit=1267%2C1361&ssl=1
        options: [],
        ratings: [
          {
            Source: "Internet Movie Database",
            Value: "8.1/10",
          },
          {
            Source: "Rotten Tomatoes",
            Value: "90%",
          },
          {
            Source: "Metacritic",
            Value: "73/100",
          },
        ],
        sessions: [
          {
            id: "tt0379225_0",
            time: "2022-12-02T09:45:00.000Z",
            seats: 175,
          },
          {
            id: "tt0379225_1",
            time: "2022-12-02T12:00:00.000Z",
            seats: 178,
          },
          {
            id: "tt0379225_2",
            time: "2022-12-02T14:00:00.000Z",
            seats: 181,
          },
          {
            id: "tt0379225_2",
            time: "2022-12-03T08:30:00.000Z",
            seats: 185,
          },
          {
            id: "tt0379225_3",
            time: "2022-12-03T08:30:00.000Z",
            seats: 190,
          },
          {
            id: "tt0379225_0",
            time: "2022-12-03T12:00:00.000Z",
            seats: 175,
          },
          {
            id: "tt0379225_4",
            time: "2022-12-03T12:00:00.000Z",
            seats: 195,
          },
          {
            id: "tt0379225_1",
            time: "2022-12-03T14:00:00.000Z",
            seats: 180,
          },
          {
            id: "tt0379225_2",
            time: null,
            seats: 185,
          },
          {
            id: "tt0379225_5",
            time: null,
            seats: 210,
          },
          {
            id: "tt0379225_4",
            time: "2022-12-04T07:30:00.000Z",
            seats: 203,
          },
          {
            id: "tt0379225_1",
            time: "2022-12-04T08:30:00.000Z",
            seats: 182,
          },
          {
            id: "tt0379225_2",
            time: "2022-12-04T08:30:00.000Z",
            seats: 189,
          },
          {
            id: "tt0379225_3",
            time: "2022-12-04T12:00:00.000Z",
            seats: 196,
          },
          {
            id: "tt0379225_6",
            time: "2022-12-04T12:00:00.000Z",
            seats: 217,
          },
          {
            id: "tt0379225_0",
            time: "2022-12-04T14:00:00.000Z",
            seats: 175,
          },
          {
            id: "tt0379225_0",
            time: "2022-12-05T08:30:00.000Z",
            seats: 175,
          },
          {
            id: "tt0379225_1",
            time: "2022-12-05T08:30:00.000Z",
            seats: 177,
          },
          {
            id: "tt0379225_0",
            time: "2022-12-06T08:30:00.000Z",
            seats: 175,
          },
          {
            id: "tt0379225_1",
            time: "2022-12-06T12:00:00.000Z",
            seats: 177,
          },
          {
            id: "tt0379225_1",
            time: "2022-12-07T07:30:00.000Z",
            seats: 180,
          },
          {
            id: "tt0379225_0",
            time: "2022-12-07T12:00:00.000Z",
            seats: 175,
          },
          {
            id: "tt0379225_3",
            time: "2022-12-07T12:00:00.000Z",
            seats: 190,
          },
          {
            id: "tt0379225_4",
            time: "2022-12-07T14:00:00.000Z",
            seats: 195,
          },
          {
            id: "tt0379225_0",
            time: "2022-12-08T07:30:00.000Z",
            seats: 175,
          },
        ],
        genres: ["Documentary ,History", "Animation, Comedy, Fantasy", "Animation, Short, Fantasy", "Comedy, Drama"],
        movie: {},
      }
    },
    methods: {
      addTag(newTag) {
        const tag = {
          name: newTag,
          code: newTag.substring(0, 2) + Math.floor(Math.random() * 10000000),
        }
        this.options.push(tag)
        this.value.push(tag)
      },
      AddMovie() {
        console.log("movie", this.movie)
        axios
          .post("/o/insertMovie", { id:this.movie.imdbId,movie: this.movie, sessions: this.movie.sessions })
          .then(function(response) {
            console.log(response)
          })
          .catch(function(error) {
            console.log(error)
          })
      },
    },
  }
</script>
