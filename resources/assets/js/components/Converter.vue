<template>
    <div class="row equal">
        <div class="col-sm-3 col-sm-offset-2 centered">
            <label :for="_uid">{{ from }}</label>
            <input type="text" :id="_uid" name="query" :placeholder="placeholder" autocomplete="off" class="form-control" aria-describedby="help" v-model="query">
            <span class="help-block">{{ help }}</span>
        </div>
        <div class="col-sm-2 centered">
            <button type="submit" class="btn btn-success btn-sm">
                <i class="far fa-2x fa-fw loader" v-bind:class="loader" aria-hidden="true"></i>
            </button>
        </div>
        <div class="col-sm-4">
            <h4>{{ to }}</h4>
            <div class="form-control conversion">
                <table class="table table-striped">
                    <tbody>
                        <tr v-for="(item, key) in result">
                            <td class="min">{{ key }}:&nbsp;</td>
                            <td class="result">
                                <div>
                                    {{ item }}
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</template>

<script>
var VueAutosize = require('./vue-autosize')
Vue.use(VueAutosize)

export default {
    data: function() {
        return {
            query: '',
            response: [],
            result: {},
            loading: false
        };
    },
    props: ['from', 'to', 'url', 'placeholder', 'help', 'joinstr', 'value'],
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
            axios.get(this.url + this.query)
                .then(function (response) {
                    vue.response = response.data.result
                    vue.loading = false
                    vue.createResult()
                });
        },
        createResult: function() {
            this.result = {}
            var vue = this
            _.forEach(vue.response, function(value, key) {
                var it = '';
                if (typeof value == 'string')
                    it += value
                else
                    it += value.join(', ')

                vue.result[key] = it
            })
        }
    },
    created: function() {
        if (this.value) {
            this.query = this.value
            this.getResponse()
        }
    }
}
</script>
