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
            Poster
          </label>
          <input type="text" placeholder="Enter poster url" />
        </div>
        <div class="input-group">
          <label>
            Title
          </label>
          <input type="text" placeholder="Enter movie title" />
        </div>
        <div class="input-group">
          <label>
            Cast
          </label>
          <v-select 
    :options="options"
    :value="selected"
    label="name"
    :multiple="true"
    :taggable="true"
    placeholder="Enter the Cast"
  ></v-select>
        </div>
        <div class="input-group">
          <label>
            Rating
          </label>
          <select v-model="$i18n.rating" 
    placeholder="Enter the Rating"
          >
            <option v-for="(rating, i) in ratings" :key="i" :value="rating">
              {{ rating }}
            </option>
          </select>
        </div>
        <div class="input-group">
          <label>
            Plot
          </label>
          <textarea
            name=""
            rows="2"
            placeholder="Enter the movie plot"
          ></textarea>
        </div>
       
      </div>
      <div class="col-6">
        <div class="input-group">
          <label>Genre</label>
          <div>
            <v-select 
    :options="options"
    :value="selected"
    label="name"
    :multiple="true"
    :taggable="true"
    placeholder="Enter the Genre"
  ></v-select>
          </div>
        </div>
        <div class="input-group">
          <label>
            Director
          </label>
          <v-select 
    :options="options"
    :value="selected"
    label="name"
    :multiple="true"
    :taggable="true"
    placeholder="Enter the Director"
  ></v-select>
        </div>
        <div class="input-group">
          <label>
            Release Date
          </label>
          <input type="date" placeholder="Enter Release Date" />
        </div>
        <div class="input-group">
          <label>
            Running Time
          </label>
          <input type="text" placeholder="Enter Running Time" />
        </div>
       
       
      </div>
    </div>
    <div class="btn-holder">
      <button class="action-btn">Add movie</button>
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import vSelect from 'vue-select';
import 'vue-select/dist/vue-select.css';
export default {
  components: {
    vSelect
  },
  data() {
    return {
      ratings: ["R", "G", "N/A", "PG-13"],

      value: [],
      options: [
        { name: "Vue.js", code: "vu" },
        { name: "Javascript", code: "js" },
        { name: "Open Source", code: "os" }
      ]
    };
  },
  methods: {
    addTag(newTag) {
      const tag = {
        name: newTag,
        code: newTag.substring(0, 2) + Math.floor(Math.random() * 10000000)
      };
      this.options.push(tag);
      this.value.push(tag);
    }
  }
};
</script>
