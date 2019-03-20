<template>
  <div class="banner">
    <div class="container">
      <div class="center">
        <ClientOnly>
            <vueCanvasNest
              :config="{color:'0,0,0', count: 10, opacity: 0.7}"
              :el="'.center'"
            />
        </ClientOnly>
        <slot />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TheHeaderBanner',
  mounted () {
    this.$options.components.vueCanvasNest = require('vue-canvas-nest').default
  },
}
</script>

<style lang="stylus" scoped>
@require '~@theme/styles/variables'

.banner
  cursor default
  height $bannerHeight
  border-bottom 1px solid $borderColor
  color #fff
  .container
    height 100%
    display flex
    align-items center
    justify-content center
    .center
      z-index 19
      text-align center
      transition all 0.2s ease-in-out
      &:hover
        transform scale(1.2)

</style>
