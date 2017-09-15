@extends('body')

@section('title', 'Converters')

@section('content')
    <div class="container" id="vue">
        <h1>Converters</h1>
        <hr>
        <converter from="Word(s)" to="Number" url="/api/to/num/" placeholder="example" help="Enter any word to convert" joinstr=", " value="{{ $word }}"></converter>
        <hr>
        <converter from="Number" to="Word" url="/api/to/word/" placeholder="70395" help="Enter any number to convert" joinstr=",  " value="{{ $number }}"></converter>
    </div>
@endsection
