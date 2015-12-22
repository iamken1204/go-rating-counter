import Vue from 'vue'
import AddPanel from './components/add-panel.vue'
import List from './components/list.vue'

Vue.use(require('vue-resource'))
Vue.http.options.emulateJSON = true

new Vue({
    el: 'body',
    components: {
      addpanel: AddPanel,
      list: List
    }
})
