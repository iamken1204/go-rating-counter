<template>
  <div v-for="target in targets">
    <h5>{{ target.id }}</h5>
    <h5>{{ target.keyword }}</h5>
    <h5>{{ target.url }}</h5>
  </div>
</template>

<script>
export default {
  data () {
    return {
        message: "",
        targets: []
    }
  },
  ready () {
    this.init()
  },
  methods: {
    init () {
        this.$http.post('/api/targets', function(targets) {
          this.targets = targets
        })
    }
  },
  destroyed () {
    clearInterval(this.handle)
  }
}
</script>
