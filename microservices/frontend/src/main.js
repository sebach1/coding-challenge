// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import { MdTable, MdContent } from 'vue-material/dist/components'
import 'vue-material/dist/vue-material.min.css'

Vue.use(MdContent)
Vue.use(MdTable)

Vue.config.productionTip = false

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');
