import { mdiPound, mdiSpeaker, mdiComment, mdiHeadSnowflake } from "@mdi/js";

class Rule {
  constructor(number, sounds, letters, comments) {
    this.number = number;
    this.sounds = sounds;
    this.letters = letters;
    this.comments = comments;
  }
}

export default {
  headers: [
    { text: "Number", icon: mdiPound },
    { text: "Sounds", icon: mdiSpeaker },
    { text: "Letters", icon: mdiComment },
    { text: "Comments/Memorization", icon: mdiHeadSnowflake },
  ],
  rules: [
    new Rule(
      "0",
      ["/s/,", "/z/"],
      ["<i>s</i>,", "soft <i>c</i>,", "<i>z</i>,", "<i>x</i> (in <i>xylophone</i>)"],
      [
        "<i>Zero</i> begins with <i>z</i>.",
        "Upper case <i>S</i> and <i>Z</i>, as well as lower case <i>s</i> and <i>z</i>, have <i>zero</i> vertical strokes each, as with the number <i>0</i>.",
        'The alveolar fricatives /s/ and /z/ form a <a href="https://en.wikipedia.org/wiki/Voice_(phonetics)" target="_blank">voiceless and voiced pair</a>.',
      ],
    ),
    new Rule(
      "1",
      ["/t/,", "/d/,", "(/θ/, /ð/)"],
      ["<i>t</i>,", "<i>d</i>,", "(<i>th</i> in <i>thing</i> and <i>this</i>)"],
      [
        "Upper case <i>T</i> and <i>D</i>, as well as lower case <i>t</i> and <i>d</i> have one vertical stroke each, as with the number <i>1</i>.",
        "The alveolar stops /t/ and /d/ form a voiceless and voiced pair, as do the similar-sounding dental fricatives /θ/ and /ð/, though some variant systems may omit the latter pair.",
      ],
    ),
    new Rule(
      "2",
      ["/n/"],
      ["<i>n</i>"],
      [
        'Upper case <i>N</i> and lower case <i>n</i> each have <i>two</i> vertical strokes and <i>two</i> points on the <a href="https://en.wikipedia.org/wiki/Baseline_(typography)" target="_blank">baseline</a>.',
      ],
    ),
    new Rule(
      "3",
      ["/m/"],
      ["<i>m</i>"],
      [
        "Lower case <i>m</i> has <i>three</i> vertical strokes.",
        "Both upper case <i>M</i> and lower case <i>m</i> each have <i>three</i> points on the baseline and look like the number <i>3</i> on its side.",
      ],
    ),
    new Rule(
      "4",
      ["/r/"],
      ["<i>r</i>,", "<i>l</i> (in <i>colonel</i>)"],
      ["<i>Four</i> ends with <i>r</i> (and /r/ in rhotic accents)."],
    ),
    new Rule(
      "5",
      ["/l/"],
      ["<i>l</i>"],
      [
        '<i>L</i> is the <a href="https://en.wikipedia.org/wiki/Roman_number" target="_blank">Roman number</a> for <i>5</i>0.',
        "Among the <i>five</i> digits of one's <i>l</i>eft hand, the thumb and index fingers also form an <i>L</i>",
      ],
    ),
    new Rule(
      "6",
      ["/tʃ/,", "/dʒ/,", "/ʃ/,", "/ʒ/"],
      [
        "<i>ch</i> (in <i>cheese</i> and <i>chef</i>),",
        "<i>j</i>,",
        "soft <i>g</i>,",
        "<i>sh</i>,",
        "<i>c</i> (in <i>cello</i> and <i>special</i>),",
        "<i>cz</i> (in <i>Czech</i>),",
        "<i>s</i> (in <i>tissue</i> and <i>vision</i>),",
        "<i>sc</i> (in <i>fascist</i>),",
        "<i>sch</i> (in <i>schwa</i> and <i>eschew</i>),",
        "<i>t</i> (in <i>ration</i> and <i>equation</i>),",
        "<i>tsch</i> (in <i>putsch</i>),",
        "<i>z</i> (in <i>seizure</i>)",
      ],
      [
        "Upper case <i>G</i> and lower case <i>g</i> look like the number <i>6</i> flipped horizontally and rotated 180° respectively.",
        'Lower case <a href="https://en.wikipedia.org/wiki/Script_typeface" target="_blank">script</a> <i>j</i> tends to have a lower loop, like the number <i>6</i>.',
        'In some <a href="https://en.wikipedia.org/wiki/Serif" target="_blank">serif</a> fonts, upper case <i>CH</i>, <i>SH</i> and <i>ZH</i> each have <i>six</i> serifs.',
        "The postalveolar affricates /tʃ/ and /dʒ/ form a voiceless and voiced pair, as do the similar-sounding postalveolar fricatives /ʃ/ and /ʒ/.",
        "<i>CH</i>urch has six letters.",
      ],
    ),
    new Rule(
      "7",
      ["/k/,", "/ɡ/"],
      ["<i>k</i>,", "hard <i>c</i>,", "<i>q</i>,", "<i>ch</i> (in loch),", "hard <i>g</i>"],
      [
        "Both upper case <i>K</i> and lower case <i>k</i> look like two small <i>7</i>s on their sides.",
        "In some fonts, the lower-right part of the upper case <i>G</i> looks like a <i>7</i>.",
        "<i>G</i> is also the <i>7</i>th letter of the alphabet.",
        "The velar stops /k/ and /g/ form a voiceless and voiced pair.",
      ],
    ),
    new Rule(
      "8",
      ["/f/,", "/v/"],
      ["<i>f</i>,", "<i>ph</i> (in <i>phone</i>),", "<i>v</i>,", "<i>gh</i> (in <i>laugh</i>)"],
      [
        "Lower case script <i>f</i>, which tends to have an upper and lower loop, looks like a <i>f</i>igure-<i>8</i>.",
        "The labiodental fricatives /f/ and /v/ form a voiceless and voiced pair.",
      ],
    ),
    new Rule(
      "9",
      ["/p/,", "/b/"],
      ["<i>p</i>,", "<i>b</i>,", "<i>gh</i> (in <i>hiccough</i>)"],
      [
        "Upper case <i>P</i> and lower case <i>p</i> look like the number <i>9</i> flipped horizontally.",
        "Lower case <i>b</i> looks like the number <i>9</i> turned 180°.",
        "The labial stops /p/ and /b/ form a voiceless and voiced pair.",
      ],
    ),
    new Rule(
      "Unassigned",
      [
        "/h/,",
        "/j/,",
        "/w/,",
        '<a href="https://en.wikipedia.org/wiki/English_phonology#Vowels" target="_blank">vowel sounds</a>',
      ],
      [
        "<i>h</i>,<i>y</i>,<i>w</i>,<i>a</i>,<i>e</i>,<i>i</i>,<i>o</i>,<i>u</i>,",
        '<a href="https://en.wikipedia.org/wiki/Silent_letter" target="_blank">silent letters</a>,',
        "<i>c</i> (in <i>packet</i> and <i>chutzpah</i>),",
        "<i>d</i> (in <i>judge</i>),",
        "<i>j</i> (in <i>Hallelujah</i> and <i>jalapeno</i>),",
        "<i>ll</i> (in <i>tortilla</i>),",
        "the first <i>p</i> in <i>sapphire</i>,",
        "<i>t</i> (in <i>match</i>),",
        "one of doubled letters in most contexts",
      ],
      [
        '<a href="https://en.wikipedia.org/wiki/Vowel" target="_blank">Vowel</a> sounds, <a href="https://en.wikipedia.org/wiki/Semivowel" target="_blank">semivowels</a> (/j/ and /w/) and /h/ do not correspond to any number. They can appear anywhere in a word without changing its number value.',
      ],
    ),
    new Rule(
      "(2, 27 or 7)",
      ["/ŋ/"],
      [
        "<i>ng</i>,",
        "<i>n</i> before <i>k</i>,",
        "hard <i>c</i>,",
        "<i>q</i>,",
        "hard <i>g</i> or <i>x</i>",
      ],
      [
        "Variant systems differ about whether /ŋ/ should encode <i>2</i> and classified together with /n/, <i>7</i> and classified together with /k/ and /g/ or even <i>27</i> (e.g. <i>ring</i> could be <i>42</i>, <i>47</i> or <i>427</i>). When a /k/ and /g/ is pronounced separately after the /ŋ/, variant systems that chose /ŋ/ to be <i>27</i> also disagree if an extra <i>7</i> should be written (e.g. finger could be <i>8274</i> or <i>82774</i>, or if /ŋ/ is chosen to be <i>7</i>, <i>8774</i>).",
      ],
    ),
  ],
};
