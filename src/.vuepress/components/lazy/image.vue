<template>
  <v-lazy-image :src="src" />
</template>

<script>
export default {
  props: {
    src: {
      type: String,
      default: '',
    },
  },
}
</script>

<style scoped>
.v-lazy-image {
  filter: blur(10px);
  transition: filter 0.7s;
}
.v-lazy-image-loaded {
  filter: blur(0);
}
</style>
