<template>
  <div>
    <h1>{{message}}</h1>
    <button v-on:click="setM">click me</button>
  </div>
</template>

<script>
export default {
  data () {
    return {
        message: ""
    }
  },
  ready () {
    this.init()
  },
  methods: {
    init () {
        this.message = "Init Method-2"
    },
    setM () {
        this.message = "Setted message."
    }
  },
  destroyed () {
    clearInterval(this.handle)
  }
}
</script>
