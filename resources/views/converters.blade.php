@extends('body')

@section('title', 'Converters')

@section('content')
    <div class="container" id="vue">
        <h1>Converters</h1>
        <div class="row">
            <div class="col-md-8 col-sm-12 col-md-offset-2">
                <p>The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you. <a href="/about">Read more on the about page!</a></p>
                <p>Fill in a single value or a comma-separated list on the left and view the resulting conversions on the right.</p>
            </div>
        </div>
        <hr>
        <converter from="Word" to="Number" url="/api/to/num/" placeholder="example" help="Enter any word(s) to convert" value="{{ $word }}"></converter>
        <hr>
        <converter from="Number" to="Word" url="/api/to/word/" placeholder="70395" help="Enter any number(s) to convert" value="{{ $number }}"></converter>
    </div>
@endsection
