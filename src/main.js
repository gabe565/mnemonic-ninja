import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import "@fortawesome/fontawesome-pro/css/fontawesome.css";
import "@fortawesome/fontawesome-pro/css/solid.css";
import "@fortawesome/fontawesome-pro/css/brands.css";

createApp(App).use(router).use(vuetify).mount("#app");
