<template>
  <div id="content">
    <div class="sidebar">
      <day-select v-bind:day="day" v-on:day-selected="daySelected" />
    </div>
    <div class="main-content">
      <!-- Content here-->
      <router-view
        v-bind:genres="genres"
        v-bind:movies="movies"
        v-bind:day="day"
        v-bind:times="times"
        v-on:check-filter="checkFilter"
        v-on:day-selected="daySelected"
        v-slot="{ Component }"
      >
        <keep-alive>
          <component :is="Component"></component>
        </keep-alive>
      </router-view>
    </div>
  </div>
</template>
<script>
  /* eslint-disable */
  import axios from "axios"
  import dayjs from "dayjs"

  import DaySelect from "@/components/DaySelect.vue"
  export default {
    components: {
      DaySelect,
    },
    data() {
      return {
        genres: [],
        times: [],
        movies: [],
        day: dayjs(),
        showFilter: true,
        basicInfoData: {
          name: "vinanya",
          qualification: "ME Comp",
        },
      }
    },
    methods: {
      setBasicData() {
        this.$store.dispatch("actionBasicInfo", this.basicInfoData)
      },
      addCounter() {
        this.$store.dispatch("increment", 200)
      },
      decrement() {
        this.$store.dispatch("decrement", 200)
      },
      checkFilter(checked, title, group) {
        if (checked) {
          this[group].push(title)
        } else {
          this[group] = this[group].filter((item) => item !== title)
        }
      },
      daySelected(day) {
        console.log("router", day)
        this.day = day
      },
    },
    async created() {
      var config = {
        method: "get",
        url: "/o/fetchMovies",
        headers: {
          Authorization: "Basic dmluYXlhOnZpbmF5YUAxMTIy",
        },
      }
      const { data } = await axios(config)
      this.movies = data
    },
  }
</script>
<style></style>
