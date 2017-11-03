<template>
    <div class="converter">
        <div class="row">
            <div class="col-md-10 col-sm-12 col-md-offset-1">
                <p v-html="description"></p>
            </div>
        </div>
        <div class="row equal">
            <div class="col-sm-5 col-md-4 col-md-offset-1 centered" :class="{ 'has-error': error }">
                <label :for="_uid">{{ from.label }}</label>
                <textarea :id="_uid" class="form-control conversion" name="query" :placeholder="from.placeholder" v-model="query"></textarea>
            </div>
            <div class="col-sm-2 centered">
                <button type="submit" class="btn btn-success btn-sm" v-on:click="manualUpdate">
                    <svgicon name="arrow-right" class="svg-fh svg-fw svg-2x hidden-xs" v-if="!loading"></svgicon>
                    <svgicon name="arrow-down" class="svg-fh svg-fw svg-2x visible-xs-inline-block" v-if="!loading"></svgicon>
                    <svgicon name="sync-alt" class="svg-fh svg-fw svg-2x svg-spin" v-if="loading"></svgicon>
                </button>
            </div>
            <div class="col-sm-5 col-md-4 centered">
                <h4>{{ to.label }}</h4>
                <div class="form-control conversion">
                    <table class="table table-striped">
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
    methods: {
        manualUpdate: _.throttle(
            function() {
                this.getResponse()
            }, 1000),
        getResponse: function () {
            if (this.query == '') {
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
