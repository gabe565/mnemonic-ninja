<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <meta name="description" content="The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you.">
        <meta name="author" content="Gabe Cook">
        <meta name="keywords" content="mnemonic major system, convert, memorize numbers, phonetic, memorization technique">

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
                    <ul class="nav navbar-nav navbar-right">
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
