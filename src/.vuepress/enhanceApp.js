import ElementUI from 'element-ui'
import { VLazyImagePlugin } from 'v-lazy-image'
import VueLazyload from 'vue-lazyload'
import 'element-ui/lib/theme-chalk/index.css'
import 'aplayer/dist/APlayer.min.css'
import './styles/prismjs-yuchanns.css'

export default ({
  Vue,
  options,
  router,
  siteData
}) => {
  Vue.use(ElementUI)
  Vue.use(VLazyImagePlugin)
  Vue.use(VueLazyload, {
    lazyComponent: true,
  })
}
