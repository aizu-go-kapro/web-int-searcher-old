<template>
  <div class="searchForm">
    <input v-model="words" placeholder="">
    <button v-on:click="">Search</button>
  </div>
</template>

<script>
import { mapActions, mapGetters } from 'vuex'

export default {
    name: 'seach from',
    props: {
        words: String,
    },
    methods: mapActions('Links', {
      'buttonAction': 'buttonAction'
    }),
}
</script>

<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>