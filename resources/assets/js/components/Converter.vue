<template>
    <div class="row equal">
        <div class="col-sm-5 col-md-4 col-md-offset-1 centered" :class="{ 'has-error': error }">
            <label :for="_uid">{{ from.label }}</label>
            <textarea :id="_uid" class="form-control conversion" name="query" :placeholder="from.placeholder" autocomplete="off" aria-describedby="help" v-model="query"></textarea>
            <span class="help-block">{{ from.help }}</span>
        </div>
        <div class="col-sm-2 centered">
            <button type="submit" class="btn btn-success btn-sm">
                <i class="far fa-2x fa-fw loader" :class="loader" aria-hidden="true"></i>
            </button>
        </div>
        <div class="col-sm-5 col-md-4 centered">
            <h4>{{ to.label }}</h4>
            <div class="form-control conversion">
                <table class="table table-striped">
                    <tbody>
                        <tr v-for="(item, key) in result">
                            <td class="min">{{ key }}:</td>
                            <td class="result">
                                <div>
                                    {{ item }}
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <span class="help-block">{{ to.help }}</span>
        </div>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            query: '',
            response: [],
            result: {},
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
        }
    },
    computed: {
        loader: function() {
            if (this.loading) {
                return ['fa-sync-alt', 'fa-spin']
            } else {
                return ['arrow']
            }
        },
        queries: function() {
            return this.query.split(' ');
        }
    },
    watch: {
        query: _.debounce(
            function() {
                this.getResponse()
            }, 250)
    },
    methods: {
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
            this.result = {}
            var vue = this
            _.forEach(vue.response.result, function(value, key) {
                var it = '';
                if (typeof value == 'string')
                    it += value
                else
                    it += value.join(', ')

                vue.result[key] = it
            })
        }
    },
    created: function() {
        if (this.from.value) {
            this.query = this.from.value
            this.getResponse()
        }
    }
}
</script>
