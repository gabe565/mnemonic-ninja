import { nextTick } from "vue";
import { createRouter, createWebHistory } from "vue-router";
import HomeIcon from "~icons/material-symbols/home-rounded";
import InfoIcon from "~icons/material-symbols/info-rounded";
import SwapIcon from "~icons/material-symbols/swap-horizontal-circle-rounded";

const routes = [
  {
    path: "/",
    name: "Home",
    component: () => import("../views/HomePage.vue"),
    meta: {
      showInNav: true,
      icon: HomeIcon,
      group: "Info",
    },
  },
  {
    path: "/about",
    name: "About",
    component: () => import("../views/AboutPage.vue"),
    meta: {
      showInNav: true,
      icon: InfoIcon,
      group: "Info",
    },
  },
  {
    path: "/convert",
    redirect: "/convert/interactive",
  },
  {
    path: "/convert/interactive",
    name: "Interactive",
    component: () => import("../views/InteractiveConverter.vue"),
    props: (route) => route.query,
    meta: {
      showInNav: true,
      icon: SwapIcon,
      group: "Converters",
    },
  },
  {
    path: "/convert/number",
    name: "Number to Word",
    component: () => import("../views/NumberConverter.vue"),
    props: (route) => route.query,
    meta: {
      showInNav: true,
      icon: SwapIcon,
      group: "Converters",
      short: "Num to Word",
    },
  },
  {
    path: "/convert/word",
    name: "Word to Number",
    component: () => import("../views/WordConverter.vue"),
    props: (route) => route.query,
    meta: {
      showInNav: true,
      icon: SwapIcon,
      group: "Converters",
      short: "Word to Num",
    },
  },
  {
    path: "/:pathMatch(.*)*",
    name: "404 Not Found",
    component: () => import("../views/NotFoundPage.vue"),
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
