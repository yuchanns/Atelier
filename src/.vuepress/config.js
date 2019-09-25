const path = require('path')

module.exports = {
  title: '🍃(yuchanns (Atelier))',

  description: 'focus on python|php',

  locales: {
    '/': {
      lang: 'zh-CN',
    },
  },

  clientRootMixin: path.resolve(__dirname, 'mixin/mixin.js'),

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
    extendMarkdown: md => {
      md.use(require('./components/markdown-it-controls/index'))
    },
  },

  head: [
    ['link', { rel: 'icon', href: '/yuchanns.png' }],
    ['meta', { name: 'google-site-verification', content: 'h0GK-apopUhINJJe5Jp3XopZswk6EK_JQT_fVMrs6A0' }],
    ['meta', { name: 'baidu-site-verification', content: 'OeVj0fzw4S' }],
  ],

  theme: 'meteorlxy',

  themeConfig: {
    lang: {
      home: 'Atelier',
      posts: 'Recipes',
      category: 'category',
      categories: 'categories',
      allCategories: 'all',
      tag: 'tag',
      tags: 'tags',
      createdAt: 'createdAt',
      updatedAt: 'updatedAt',
      prevPost: 'prev',
      nextPost: 'next',
    },
    personalInfo: {
      nickname: 'yuchanns',
      description: 'Plus Ultra!',
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

    header: {
      background: {
        url: [
          // '/assets/img/1.jpeg',
          // '/assets/img/2.jpeg',
          // '/assets/img/3.jpeg',
          // '/assets/img/4.jpeg',
          // '/assets/img/5.jpeg',
          // '/assets/img/6.jpeg',
          // '/assets/img/7.jpeg',
          // '/assets/img/8.jpeg',
          '/assets/img/houkai3.jpeg',
          '/assets/img/godkiana.jpeg',
        ],
        useGeo: false,
      },
      showTitle: true,
    },

    footer: {
      // 是否显示 Powered by VuePress
      poweredBy: true,

      // 是否显示使用的主题
      poweredByTheme: true,

      // 添加自定义 footer (支持 HTML)
      custom: '<div style="margin: 0.2em 0"><a href="https://app.netlify.com/sites/yuchanns/deploys" ref="nofollow"><img src="https://api.netlify.com/api/v1/badges/d5cbefe0-4b3d-437a-8064-0181fbe0dd23/deploy-status"></img></a> <a href="https://circleci.com/gh/yuchanns/Atelier" ref="nofollow"><img src="https://circleci.com/gh/yuchanns/Atelier.svg?style=svg"></img></a></div><span id="busuanzi_container_site_pv">你是第<span id="busuanzi_value_site_pv" style="color: #F56C6C"></span>位访问者</span><div style="margin: 0.3em 0">Since Year 2018 - <span id="present"></span></div><a href="https://github.com/yuchanns"><img width="149" height="149" src="/images/forkme_left_gray_6d6d6d.png" style="position: fixed; top: -7px; left: -7px; border: 0; z-index: 999;" alt="Fork me on GitHub"></a>',
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
