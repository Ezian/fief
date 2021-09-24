import './plugins/fontawesome'
import { createApp } from 'vue'
import App from './App.vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

createApp(App)
  .component('font-awesom-icon', FontAwesomeIcon)
  .mount('#app')
