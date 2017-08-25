<template>
    <div class="row equal">
        <div class="col-sm-3 col-sm-offset-2 centered">
            <label :for="_uid">{{ from }}</label>
            <input type="text" :id="_uid" name="query" :placeholder="placeholder" autocomplete="off" class="form-control" aria-describedby="help" autofocus v-model="query">
            <span class="help-block">{{ help }}</span>
        </div>
        <div class="col-sm-2 centered">
            <button type="submit" class="btn btn-success btn-sm">
                <i class="fa fa-2x loader" v-bind:class="loader"></i>
            </button>
        </div>
        <div class="col-sm-4">
            <h4>{{ to }}</h4>
            <textarea class="result form-control" readonly v-model="display" v-autosize></textarea>
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
            loading: false
        };
    },
    props: ['from', 'to', 'url', 'placeholder', 'help', 'joinstr'],
    computed: {
        loader: function() {
            if (this.loading) {
                return ['fa-refresh', 'fa-spin'];
            } else {
                return ['arrow'];
            }
        },
        display: function() {
            return this.response.join(this.joinstr);
        }
    },
    watch: {
        query: function() {
            this.getResponse();
        }
    },
    methods: {
        getResponse: _.debounce(
            function () {
                this.loading = true;
                var vue = this;
                axios.get(this.url + this.query)
                    .then(function (response) {
                        vue.response = response.data;
                        vue.loading = false;
                    });
            }, 250)
    }
}
</script>
