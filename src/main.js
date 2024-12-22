import { createApp } from "vue";
import App from "./App.vue";
import "./plugins/plausible";
import router from "./plugins/router";
import vuetify from "./plugins/vuetify";

createApp(App).use(router).use(vuetify).mount("#app");
