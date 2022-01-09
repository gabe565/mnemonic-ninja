import Vue from 'vue';
import VueRouter from 'vue-router';
import Converters, { Tabs } from '@/views/Converters.vue';
import About from '@/views/About.vue';
import NotFound from '@/views/NotFound.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/converters',
    name: 'Converters',
    component: Converters,
    props: true,
    meta: {
      icon: 'fa-exchange-alt',
      showInNav: true,
    },
  },
  {
    path: '/convert/word/:word',
    redirect: ({ params }) => ({ name: 'Converters', params: { word: params.word, startTab: Tabs.WordToNumber } }),
  },
  {
    path: '/convert/num/:num',
    redirect: ({ params }) => ({ name: 'Converters', params: { num: params.num } }),
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
    redirect: { name: 'Converters' },
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
