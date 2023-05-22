import { createRouter, createWebHistory } from "vue-router";
import Converters from "../views/ConvertersPage.vue";
import About from "../views/AboutPage.vue";
import NotFound from "../views/NotFoundPage.vue";

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
    path: "/convert/:startTab(interactive|word|number)",
    name: "Convert",
    component: Converters,
    props: true,
  },
  {
    path: "/about",
    name: "About",
    component: About,
  },
  {
    path: "/:pathMatch(.*)*",
    name: "404 Not Found",
    component: NotFound,
  },

  // Deprecated route redirects
  {
    path: "/converters:pathMatch(/?.*)*",
    redirect: (to) => to.path.replace("/converters", "/convert"),
  },
  {
    path: "/convert/num:pathMatch(ber)?/:query",
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

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

const defaultTitle = document.title;
router.afterEach(async (to) => {
  // Use next tick to handle router history correctly
  // see: https://github.com/vuejs/vue-router/issues/914#issuecomment-384477609
  await nextTick();
  if (to.name) {
    document.title = `${to.name} - ${defaultTitle}`;
  } else {
    document.title = defaultTitle;
  }
});

export default router;
