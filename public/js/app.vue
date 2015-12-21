<template>
<div>
  <h1>{{ msg }}</h1>
  <button v-on:click="addTarget">新增</button
</div>
</template>

<script>
export default {
  data () {
    return {
      msg: "Hello Vue app!"
    }
  },

  methods: {
    addTarget () {
        window.location.replace("/add-target")
    }
  }
}
</script>
