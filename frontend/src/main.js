import Vue from 'vue';
import App from './App.vue';
import '@/registerServiceWorker';
import router from '@/router';
import vuetify from '@/plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css';
import 'roboto-fontface/css/roboto-slab/roboto-slab-fontface.css';
import '@fortawesome/fontawesome-pro/css/all.css';

Vue.config.productionTip = false;

new Vue({
  router,
  vuetify,
  render: (h) => h(App),
}).$mount('#app');
