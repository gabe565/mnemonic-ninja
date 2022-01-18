import Vue from 'vue';
import VueRouter from 'vue-router';
import Converters from '@/views/Converters.vue';
import About from '@/views/About.vue';
import NotFound from '@/views/NotFound.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    redirect: { name: 'Converters', params: { startTab: 'interactive' } },
  },
  {
    path: '/converters',
    redirect: { name: 'Converters', params: { startTab: 'interactive' } },
  },
  {
    path: '/converters/:startTab(interactive|word|number)',
    alias: '/converters',
    name: 'Converters',
    component: Converters,
    props: true,
    meta: {
      icon: 'fa-exchange-alt',
      showInNav: true,
    },
  },
  {
    path: '/about',
    name: 'About',
    component: About,
    meta: {
      icon: 'fa-info-circle',
      showInNav: true,
    },
  },
  {
    path: '*',
    name: '404 Not Found',
    component: NotFound,
    meta: {
      icon: 'fa-exclamation-triangle',
    },
  },

  // Deprecated route redirects
  {
    path: '/converters/num/:query?',
    redirect: (to) => ({
      name: 'Converters',
      params: { startTab: 'number' },
      query: { q: to.params.query },
    }),
  },
  {
    path: '/converters/:startTab/:query',
    redirect: (to) => ({
      name: 'Converters',
      params: { startTab: to.params.startTab },
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
