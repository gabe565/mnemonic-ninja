<template>
    <div class="container">
        <h1>Converters</h1>
        <hr>
        <div class="row">
            <div class="col-lg-10 col-md-12 mx-md-auto">
                <p>The mnemonic major system aids in numeric memorization by linking numbers with specific phonetic sounds, allowing you to convert a number to a word. This site will help do these conversions for you. <router-link to="/about">Read more on the about page!</router-link></p>
                <p>Fill in a single value or a comma-separated list on the left and view the resulting conversions on the right.</p>
            </div>
        </div>
        <hr>
        <converter :from="{ label: 'Word', placeholder: 'example', value: word, regex: /^[A-Za-z-\'\s,;\.]*$/g }" :to="{ label: 'Number' }" description="Enter a word or a list of words to get the number conversions.<br>Note that more than one number may show up from a single word. This means that there is more than one pronounciation!" url="/api/to/num/"></converter>
        <hr>
        <converter :from="{ label: 'Number', placeholder: 70395, value: num, regex: /^[0-9\s,;]*$/g }" :to="{ label: 'Word' }" description="Enter a number or a list of numbers to get all associated words.<br>Note that if many words show up for a single number, you may have to scroll through the results!" url="/api/to/word/"></converter>
    </div>
</template>

<script>
import Converter from './Converter.vue'

export default {
    computed: {
        word() {
            if (this.$route.params.hasOwnProperty('word'))
                return this.$route.params.word
            else
                return ''
        },
        num() {
            if (this.$route.params.hasOwnProperty('num'))
                return this.$route.params.num
            else
                return ''
        }
    },
    components: {
        converter: Converter
    }
}
</script>
