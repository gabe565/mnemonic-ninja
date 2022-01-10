import Vue from 'vue';
import VueRouter from 'vue-router';
import Converters from '@/views/Converters.vue';
import About from '@/views/About.vue';
import NotFound from '@/views/NotFound.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/converters',
    redirect: { name: 'Converters', params: { startTab: 'interactive' } },
  },
  {
    path: '/converters/:startTab?/:query?',
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
    path: '/',
    redirect: { name: 'Converters', params: { startTab: 'interactive' } },
  },
  {
    path: '*',
    name: '404 Not Found',
    component: NotFound,
    meta: {
      icon: 'fa-exclamation-triangle',
    },
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
