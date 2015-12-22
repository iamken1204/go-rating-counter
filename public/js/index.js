import Vue from 'vue'
import App from './components/app.vue'
import List from './components/list.vue'
import SearchPanel from './components/search.vue'

Vue.use(require('vue-resource'))
Vue.http.options.emulateJSON = true

new Vue({
    el: 'body',
    components: {
        app: App,
        search: SearchPanel,
        list: List
    }
})
