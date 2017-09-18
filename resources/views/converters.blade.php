@extends('body')

@section('title', 'Converters')

@section('content')
    <div class="container" id="vue">
        <h1>Converters</h1>
        <hr>
        <converter from="Words" to="Number" url="/api/to/num/" placeholder="example" help="Enter any word to convert" value="{{ $word }}"></converter>
        <hr>
        <converter from="Numbers" to="Word" url="/api/to/word/" placeholder="70395" help="Enter any number to convert" value="{{ $number }}"></converter>
    </div>
@endsection
