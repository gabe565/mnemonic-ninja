import Vue from 'vue';
import VueRouter from 'vue-router';
import Converters from '@/views/Converters.vue';
import About from '@/views/About.vue';
import NotFound from '@/views/NotFound.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    redirect: '/converters/interactive',
  },
  {
    path: '/converters',
    redirect: '/converters/interactive',
  },
  {
    path: '/converters/interactive',
    name: 'Interactive',
    component: Converters,
    props: { startTab: 'interactive' },
  },
  {
    path: '/converters/number',
    name: 'Number to Word',
    component: Converters,
    props: { startTab: 'number' },
  },
  {
    path: '/converters/word',
    name: 'Word to Number',
    component: Converters,
    props: { startTab: 'word' },
  },
  {
    path: '/about',
    name: 'About',
    component: About,
  },
  {
    path: '*',
    name: '404 Not Found',
    component: NotFound,
  },

  // Deprecated route redirects
  {
    path: '/converters/num(ber)?/:query?',
    redirect: (to) => ({
      path: '/converters/number',
      query: { q: to.params.query },
    }),
  },
  {
    path: '/converters/word/:query',
    redirect: (to) => ({
      path: '/converters/word',
      query: { q: to.params.query },
    }),
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
