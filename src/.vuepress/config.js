const markdownItCenterText = require('markdown-it-center-text')

module.exports = {
  title: 'C-Sekai',

  description: '代码爱好者，半吊子码农',

  theme: 'yuchanns',

  head: [
    ['link', { rel: 'shortcut icon', href: '/favicon.ico' }],
    ['meta', { name: 'keywords', content: 'Atelier,yuchanns,PHP,Python,Golang,Go,Go语言,Scheme,Lisp,码农,程序猿,炼金工坊,vuepress' }],
    ['meta', { name: 'google-site-verification', content: 'K1weNq67k6Udp1q4Jr7AZahSmfHcaYXBnYHHCxPgh_I' }],
    ['meta', { name: 'baidu-site-verification', content: 'YIQSNieN3r' }],
  ],

  locales: {
    '/': {
      lang: 'en-US',
    },
  },

  evergreen: true,

  plugins: [
    ['mathjax', {
      target: 'svg',
      macros: {
        '*': '\\times',
      },
    }],
    ['vuepress-plugin-zooming', {
      selector: '.content__default img',
    }],
    'reading-progress',
    ['vuepress-plugin-sitemap', {
      hostname: 'https://yuchanns.org',
    }],
  ],

  markdown: {
    lineNumbers: false,
    extendMarkdown: md => md.use(markdownItCenterText),
  },

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

  themeConfig: {
    lang: {
      home: 'Home',
      navigation: 'Navigation',
      categories: 'Category',
      tags: 'Tags',
      archive: 'Archive',
      prev: '上一篇',
      next: '下一篇',
      more: 'More',
      createdAt: '创建于'
    },

    pagination: {
      lengthPerPage: 10,
    },

    sns: {
      github: {
        account: 'yuchanns',
        link: 'https://github.com/yuchanns',
      },
      twitter: {
        account: 'airamusume',
        link: 'https://twitter.com/airamusume',
      },
    },

    categories: {
      Vue: { color: '#2c3e50', logo: '/images/vue.png', desc: 'Vue.js is a JavaScript framework for building interactive web applications.' },
      Document: { color: '#e34c26', desc: 'Document is a guidebook for users.' },
      golang: { color: '#00ADD8', desc: 'go语言学习使用笔记', logo: '/images/categories/golang.jpeg' },
      python: { color: '#3572A5', desc: '人生苦短，我选py！', logo: '/images/categories/python.png' },
      php: { color: '#4F5D95', desc: '世界上最好的语言！', logo: '/images/categories/php.png' },
      学习笔记: { color: '#e34c26', desc: 'study of sicp' },
      lisp: { color: '#c065db', desc: '学着玩儿！', logo: '/images/categories/lisp.png' },
    },
  },
}
