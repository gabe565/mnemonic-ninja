<template>
    <div class="converter">
        <div class="row">
            <div class="col-lg-10 col-md-12 mx-md-auto">
                <p v-html="description"></p>
            </div>
        </div>
        <div class="row equal text-center">
            <div class="col-md-4 ml-md-auto my-auto" :class="{ 'has-error': error }">
                <label :for="_uid" class="h4">{{ from.label }}</label>
                <textarea :id="_uid" class="form-control conversion" :class="{ 'is-invalid': !valid }" name="query" :placeholder="from.placeholder" v-model="query"></textarea>
            </div>
            <div class="col-md-2 my-auto">
                <button type="submit" class="btn btn-success btn-sm" v-on:click="manualUpdate">
                    <svgicon name="arrow-right" class="svg-fh svg-2x d-none d-md-inline-block" v-if="!loading"></svgicon>
                    <svgicon name="arrow-down" class="svg-fh svg-2x d-inline-block d-md-none" v-if="!loading"></svgicon>
                    <svgicon name="sync-alt" class="svg-fh svg-2x svg-spin" v-if="loading"></svgicon>
                </button>
            </div>
            <div class="col-md-4 mr-md-auto my-auto">
                <label class="h4">{{ to.label }}</label>
                <div class="form-control conversion">
                    <table class="table table-striped text-left">
                        <tbody>
                            <tr v-for="item in result">
                                <td class="min">{{ item.q }}:</td>
                                <td class="result">
                                    <div>
                                        {{ item.r }}
                                    </div>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import '../svg/arrow-right'
import '../svg/arrow-down'
import '../svg/sync-alt'

export default {
    data: function() {
        return {
            query: '',
            response: [],
            result: [],
            loading: false,
            error: false
        };
    },
    props: {
        'from': {
            type: Object
        },
        'to': {
            type: Object
        },
        'url': {
            type: String
        },
        'description': {
            type: String
        }
    },
    watch: {
        query: _.debounce(
            function() {
                this.getResponse()
            }, 300)
    },
    computed: {
        valid: function() {
            return this.query.match(this.from.regex)
        },
        empty: function() {
            return this.query == ''
        }
    },
    methods: {
        manualUpdate: _.throttle(
            function() {
                this.getResponse()
            }, 1000),
        getResponse: function () {
            if (this.empty || !this.valid) {
                this.response = []
                this.result = {}
                return
            }
            this.loading = true
            var vue = this
            axios.get(this.url + encodeURIComponent(this.query))
                .then(function (response) {
                    vue.error = false
                    vue.response = response.data
                    vue.createResult()
                    vue.loading = false
                })
                .catch(function (response) {
                    vue.error = true
                    vue.loading = false
                })
        },
        createResult: function() {
            this.result = []
            var vue = this
            _.forEach(vue.response.result, function(value) {
                vue.result.push({
                    q: value.q,
                    r: value.r.join(', ')
                })
            })
        }
    },
    created: function() {
        if (this.from.value) {
            this.query = this.from.value
            this.getResponse()
        }
    },
}
</script>
