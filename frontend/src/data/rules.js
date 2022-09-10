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
    { text: "Number", icon: "fa-hashtag" },
    { text: "Sounds", icon: "fa-lips" },
    { text: "Letters", icon: "fa-comment" },
    { text: "Comments/Memorization", icon: "fa-head-side-brain" },
  ],
  rules: [
    new Rule(
      "0",
      ["/s/,", "/z/"],
      ["__s__,", "soft __c__,", "__z__,", "__x__ (in __xylophone__)"],
      [
        "__Zero__ begins with __z__.",
        "Upper case __S__ and __Z__, as well as lower case __s__ and __z__, have __zero__ vertical strokes each, as with the number __0__.",
        "The alveolar fricatives /s/ and /z/ form a [voiceless and voiced pair](https://en.wikipedia.org/wiki/Voice_(phonetics)).",
      ]
    ),
    new Rule(
      "1",
      ["/t/,", "/d/,", "(/θ/, /ð/)"],
      ["__t__,", "__d__,", "(__th__ in __thing__ and __this__)"],
      [
        "Upper case __T__ and __D__, as well as lower case __t__ and __d__ have one vertical stroke each, as with the number __1__.",
        "The alveolar stops /t/ and /d/ form a voiceless and voiced pair, as do the similar-sounding dental fricatives /θ/ and /ð/, though some variant systems may omit the latter pair.",
      ]
    ),
    new Rule(
      "2",
      ["/n/"],
      ["__n__"],
      [
        "Upper case __N__ and lower case __n__ each have __two__ vertical strokes and __two__ points on the [baseline](https://en.wikipedia.org/wiki/Baseline_(typography)).",
      ]
    ),
    new Rule(
      "3",
      ["/m/"],
      ["__m__"],
      [
        "Lower case __m__ has __three__ vertical strokes.",
        "Both upper case __M__ and lower case __m__ each have __three__ points on the baseline and look like the number __3__ on its side.",
      ]
    ),
    new Rule(
      "4",
      ["/r/"],
      ["__r__,", "__l__ (in __colonel__)"],
      ["__Four__ ends with __r__ (and /r/ in rhotic accents)."]
    ),
    new Rule(
      "5",
      ["/l/"],
      ["__l__"],
      [
        "__L__ is the [Roman number](https://en.wikipedia.org/wiki/Roman_number) for __5__0.",
        "Among the __five__ digits of one's __l__eft hand, the thumb and index fingers also form an __L__",
      ]
    ),
    new Rule(
      "6",
      ["/tʃ/,", "/dʒ/,", "/ʃ/,", "/ʒ/"],
      [
        "__ch__ (in __cheese__ and __chef__),",
        "__j__,",
        "soft __g__,",
        "__sh__,",
        "__c__ (in __cello__ and __special__),",
        "__cz__ (in __Czech__),",
        "__s__ (in __tissue__ and __vision__),",
        "__sc__ (in __fascist__),",
        "__sch__ (in __schwa__ and __eschew__),",
        "__t__ (in __ration__ and __equation__),",
        "__tsch__ (in __putsch__),",
        "__z__ (in __seizure__)",
      ],
      [
        "Upper case __G__ and lower case __g__ look like the number __6__ flipped horizontally and rotated 180° respectively.",
        'Lower case <a href="https://en.wikipedia.org/wiki/Script_typeface" target="_blank">script</a> __j__ tends to have a lower loop, like the number __6__.',
        'In some <a href="https://en.wikipedia.org/wiki/Serif" target="_blank">serif</a> fonts, upper case __CH__, __SH__ and __ZH__ each have __six__ serifs.',
        "The postalveolar affricates /tʃ/ and /dʒ/ form a voiceless and voiced pair, as do the similar-sounding postalveolar fricatives /ʃ/ and /ʒ/.",
        "__CH__urch has six letters.",
      ]
    ),
    new Rule(
      "7",
      ["/k/,", "/ɡ/"],
      ["__k__,", "hard __c__,", "__q__,", "__ch__ (in loch),", "hard __g__"],
      [
        "Both upper case __K__ and lower case __k__ look like two small __7__s on their sides.",
        "In some fonts, the lower-right part of the upper case __G__ looks like a __7__.",
        "__G__ is also the __7__th letter of the alphabet.",
        "The velar stops /k/ and /g/ form a voiceless and voiced pair.",
      ]
    ),
    new Rule(
      "8",
      ["/f/,", "/v/"],
      ["__f__,", "__ph__ (in __phone__),", "__v__,", "__gh__ (in __laugh__)"],
      [
        "Lower case script __f__, which tends to have an upper and lower loop, looks like a __f__igure-__8__.",
        "The labiodental fricatives /f/ and /v/ form a voiceless and voiced pair.",
      ]
    ),
    new Rule(
      "9",
      ["/p/,", "/b/"],
      ["__p__,", "__b__,", "__gh__ (in __hiccough__)"],
      [
        "Upper case __P__ and lower case __p__ look like the number __9__ flipped horizontally.",
        "Lower case __b__ looks like the number __9__ turned 180°.",
        "The labial stops /p/ and /b/ form a voiceless and voiced pair.",
      ]
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
        "__h__,__y__,__w__,__a__,__e__,__i__,__o__,__u__,",
        '<a href="https://en.wikipedia.org/wiki/Silent_letter" target="_blank">silent letters</a>,',
        "__c__ (in __packet__ and __chutzpah__),",
        "__d__ (in __judge__),",
        "__j__ (in __Hallelujah__ and __jalapeno__),",
        "__ll__ (in __tortilla__),",
        "the first __p__ in __sapphire__,",
        "__t__ (in __match__),",
        "one of doubled letters in most contexts",
      ],
      [
        '<a href="https://en.wikipedia.org/wiki/Vowel" target="_blank">Vowel</a> sounds, <a href="https://en.wikipedia.org/wiki/Semivowel" target="_blank">semivowels</a> (/j/ and /w/) and /h/ do not correspond to any number. They can appear anywhere in a word without changing its number value.',
      ]
    ),
    new Rule(
      "(2, 27 or 7)",
      ["/ŋ/"],
      ["__ng__,", "__n__ before __k__,", "hard __c__,", "__q__,", "hard __g__ or __x__"],
      [
        "Variant systems differ about whether /ŋ/ should encode __2__ and classified together with /n/, __7__ and classified together with /k/ and /g/ or even __27__ (e.g. __ring__ could be __42__, __47__ or __427__). When a /k/ and /g/ is pronounced separately after the /ŋ/, variant systems that chose /ŋ/ to be __27__ also disagree if an extra __7__ should be written (e.g. finger could be __8274__ or __82774__, or if /ŋ/ is chosen to be __7__, __8774__).",
      ]
    ),
  ],
};
