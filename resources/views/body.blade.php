<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <meta name="description" content="The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you.">
        <meta name="author" content="Gabe Cook">
        <meta name="keywords" content="mnemonic major system, convert, memorize numbers, phonetic, memorization technique, mnemonic major converter">

        <meta property='og:title' content='Mnemonic Major Converter'/>
        <meta property='og:image' content='{{ config('app.url') }}/images/thumb.png'/>
        <meta property='og:description' content='The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you.'/>
        <meta property='og:url' content='{{ config('app.url') }}' />
        <meta property="og:type" content="website" />

        <meta name="csrf-token" content="{{ csrf_token() }}">

        <title>{{ config('app.name') }}</title>

        <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png?v=zXrE0PzqQg">
        <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png?v=zXrE0PzqQg">
        <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png?v=zXrE0PzqQg">
        <link rel="manifest" href="/manifest.json?v=zXrE0PzqQg">
        <link rel="mask-icon" href="/safari-pinned-tab.svg?v=zXrE0PzqQg" color="#607d8b">
        <link rel="shortcut icon" href="/favicon.ico?v=zXrE0PzqQg">
        <meta name="theme-color" content="#607d8b">

        <link href="{{ mix('css/app.css') }}" rel="stylesheet">
        <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
        <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
        <!--[if lt IE 9]>
            <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
            <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
        <![endif]-->

        <!-- Global Site Tag (gtag.js) - Google Analytics -->
        <script async src="https://www.googletagmanager.com/gtag/js?id=UA-107254206-1"></script>
        <script>
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments)};
gtag('js', new Date());

gtag('config', 'UA-107254206-1');
        </script>
    </head>
    <body>
        <div id="app">

            <!-- Navigation -->
            <nav class="navbar navbar-expand-md navbar-dark bg-secondary fixed-top">
                <div class="container">
                    <router-link to="/" class="navbar-brand">
                        <img src="/images/mnemonic-ninja.svg" class="d-inline-block" alt="{{ config('app.name') }} Logo" style="height: 55px">
                        &nbsp;{{ config('app.name') }}
                    </router-link>
                    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                        <span class="navbar-toggler-icon"></span>
                    </button>
                    <div class="collapse navbar-collapse" id="navbarResponsive">
                        <ul class="navbar-nav ml-auto">
                            <li class="nav-item">
                                <router-link to="/convert" class="nav-link">
                                    <svgicon name="exchange"></svgicon>
                                    &nbsp;Converters
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/about" class="nav-link">
                                    <svgicon name="info-circle"></svgicon>
                                    &nbsp;About
                                </router-link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <transition :name="transitionName" mode="out-in">
                <router-view class="child-view"></router-view>
            </transition>

            <footer class="footer bg-secondary">
                <div class="container">
                    <a href="https://github.com/gabe565/mnemonic-major-converter/blob/master/LICENSE" target="_blank" class="float-left">
                        &copy; 2017 Gabe Cook
                    </a>
                    <a href="https://github.com/gabe565/mnemonic-major-converter" target="_blank" class="float-right">
                        <svgicon name="github"></svgicon>
                        &nbsp;View on GitHub
                    </a>
                </div>
            </footer>

        </div>

        <script src="https://cdn.polyfill.io/v2/polyfill.js?unknown=polyfill&features=Object.assign|gated,Promise|gated"></script>
        <script src="{{ mix('js/manifest.js') }}"></script>
        <script src="{{ mix('js/vendor.js') }}"></script>
        <script src="{{ mix('js/app.js') }}"></script>
    </body>
</html>
