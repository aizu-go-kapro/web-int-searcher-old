<template>
  <div id="app">
    <!-- <img src="./assets/logo.png"> -->
    <Title msg="welcome wisa!"/>
    <Form v-bind:keyword="keyword"/>
    <p>{{keyword}}</p>
    <Results
      v-for="(link) in links"
      v-bind:link="link"
    ></Results>
  </div>
</template>

<script>
import Title from './components/HelloWorld.vue'
import Form from './components/SearchForm.vue'
import Results from './components/SearchResult.vue'
import { mapActions, mapGetters } from 'vuex'

export default {
  name: 'app',
  computed: mapGetters(['Links']),
  components: {
    Title,
    Form,
    Results
  },
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
