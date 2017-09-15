@extends('body')

@section('title', 'About')

@section('content')
    <div class="container">
        <h1 id="About">About</h1>
        <hr>
        <div class="row">
            <div class="col-sm-12">
                <p>
                The <a href="https://en.wikipedia.org/wiki/Mnemonic_major_system" target="_blank">Mnemonic Major System</a> is a great way to easily remember numbers, but learning to remember which sounds and numbers correspond to each other takes time.<br>
                There are other online converters, but they all performed conversions by directly changing specific letters to a number. This is not how the system is meant to work, though. The system is meant to work with phonetic sounds. That is how this site does conversions. If a word or number is entered that is in its database (Which contains almost 126,000 words!), then it will perform the conversion according to how the word is pronounced.
                </p>
            </div>
        </div>
        <h2 id="Conversions">Conversions</h2>
        <div class="row">
            <div class="col-sm-12">
                <p>Based off of the <a href="https://en.wikipedia.org/wiki/Mnemonic_major_system#The_system" target="_blank">Wikipedia page</a>, here are the numbers and their corresponding sounds:</p>
                <table class="table datatable table-bordered table-condensed" style="width: 100%">
                    <thead>
                        <tr>
                            <th class="all">
                                <i class="far fa-hashtag fa-fw" aria-hidden="true"></i>&nbsp;
                                Number
                            </th>
                            <th class="all">
                                <i class="far fa-volume-up fa-fw" aria-hidden="true"></i>&nbsp;
                                Sound
                            </th>
                            <th class="all">
                                <i class="far fa-bars fa-fw" aria-hidden="true"></i>&nbsp;
                                Letters
                            </th>
                            <th class="min-desktop">
                                <i class="far fa-comments fa-fw" aria-hidden="true"></i>&nbsp;
                                Comments/Remembering
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td>0</td>
                            <td>/s/,<br>/z/</td>
                            <td><i>s</i>,<br>soft <i>c</i>,<br><i>z</i>,<br><i>x</i> (in <i>xylophone</i>)</td>
                            <td>
                                <ul>
                                    <li><i>Zero</i> begins with <i>z</i>.</li>
                                    <li>Upper case <i>S</i> and <i>Z</i>, as well as lower case <i>s</i> and <i>z</i>, have <i>zero</i> vertical strokes each, as with the number <i>0</i>.</li>
                                    <li>The alveolar fricatives /s/ and /z/ form a <a href="https://en.wikipedia.org/wiki/Voice_(phonetics)" title="Voice (phonetics)" target="_blank">voiceless and voiced pair</a>.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>1</td>
                            <td>/t/,<br>/d/,<br>(/θ/, /ð/)</td>
                            <td><i>t</i>,<br><i>d</i>,<br>(<i>th</i> in <i>thing</i> and <i>this</i>)</td>
                            <td>
                                <ul>
                                    <li>Upper case <i>T</i> and <i>D</i>, as well as lower case <i>t</i> and <i>d</i> have <i>one</i> vertical stroke each, as with the number <i>1</i>.</li>
                                    <li>The alveolar stops /t/ and /d/ form a voiceless and voiced pair, as do the similar-sounding dental fricatives /θ/ and /ð/, though some variant systems may omit the latter pair.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>2</td>
                            <td>/n/</td>
                            <td><i>n</i>
                            </td>
                            <td>
                                <ul>
                                    <li>Upper case <i>N</i> and lower case <i>n</i> each have <i>two</i> vertical strokes and <i>two</i> points on the <a href="https://en.wikipedia.org/wiki/Baseline_(typography)" title="Baseline (typography)" target="_blank">baseline</a>.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>3</td>
                            <td>/m/</td>
                            <td><i>m</i>
                            </td>
                            <td>
                                <ul>
                                    <li>Lower case <i>m</i> has <i>three</i> vertical strokes.</li>
                                    <li>Both upper case <i>M</i> and lower case <i>m</i> each have <i>three</i> points on the baseline and look like the number <i>3</i> on its side.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>4</td>
                            <td>/r/</td>
                            <td><i>r</i>,<br><i>l</i> (in <i>colonel</i>)</td>
                            <td>
                                <ul>
                                    <li><i>Four</i> ends with <i>r</i> (and /r/ in rhotic accents).</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>5</td>
                            <td>/l/</td>
                            <td><i>l</i>
                            </td>
                            <td>
                                <ul>
                                    <li><i>L</i> is the <a href="https://en.wikipedia.org/wiki/Roman_number" class="mw-redirect" title="Roman number" target="_blank">Roman number</a> for <i>5</i>0.</li>
                                    <li>Among the <i>five</i> digits of one's <i>l</i>eft hand, the thumb and index fingers also form an <i>L</i>.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>6</td>
                            <td>/tʃ/,<br>/dʒ/,<br>/ʃ/,<br>/ʒ/</td>
                            <td><i>ch</i> (in <i>cheese</i> and <i>chef</i>),<br><i>j</i>,<br>soft <i>g</i>,<br><i>sh</i>,<br><i>c</i> (in <i>cello</i> and <i>special</i>),<br><i>cz</i> (in <i>Czech</i>),<br><i>s</i> (in <i>tissue</i> and <i>vision</i>),<br><i>sc</i> (in <i>fascist</i>),<br><i>sch</i> (in <i>schwa</i> and <i>eschew</i>),<br><i>t</i> (in <i>ration</i> and <i>equation</i>),<br><i>tsch</i> (in <i>putsch</i>),<br><i>z</i> (in <i>seizure</i>)</td>
                            <td>
                                <ul>
                                    <li>Upper case <i>G</i> and lower case <i>g</i> look like the number <i>6</i> flipped horizontally and rotated 180° respectively.</li>
                                    <li>Lower case <a href="https://en.wikipedia.org/wiki/Script_typeface" title="Script typeface" target="_blank">script</a> <i>j</i> tends to have a lower loop, like the number <i>6</i>.</li>
                                    <li>In some <a href="https://en.wikipedia.org/wiki/Serif" title="Serif" target="_blank">serif</a> fonts, upper case <i>CH</i>, <i>SH</i> and <i>ZH</i> each have <i>six</i> serifs.</li>
                                    <li>The postalveolar affricates /tʃ/ and /dʒ/ form a voiceless and voiced pair, as do the similar-sounding postalveolar fricatives /ʃ/ and /ʒ/.</li>
                                    <li><i>CH</i>ur<i>ch</i> has six letters.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>7</td>
                            <td>/k/,<br>/ɡ/</td>
                            <td><i>k</i>,<br>hard <i>c</i>,<br><i>q</i>,<br><i>ch</i> (in <i>loch</i>),<br>hard <i>g</i>
                            </td>
                            <td>
                                <ul>
                                    <li>Both upper case <i>K</i> and lower case <i>k</i> look like two small <i>7</i>s on their sides.</li>
                                    <li>In some fonts, the lower-right part of the upper case <i>G</i> looks like a <i>7</i>.</li>
                                    <li><i>G</i> is also the 7th letter of the alphabet.</li>
                                    <li>The velar stops /k/ and /g/ form a voiceless and voiced pair.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>8</td>
                            <td>/f/,<br>/v/</td>
                            <td><i>f</i>,<br><i>ph</i> (in <i>phone</i>),<br><i>v,<br>gh (in laugh)</i>
                            </td>
                            <td>
                                <ul>
                                    <li>Lower case script <i>f</i>, which tends to have an upper and lower loop, looks like a <i>f</i>igure-<i>8</i>.</li>
                                    <li>The labiodental fricatives /f/ and /v/ form a voiceless and voiced pair.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>9</td>
                            <td>/p/,<br>/b/</td>
                            <td><i>p</i>,<br><i>b</i>,<br><i>gh</i> (in <i>hiccough</i>)</td>
                            <td>
                                <ul>
                                    <li>Upper case <i>P</i> and lower case <i>p</i> look like the number <i>9</i> flipped horizontally.</li>
                                    <li>Lower case <i>b</i> looks like the number <i>9</i> turned 180°.</li>
                                    <li>The labial stops /p/ and /b/ form a voiceless and voiced pair.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>Unassigned</td>
                            <td>/h/,<br>/j/,<br>/w/, <a href="https://en.wikipedia.org/wiki/English_phonology#Vowels" title="English phonology" target="_blank">vowel sounds</a>
                            </td>
                            <td><i>h</i>,<i>y</i>,<i>w</i>,<i>a</i>,<i>e</i>,<i>i</i>,<i>o</i>,<i>u</i>,<br><a href="https://en.wikipedia.org/wiki/Silent_letter" title="Silent letter" target="_blank">silent letters</a>,<br><i>c</i> (in <i>packet</i> and <i>chutzpah</i>),<br><i>d</i> (in <i>judge</i>),<br><i>j</i> (in <i>Hallelujah</i> and <i>jalapeno</i>),<br><i>ll</i> (in <i>tortilla</i>),<br>the first <i>p</i> in <i>sapphire</i>,<br><i>t</i> (in <i>match</i>),<br>one of doubled letters in most contexts</td>
                            <td>
                                <ul>
                                    <li><a href="https://en.wikipedia.org/wiki/Vowel" title="Vowel" target="_blank">Vowel</a> sounds, <a href="https://en.wikipedia.org/wiki/Semivowel" title="Semivowel" target="_blank">semivowels</a> (/j/ and /w/) and /h/ do not correspond to any number. They can appear anywhere in a word without changing its number value.</li>
                                </ul>
                            </td>
                        </tr>
                        <tr>
                            <td>(2, 27 or 7)</td>
                            <td>/ŋ/</td>
                            <td><i>ng</i>,<br><i>n</i> before <i>k</i>,<br>hard <i>c</i>,<br><i>q</i>,<br>hard <i>g</i> or <i>x</i>
                            </td>
                            <td>Variant systems differ about whether /ŋ/ should encode <i>2</i> and classified together with /n/, <i>7</i> and classified together with /k/ and /g/ or even <i>27</i> (e.g. <i>ring</i> could be <i>42</i>, <i>47</i> or <i>427</i>). When a /k/ and /g/ is pronounced separately after the /ŋ/, variant systems that chose /ŋ/ to be <i>27</i> also disagree if an extra <i>7</i> should be written (e.g. <i>finger</i> could be <i>8274</i> or <i>82774</i>, or if /ŋ/ is chosen to be <i>7</i>, <i>8774</i>).</td>
                        </tr>    
                    </tbody>
                </table>
            </div>
        </div>
        <h2 id="Development">Development</h2>
        <div class="row">
            <div class="col-sm-12">
                <hr>
                <p>
                To create this application, I started with the <a href="http://www.speech.cs.cmu.edu/cgi-bin/cmudict" target="_blank">CMU Dictionary</a> (I actually used <a href="https://github.com/cmusphinx/cmudict">this updated version on GitHub</a>).<br>
                Then I converted the <a href="https://en.wikipedia.org/wiki/Arpabet" target="_blank">Arpabet pronounciations</a> to the <a href="https://en.wikipedia.org/wiki/International_Phonetic_Alphabet" target="_blank">International Phonetic Alphabet (IPA)</a> using <a href="https://github.com/wwesantos/arpabet-to-ipa" target="_blank">this arpabet-to-ipa converter</a>.</p>
                <p>
                The results are stored in the database so that conversions are lighter (The server does not have to perform the conversion every time it is queried, instead it just looks up its saved conversion list) and faster (Database has indexes to speed up the lookups).
                </p>
                <p>
                When performing conversions, the application will attempt to perform the conversion from the pronounciation in the database.<br>
                If the word cannot be found, then it falls back to estimating the conversion by comparing letters instead of sounds.<br>
                To see this in action, <a href="/?word=garage">try converting "garage" to a number</a>. The result is "746" because the last /ge/ is not a hard G, but a soft G which corresponds to the number 6.<br>
                <a href="/?word=garages">If you change it to "garages"</a>, it will successfully be converted to "7460" because of the /s/ at the end of the word.<br>
                Now is where we can test the word list: <a href="/?word=garagez">if something like "garagez" is inserted</a>, it will convert to "7470" because /z/ does convert to "0", but the word "garagez" is not in the wordlist, so it guesses on conversions.
                </p>
                <p><a href="https://github.com/gabe565/mnemonic-major-converter" target="_blank">This project is now on GitHub</a></p>
            </div>
        </div>
    </div>
@endsection
