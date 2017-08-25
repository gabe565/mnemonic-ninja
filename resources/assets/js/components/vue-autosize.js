window.autosize = require('autosize')

exports.install = function(Vue) {
    Vue.directive('autosize', {
        bind: function(el, binding) {
            var tagName = el.tagName
            autosize(el)
        },

        componentUpdated: function(el, binding, vnode) {
            var tagName = el.tagName
            autosize.update(el)
        },

        unbind: function(el) {
            autosize.destroy(el)
        }
    })
}
