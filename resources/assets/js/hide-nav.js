import { throttle } from 'lodash'

$(function(){
    "use strict"

    var prev = 50
    var $window = $(window)
    var $nav = $('nav')

    $window.on('scroll', _.throttle(function(){
        var scrollTop = $window.scrollTop()
        $nav.toggleClass('navbar-hide', scrollTop > prev)
        prev = scrollTop > 50 ? scrollTop : 50
    }, 250, { leading: true, trailing: true }))
})
