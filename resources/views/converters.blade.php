@extends('body')

@section('title', 'Converters')

@section('content')
    <div class="container" id="vue">
        <h1>Converters</h1>
        <hr>
        <div class="row">
            <div class="col-md-10 col-sm-12 col-md-offset-1">
                <p>The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you. <a href="/about">Read more on the about page!</a></p>
                <p>Fill in a single value or a comma-separated list on the left and view the resulting conversions on the right.</p>
            </div>
        </div>
        <hr>
        <converter :from="{ label: 'Word', placeholder: 'example', value: '{{ $word }}' }" :to="{ label: 'Number' }" description="Enter a word or a list of words to get the number conversions.<br>Note that more than one number may show up from a single word. This means that there is more than one pronounciation!" url="/api/to/num/"></converter>
        <hr>
        <converter :from="{ label: 'Number', placeholder: 70395, value: '{{ $number }}' }" :to="{ label: 'Word' }" description="Enter a number or a list of numbers to get all associated words.<br>Note that if many words show up for a single number, you may have to scroll through the results!" url="/api/to/word/"></converter>
    </div>
@endsection
