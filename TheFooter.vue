<template>
  <footer class="footer">
    <p
      v-if="sns"
      class="sns-links"
    >
      <a
        v-for="(item, name) in sns"
        :key="name"
        class="sns-link"
        :href="item.link"
        target="_blank"
      >
        <IconSns
          :name="name"
          :account="item.account"
        />
      </a>
    </p>

    <p>
      <span>Powered by </span>

      <a
        href="https://vuepress.vuejs.org"
        target="_blank"
      >
        Vuepress
      </a>
    </p>
    <span id="busuanzi_container_site_pv">本站总访问量<span id="busuanzi_value_site_pv"></span>次</span>
    <p>
      <a 
        href="http://beian.miit.gov.cn/"
        target="_blank"
        style="color: #666"
      >
          闽ICP备19005068号
      </a>
    </p>
  </footer>
</template>

<script>
import IconSns from './IconSns.vue'

export default {
  name: 'TheFooter',

  components: {
    IconSns,
  },

  computed: {
    sns () {
      return this.$site.themeConfig.personalInfo.sns || null
    },
  },
  mounted () {
    require('busuanzi/bsz.pure.mini')
  }
}
</script>

<style lang="stylus" scoped>
@require '~@theme/styles/variables'

.footer
  color $grayTextColor
  padding-bottom 1.5rem
  text-align center
  border-top 1px solid $borderColor
</style>
