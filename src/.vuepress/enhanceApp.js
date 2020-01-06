import { VLazyImagePlugin } from 'v-lazy-image'
import VueLazyload from 'vue-lazyload'

export default ({
  Vue,
  options,
  router,
  siteData,
}) => {
  Vue.use(VLazyImagePlugin)
  Vue.use(VueLazyload, {
    lazyComponent: true,
  })
}
