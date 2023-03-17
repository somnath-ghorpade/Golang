<template>
  <div id="overview">
    <div id="main" class="main">
      <div id="title">
        <!-- <img src="logo.png" /> -->
        <h1>{{ $t("vueProjectName") }}</h1>
        <div class="box-filter">
          <router-link to="/addMovie" class="title">
            Add Movie
          </router-link>
          <div class="title" @click="showFilterSidebar = true">
            Filter
          </div>

          <!-- I18n implementation -->
          <!-- Change language using locals -->
          <changeLocale></changeLocale>
          <!-- Ends here -->
        </div>
      </div>
      <movie-list
        v-bind:genres="genres"
        v-bind:movies="movies"
        v-bind:day="day"
        v-bind:times="times"
      ></movie-list>
      <!-- Sidebar showing filter by genre and before and after 6pm -->
      <div class="filter-sidebar" :class="showFilterSidebar ? 'active' : ''">
        <div class="heading">
          <h2>Filter</h2>
          <div
            class="close-icon"
            @click="showFilterSidebar = !showFilterSidebar"
          >
            &times;
          </div>
        </div>
        <movie-filter v-on:check-filter="checkFilter"></movie-filter>
      </div>
      <!-- Ends here -->
    </div>
  </div>
</template>

<script>
/* eslint-disable */
import MovieList from "@/components/MovieList.vue";
import MovieFilter from "@/components/MovieFilter.vue";
import changeLocale from "@/components/ChangeLocale.vue";

export default {
  name: "Home-Page",
  props: {
    genres: Array,
    movies: Array,
    day: Object,
    times: Array
  },
  emits: ["check-filter", "day-selected"],
  methods: {
    checkFilter(checked, title, group) {
      this.$emit("check-filter", checked, title, group);
    },
    daySelected(day) {
      console.log(day);
      this.$emit("day-selected", day);
    }
  },
  components: {
    MovieList,
    MovieFilter,
    changeLocale
  },
  data() {
    return {
      showFilterSidebar: true
    };
  }
};
</script>
