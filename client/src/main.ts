import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'

Vue.config.productionTip = false

// Set up vuetify.
Vue.use(Vuetify)

// Mount the App component onto the #app element.
new Vue({
  render: h => h(App)
}).$mount('#app')
