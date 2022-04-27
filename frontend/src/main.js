import Vue from 'vue';
import '@/registerServiceWorker';
import router from '@/router';
import vuetify from '@/plugins/vuetify';
import App from './App.vue';
import '@fortawesome/fontawesome-pro/css/fontawesome.css';
import '@fortawesome/fontawesome-pro/css/solid.css';
import '@fortawesome/fontawesome-pro/css/brands.css';

Vue.config.productionTip = false;

new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
