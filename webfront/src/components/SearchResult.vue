<template>
  <div class="result">
  <p v-bind:url="url">{{links.url}}</p>
  <p v-bind:title="title">{{links.title}}</p>
  <p v-bind:text="text">{{links.text}}</p>
  </div>
</template>

<script>
export default {
    name: 'result row',
    props: {
       links: {},
    }
}
</script>