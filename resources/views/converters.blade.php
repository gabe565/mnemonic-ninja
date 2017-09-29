@extends('body')

@section('title', 'Converters')

@section('content')
    <div class="container" id="vue">
        <h1>Converters</h1>
        <div class="row">
            <div class="col-md-10 col-sm-12 col-md-offset-1">
                <p>The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you. <a href="/about">Read more on the about page!</a></p>
                <p>Fill in a single value or a comma-separated list on the left and view the resulting conversions on the right.</p>
            </div>
        </div>
        <hr>
        <converter :from="{ label: 'Word', placeholder: 'example', help: 'Enter any word(s) to convert', value: '{{ $word }}' }" :to="{ label: 'Number', help: 'Any corresponding numbers will be shown here.' }" url="/api/to/num/"></converter>
        <hr>
        <converter :from="{ label: 'Number', placeholder: 70395, help: 'Enter any number(s) to convert', value: '{{ $number }}' }" :to="{ label: 'Word', help: 'All associated words will be shown here. Note that you may need to scroll through to see all results.' }" url="/api/to/word/"></converter>
    </div>
@endsection
