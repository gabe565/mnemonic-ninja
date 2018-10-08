<template>
    <div class="converter">
        <div class="row">
            <div class="col-lg-10 col-md-12 mx-md-auto">
                <p v-html="description"></p>
            </div>
        </div>
        <div class="row equal text-center">
            <div class="col-md-4 ml-md-auto my-md-auto mb-4" :class="{ 'has-error': error }">
                <label :for="_uid" class="h4">{{ from.label }}</label>
                <textarea :id="_uid" class="form-control conversion" :class="{ 'is-invalid': !valid }" name="query" :placeholder="from.placeholder" v-model="query"></textarea>
            </div>
            <div class="col-md-2 my-md-auto mb-4">
                <button type="submit" class="btn btn-success btn-sm" v-on:click="manualUpdate" :disabled="disabled">
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
    data() {
        return {
            query: '',
            response: {},
            loading: false,
            error: false,
            disabled: true
        }
    },
    props: {
        from: {
            type: Object
        },
        to: {
            type: Object
        },
        url: {
            type: String
        },
        description: {
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
        valid() {
            return this.query.match(this.from.regex)
        },
        empty() {
            return this.query == ''
        },
        result() {
            if (this.response.result) {
                return this.response.result.map(value => ({
                    q: value.q,
                    r: value.r.join(', ')
                }))
            } else
                return {}
        }
    },
    methods: {
        manualUpdate() {
            this.disabled = true
            setTimeout(() => { this.disabled = false }, 750)
            this.getResponse()
        },
        getResponse() {
            if (this.empty || !this.valid) {
                this.response = {}
                this.disabled = true
                return
            }
            this.loading = true
            axios.get(this.url + encodeURIComponent(this.query))
                .then(response => {
                    this.error = false
                    this.response = response.data
                    this.loading = false
                    this.disabled = false
                })
                .catch(response => {
                    this.error = true
                    this.loading = false
                    this.disabled = false
                })
        }
    },
    created() {
        if (this.from.value) {
            this.query = this.from.value
            this.getResponse()
        }
    }
}
</script>
