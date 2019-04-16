const path = require('path')

module.exports = {
  title: 'yuchanns\'Atelier',

  description: 'focus on python|php',

  locales: {
    '/': {
      lang: 'zh-CN',
    },
  },
  plugins: [
    'flowchart',
  ],

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
      description: 'A non professional programmer which attempt to be professional',
      email: 'airamusume@gmail.com',
      location: 'Shenzhen, China',

      avatar: 'https://avatars3.githubusercontent.com/u/25029451',

      sns: {
        github: {
          account: 'yuchanns',
          link: 'https://github.com/yuchanns',
        }
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
      { text: 'Awesome', link: '/awesome/', exact: false }
    ],

    comments: {
      owner: 'yuchanns',
      repo: 'yuchanns',
      clientId: 'b2f9320aa657a856626c',
      clientSecret: 'a572d895f1daa956cb872ad9015e2ff80941c632',
    },
  },
}
