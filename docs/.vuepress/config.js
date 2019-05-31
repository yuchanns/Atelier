const path = require('path')

module.exports = {
  title: '(yuchanns (Atelier))',

  description: 'focus on python|php',

  locales: {
    '/': {
      lang: 'zh-CN',
    },
  },
  plugins: [
    ['mathjax', {
      target: 'svg',
      macros: {
        '*': '\\times',
      },
    }],
  ],
  markdown: {
    extendMarkdown: md => {
      md.use(require('./components/markdown-it-controls/index'))
    },
  },

  head: [
    ['link', { rel: 'icon', href: '/yuchanns.png' }],
    ['meta', { name: 'google-site-verification', content: 'h0GK-apopUhINJJe5Jp3XopZswk6EK_JQT_fVMrs6A0' }],
    ['meta', { name: 'baidu-site-verification', content: 'OeVj0fzw4S' }],
  ],

  theme: path.resolve(__dirname, '../../lib'),

  themeConfig: {
    lang: {
      home: '首页',
      posts: '文章',
      category: '分类',
      categories: '分类',
      allCategories: '全部',
      tag: '标签',
      tags: '标签',
      createdAt: '发布时间',
      updatedAt: '最后修改',
      prevPost: '上一篇',
      nextPost: '下一篇',
    },
    personalInfo: {
      nickname: 'yuchanns',
      description: 'Either you or time plays trick on the other one',
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

    headerBackground: {
      // url: '/assets/img/bg.jpg',

      useGeo: true,
    },

    lastUpdated: true,

    nav: [
      { text: 'Home', link: '/', exact: true },
      { text: 'Posts', link: '/posts/', exact: false },
      { text: 'Snippets', link: '/snippets/', exact: false },
      { text: 'Awesome', link: '/awesome/', exact: false },
    ],

    comments: {
      owner: 'yuchanns',
      repo: 'Atelier',
      clientId: '779fd70f4ac32b536176',
      clientSecret: '3ca68e15fd0ff7e5f0fd2679e7aa2dc56fb54009',
    },
  },
}
