import Vue from 'vue'
import App from './app.vue'
import List from './list.vue'

Vue.use(require('vue-resource'))

new Vue({
    el: 'body',
    components: {
        app: App,
        list: List
    }
})
