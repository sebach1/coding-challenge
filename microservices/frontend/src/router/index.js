import Vue from 'vue'
import Router from 'vue-router'
import MoviesIndex from '@/views/movies/Index.vue'
Vue.use(Router)

export default new Router({
  mode: 'history',

  routes: [
    {
      path: "/movies",
      name: "JPFilms - Movies",
      component: MoviesIndex,
    },
    {
      path: "*",
      redirect: "/movies",
    },
  ]
})
