<template>
  <div class="container">
    <iframe :src="src" frameborder="0" allowfullscreen class="video"></iframe>
  </div>
</template>
<script>
export default {
  name: 'video-iframe',
  props: [
    'src'
  ]
}
</script>
<style lang="stylus" scoped>
.container {
    position: relative;
    width: 100%;
    height: 0;
    padding-bottom: 56.25%;
}
.video {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
}
</style>