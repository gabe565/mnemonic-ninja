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

        <link href="{{ mix('css/bootstrap.css') }}" rel="stylesheet">
        <link href="{{ mix('css/font-awesome.css') }}" rel="stylesheet">
        <link href="{{ mix('css/datatables.css') }}" rel="stylesheet">
        <link href="{{ mix('css/main.css') }}" rel="stylesheet">
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
            <nav class="navbar navbar-inverse navbar-fixed-top" role="navigation" id="app">
                <div class="container">
                    <div class="navbar-header page-scroll">
                        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                            <span class="sr-only">Toggle navigation</span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                            <span class="icon-bar"></span>
                        </button>
                        <router-link class="navbar-brand" to="/">
                            <img src="/images/mnemonic-ninja.svg" height="100%" class="align-top" style="display: inline-block" alt="{{ config('app.name') }} Logo">
                            &nbsp;&nbsp;{{ config('app.name') }}
                        </router-link>
                    </div>

                    <!-- Collect the nav links, forms, and other content for toggling -->
                    <div class="collapse navbar-collapse navbar-collapse">
                        <ul class="nav navbar-nav navbar-right">
                            <router-link tag="li" to="/convert">
                                <a>
                                    <svgicon name="exchange"></svgicon>
                                    &nbsp;Converters
                                </a>
                            </router-link>
                            <router-link tag="li" to="/about">
                                <a>
                                    <svgicon name="info-circle"></svgicon>
                                    &nbsp;About
                                </a>
                            </router-link>
                        </ul>
                    </div>
                    <!-- /.navbar-collapse -->
                </div>
                <!-- /.container -->
            </nav>

            <transition :name="transitionName" mode="out-in">
                <router-view class="child-view"></router-view>
            </transition>

            <div class="footer">
                <div class="container">
                    <a href="https://github.com/gabe565/mnemonic-major-converter/blob/master/LICENSE" target="_blank" class="pull-left">
                        <span>&copy; 2017 Gabe Cook</span>
                    </a>
                    <a href="https://github.com/gabe565/mnemonic-major-converter" target="_blank" class="pull-right">
                        <span>
                            <svgicon name="github"></svgicon>
                            &nbsp;View on GitHub
                        </span>
                    </a>
                </div>
            </div>

        </div>

        <script src="{{ mix('js/manifest.js') }}"></script>
        <script src="{{ mix('js/vendor.js') }}"></script>
        <script src="{{ mix('js/app.js') }}"></script>
    </body>
</html>
