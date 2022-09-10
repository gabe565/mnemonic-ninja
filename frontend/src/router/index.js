import Vue from "vue";
import VueRouter from "vue-router";
import Converters from "../views/ConvertersPage.vue";
import About from "../views/AboutPage.vue";
import NotFound from "../views/NotFoundPage.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    redirect: "/convert/interactive",
  },
  {
    path: "/convert",
    redirect: "/convert/interactive",
  },
  {
    path: "/convert/interactive",
    name: "Interactive",
    component: Converters,
    props: { startTab: "interactive" },
  },
  {
    path: "/convert/number",
    name: "Number to Word",
    component: Converters,
    props: { startTab: "number" },
  },
  {
    path: "/convert/word",
    name: "Word to Number",
    component: Converters,
    props: { startTab: "word" },
  },
  {
    path: "/about",
    name: "About",
    component: About,
  },
  {
    path: "*",
    name: "404 Not Found",
    component: NotFound,
  },

  // Deprecated route redirects
  {
    path: "/converters(/.*)?",
    redirect: (to) => to.path.replace("/converters", "/convert"),
  },
  {
    path: "/convert/num(ber)?/:query?",
    redirect: (to) => ({
      path: "/convert/number",
      query: { q: to.params.query },
    }),
  },
  {
    path: "/convert/word/:query",
    redirect: (to) => ({
      path: "/convert/word",
      query: { q: to.params.query },
    }),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

const defaultTitle = document.title;
router.afterEach(async (to) => {
  // Use next tick to handle router history correctly
  // see: https://github.com/vuejs/vue-router/issues/914#issuecomment-384477609
  await Vue.nextTick();
  if (to.name) {
    document.title = `${to.name} - ${defaultTitle}`;
  } else {
    document.title = defaultTitle;
  }
});

export default router;
