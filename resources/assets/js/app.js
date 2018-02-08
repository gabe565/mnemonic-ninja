/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

import 'vue-svgicon/dist/polyfill'

require('./bootstrap')
require('./hide-nav')

import 'datatables.net'
import 'datatables.net-bs4'
import 'datatables.net-fixedheader'
import 'datatables.net-responsive'

/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

import Vue from 'vue'
import VueRouter from 'vue-router'
import VueSVGIcon from 'vue-svgicon'

Vue.use(VueRouter)
Vue.use(VueSVGIcon)

import Converters from './components/Converters.vue'
import About from './components/About.vue'
import NotFound from './components/NotFound.vue'

import './svg/bars'
import './svg/exchange'
import './svg/info-circle'
import './svg/github'

const routes = [
    { path: '/', redirect: '/convert' },
    { path: '/convert', component: Converters, meta: { title: 'Converters' } },
    { path: '/convert/word/:word', component: Converters, meta: { title: 'Converters' } },
    { path: '/convert/num/:num', component: Converters, meta: { title: 'Converters' } },
    { path: '/about', component: About, meta: { title: 'About' } },
    { path: '*', component: NotFound, meta: { title: 'Not Found' } }
]

const router = new VueRouter({
    routes,
    mode: 'history',
    linkActiveClass: 'active',
    scrollBehavior(to, from, savedPosition) {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                if (savedPosition)
                    resolve(savedPosition)
                else
                    resolve({ x: 0, y: 0 })
            }, 250)
        })
    }
})

router.beforeEach((to, from, next) => {
    document.title = to.meta.title ? to.meta.title + ' | Mnemonic Ninja' : 'Mnemonic Ninja'
    $('#app').removeClass(from.meta.title).addClass(to.meta.title)
    next()
})

new Vue({
    el: '#app',
    router,
    data() {
        return {
            transitionName: 'fade'
        }
    },
    watch: {
        '$route'(to, from) {
            var fromIndex = routes.findIndex(obj => obj.path == from.path)
            var toIndex = routes.findIndex(obj => obj.path == to.path)
            this.transitionName = (fromIndex < toIndex) ? 'slide-left' : 'slide-right'
        }
    }
})

require('./collapse-on-nav')
