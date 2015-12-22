<template>
<div>
  <h1>{{ msg }}</h1>
  <!-- <button v-on:click="addTarget">新增</button> -->
</div>
</template>

<script>
export default {
  data () {
    return {
      msg: "Rating Counter"
    }
  },

  methods: {
    addTarget () {
        window.location.replace("/add-target")
    }
  }
}
</script>
