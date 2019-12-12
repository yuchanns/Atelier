const path = require('path')
const headerconfig = require('./config/header')

module.exports = {
  title: 'yuchanns\' Atelier',

  description: '一介码农的炼金工坊。爱代码、爱学习、爱交流。php|scheme|python|golang',

  locales: {
    '/': {
      lang: 'zh-CN',
    },
  },

  clientRootMixin: path.resolve(__dirname, 'mixin/mixin.js'),

  evergreen: true,

  plugins: [
    ['mathjax', {
      target: 'svg',
      macros: {
        '*': '\\times',
      },
    }],
    ['vuepress-plugin-code-copy', true],
    'reading-progress',
    'vuepress-plugin-baidu-autopush',
    'vuepress-plugin-seo',
    ['vuepress-plugin-sitemap', {
      hostname: 'https://www.yuchanns.xyz',
    }],
    [ 'feed', {
      canonical_base: 'https://www.yuchanns.xyz',
    }],
  ],

  markdown: {
    lineNumbers: false,
    toc: { includeLevel: [2] },
    extendMarkdown: md => {
      md.use(require('./plugins/markdown-it/traffic-lights'))
    },
  },

  head: [
    ['link', { rel: 'icon', type: 'image/png', href: '/android-chrome-192x192.png', sizes: '192x192' }],
    ['link', { rel: 'shortcut icon', href: '/favicon.ico' }],
    ['meta', { name: 'google-site-verification', content: 'h0GK-apopUhINJJe5Jp3XopZswk6EK_JQT_fVMrs6A0' }],
    ['meta', { name: 'baidu-site-verification', content: 'OeVj0fzw4S' }],
  ],

  chainWebpack: (config, isServer) => {
    if (isServer === false) {
      config.optimization.splitChunks({
        maxInitialRequests: 5,
        cacheGroups: {
          vue: {
            test: /[\\/]node_modules[\\/](vue|vue-router|vssue)[\\/]/,
            name: 'vendor.vue',
            chunks: 'all',
          },
          commons: {
            test: /[\\/]node_modules[\\/]/,
            priority: -10,
            name: 'vendor.commons',
            chunks: 'all',
          },
        },
      })
    }
  },

  theme: 'meteorlxy',

  themeConfig: {
    lang: Object.assign(require('vuepress-theme-meteorlxy/lib/langs/zh-CN'), {
      home: 'Atelier',
      posts: 'Recipes',
    }),
    personalInfo: {
      nickname: 'yuchanns',
      description: '面向GitHub编程',
      email: 'airamusume@gmail.com',
      location: 'Shenzhen, China',

      avatar: '/images/bronya.gif',

      sns: {
        github: {
          account: 'yuchanns',
          link: 'https://github.com/yuchanns',
        },
      },
    },

    header: headerconfig,

    footer: {
      // 是否显示 Powered by VuePress
      poweredBy: true,

      // 是否显示使用的主题
      poweredByTheme: true,

      // 添加自定义 footer (支持 HTML)
      custom: '<div style="margin-bottom: 0.5em;display: inline-block"><span id="busuanzi_container_site_pv">你是第<span id="busuanzi_value_site_pv"></span>位访问者</span></div><br/><span id="year"></span> <a href="http://beian.miit.gov.cn">粤ICP备19127765号</a>',
    },

    lastUpdated: true,

    nav: [
      { text: 'HOME', link: '/', exact: true },
      { text: 'PHP', link: '/posts/categories/php', exact: false },
      { text: 'PYTHON', link: '/posts/categories/python', exact: false },
      { text: 'GOLANG', link: '/posts/categories/golang', exact: false },
      { text: 'WIZARD', link: '/posts/tags/魔法书', exact: false },
      { text: 'AWESOME', link: '/awesome/', exact: false },
      { text: 'TIMELINE', link: '/timeline/', exact: false },
      { text: 'YuC\'s', link: 'http://yuc.wiki/', exact: false },
    ],

    comments: {
      owner: 'yuchanns',
      repo: 'Atelier',
      clientId: '779fd70f4ac32b536176',
      clientSecret: '3ca68e15fd0ff7e5f0fd2679e7aa2dc56fb54009',
    },
    pagination: {
      perPage: 10,
    },
  },
}
