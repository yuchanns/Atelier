<template>
  <div id="vuepress-theme-meteorlxy">
    <ClientOnly>
      <vueCanvasNest
        :config="{color:'88,173,178', count: 500, opacity: 0.9}"
        :el="'#vuepress-theme-meteorlxy'"
      />
    </ClientOnly>
    <TheHeader />

    <TheMain />

    <TheFooter />
  </div>
</template>

<script>
import TheHeader from '../components/TheHeader.vue'
import TheMain from '../components/TheMain.vue'
import TheFooter from '../components/TheFooter.vue'

export default {
  name: 'ThemeMeteorlxy',

  components: {
    TheHeader,
    TheMain,
    TheFooter,
  },
  mounted () {
    this.$options.components.vueCanvasNest = require('vue-canvas-nest').default
  },
}
</script>

<style lang="stylus" src="../styles/index.styl" />
