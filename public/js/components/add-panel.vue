<template>
<div class="ff">
  <h1>新增搜尋目標</h1>
  URL: <input v-model="form.url">
  Keyword: <textarea v-model="form.keyword" rows="5" cols="50"></textarea>
  <button v-on:click="addTarget">新增</button>
</div>
</template>

<script>
export default {
  data () {
    return {
      form: {
        url: "",
        keyword: ""
      }
    }
  },

  methods: {
    addTarget () {
        this.$http.post('/api/create', this.form, function(res){
          location.reload()
        })
    }
  }
}
</script>

<style>
.ff * {
  display: inline-block;
  vertical-align: top;
}
</style>
