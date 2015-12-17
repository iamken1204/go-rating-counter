import Vue from 'vue'
import App from './app.vue'

new Vue({
    el: 'body',
    components: {
        app: App
    }
})


// var host_prefix = 'http://localhost:1234';
// $(document).ready(function() {
//     // JSONP version - add 'callback=?' to the URL - fetch the JSONP response to the request
//     $("#jsonp-button").click(function(e) {
//         e.preventDefault();
//         // The only difference on the client end is the addition of 'callback=?' to the URL
//         var url = host_prefix + '/viewtest?callback=?';
//         $.getJSON(url, function(jsonp) {
//             console.log(jsonp);
//             $("#jsonp-response").html(JSON.stringify(jsonp, null, 2));
//         });
//     });
// });

// Vue.config.delimiters = ['(%', '%)'];

// var demo = new Vue({
//   el: '#demo',
//   data: {
//     message: 'Hello Vue.js!'
//   },
//   ready: function() {
//     this.iii()
//   },
//   methods: {
//     iii: function() {
//         console.log('aa')
//         this.message = "WOWOWOWOWWOWOOW"
//     }
//   }
// })

// var demo = new Vue({
//     el: '#vtest',

//     data: {
//         message: "Hello Vue!"
//     },

//     ready: function() {
//         this.init()
//     },

//     methods: {
//         init: function() {
//             console.log(Vue.config.delimiters)
//             this.message = "Hello Vue!!!"
//         },
//     }
// })
