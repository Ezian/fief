import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import 'bulma/css/bulma.css'
import { FontAwesomeIcon } from './plugins/font-awesome'
import axios from 'axios'


axios.defaults.headers.common = {
  "Content-Type": "application/json"
}

createApp(App).use(router)
  .component("font-awesome-icon", FontAwesomeIcon)
  .mount('#app')
