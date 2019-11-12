// const path = require('path')
const headerconfig = require('./config/header')

module.exports = {
  title: 'yuchanns\' Atelier',

  description: 'Github Oriented Programming',

  locales: {
    '/': {
      lang: 'zh-CN',
    },
  },

  // clientRootMixin: path.resolve(__dirname, 'mixin/mixin.js'),

  plugins: [
    ['mathjax', {
      target: 'svg',
      macros: {
        '*': '\\times',
      },
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
      description: 'You\'re written in my memory.<br>I can\'t erase this.<br>The thought of being\'s fleeting.<br>Chasing a dream.',
      email: 'airamusume@gmail.com',
      location: 'Shenzhen, China',

      avatar: 'https://avatars3.githubusercontent.com/u/25029451',

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
      custom: '<a href="http://beian.miit.gov.cn">粤ICP备19127765号</a>',
    },

    lastUpdated: true,

    nav: [
      { text: 'Home', link: '/', exact: true },
      // { text: 'Posts', link: '/posts/', exact: false },
      { text: 'Golang', link: '/posts/categories/golang', exact: false },
      { text: 'Wizard', link: '/posts/tags/魔法书', exact: false },
      { text: 'Awesome', link: '/awesome/', exact: false },
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
