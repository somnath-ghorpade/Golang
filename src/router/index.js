import { createRouter, createWebHistory } from "vue-router";
import Home from "@/views/Home.vue";
import AddMovie from "@/views/AddMovie.vue";
import Movie from "@/views/Movie.vue";

export default createRouter({
  history: createWebHistory(),

  routes: [
    {
      path: "/",
      name: "home",
      component: Home
    },
    {
      path: "/movie/:id",
      name: "movie",
      component: Movie
    },
    {
      path: "/addMovie",
      name: "addMovie",
      component: AddMovie
    },
    {
      path: "/:catchAll(.*)",
      redirect: { name: "home" }
    }
  ]
});
