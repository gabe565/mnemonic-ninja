import { has } from 'lodash'
window._ = _

import jQuery from "jquery"
window.$ = window.jQuery = jQuery

import 'popper.js/dist/umd/popper'

import 'bootstrap'

/**
 * We'll load the axios HTTP library which allows us to easily issue requests
 * to our Laravel back-end. This library automatically handles sending the
 * CSRF token as a header based on the value of the "XSRF" token cookie.
 */

import axios from 'axios'
window.axios = axios

axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest'
