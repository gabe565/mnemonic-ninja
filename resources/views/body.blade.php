<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="description" content="">
        <meta name="author" content="">
        <meta name="csrf-token" content="{{ csrf_token() }}">

        <title>@yield('title') | {{ config('app.name') }}</title>

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

        <!-- Piwik -->
        <script type="text/javascript">
            var _paq = _paq || [];
            /* tracker methods like "setCustomDimension" should be called before "trackPageView" */
            _paq.push(["setDomains", ["*.mnemonic.gabe565.com"]]);
            _paq.push(['trackPageView']);
            _paq.push(['enableLinkTracking']);
            (function() {
                var u="//gabe565.com/analytics/";
                _paq.push(['setTrackerUrl', u+'piwik.php']);
                _paq.push(['setSiteId', '3']);
                var d=document, g=d.createElement('script'), s=d.getElementsByTagName('script')[0];
                g.type='text/javascript'; g.async=true; g.defer=true; g.src=u+'piwik.js'; s.parentNode.insertBefore(g,s);
            })();
        </script>
        <noscript><p><img src="//gabe565.com/analytics/piwik.php?idsite=3&rec=1" style="border:0;" alt="" /></p></noscript>
        <!-- End Piwik Code -->
    </head>
    <body>
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
                    <a class="navbar-brand" href="{{ route('Converters') }}">{{ config('app.name') }}</a>
                </div>

                <!-- Collect the nav links, forms, and other content for toggling -->
                <div class="collapse navbar-collapse navbar-collapse">
                    <ul class="nav navbar-nav">
                        <li class="{{ (Route::currentRouteNamed('Converters') == 'Converters') ? 'active' : '' }}"> 
                            <a href="{{ route('Converters') }}" class=""><i class="far fa-exchange fa-fw" aria-hidden="true"></i>&nbsp;Converters</a>
                        </li>
                        <li class="{{ (Route::currentRouteNamed('About') == 'About') ? 'active' : '' }}">
                            <a href="{{ route('About') }}"><i class="far fa-info-circle fa-fw" aria-hidden="true"></i>&nbsp;About</a>
                        </li>
                    </ul>
                </div>
                <!-- /.navbar-collapse -->
            </div>
            <!-- /.container -->
        </nav>

        <div class="container">
            @yield('content')
        </div>

        <div class="footer">
            <div class="container">
                <a href="https://github.com/gabe565/mnemonic-major-converter/blob/master/LICENSE" target="_blank" class="pull-left">
                    <span>&copy; 2017 Gabe Cook</span>
                </a>
                <a href="https://github.com/gabe565/mnemonic-major-converter" target="_blank" class="pull-right">
                    <span>
                        <i class="fab fa-github fa-fw"></i>&nbsp;View on GitHub
                    </span>
                </a>
            </div>
        </div>

        <script src="{{ mix('js/manifest.js') }}"></script>
        <script src="{{ mix('js/vendor.js') }}"></script>
        <script src="{{ mix('js/app.js') }}"></script>
    </body>
</html>
