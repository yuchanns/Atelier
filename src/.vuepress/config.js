const markdownItCenterText = require('markdown-it-center-text')
const markdownItFootnote = require('markdown-it-footnote')

module.exports = {
  title: 'yuchanns\' Atelier',

  description: '代码爱好者，信仰编程',

  theme: 'yuchanns',

  head: [
    ['link', { rel: 'shortcut icon', href: '/favicon.ico' }],
    ['link', { rel: 'manifest', href: '/manifest.json' }],
    ['meta', { name: 'keywords', content: 'Atelier,yuchanns,PHP,Python,Golang,Go,Go语言,Scheme,Lisp,码农,程序猿,炼金工坊,vuepress' }],
    ['meta', { name: 'google-site-verification', content: 'K1weNq67k6Udp1q4Jr7AZahSmfHcaYXBnYHHCxPgh_I' }],
    ['meta', { name: 'baidu-site-verification', content: 'o24pKRlUdn' }],
  ],

  locales: {
    '/': {
      lang: 'zh-CN',
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
    ['@vuepress/pwa', {
      serviceWorker: true,
      updatePopup: true
    }],
    ['seo', {
      description: $page => {
        if ('description' in $page.frontmatter) {
          return $page.frontmatter.description
        } else {
          return $page.title
        }
      }
    }]
  ],

  markdown: {
    lineNumbers: false,
    extendMarkdown: md => {
      md.use(markdownItCenterText)
      md.use(markdownItFootnote)
    },
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
    lang: 'zh-CN',

    author: {
      name: 'yuchanns',
      avatar: '/yuchanns.jpg',
      desc: '面向Github编程',
      job: '后端工程师',
      location: '深圳',
      email: 'airamusume@gmail.com',
      link: 'github.com/yuchanns'
    },

    vssue: {
      platform: 'github',
      owner: 'yuchanns',
      repo: 'Atelier',
      clientId: '779fd70f4ac32b536176',
      clientSecret: '3ca68e15fd0ff7e5f0fd2679e7aa2dc56fb54009'
    },

    sitemap: {
      hostname: 'https://yuchanns.org',
    },

    pagination: {
      lengthPerPage: 5,
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
      简单易懂的现代魔法: { color: '#4C0168', desc: '编程，即现代魔法；程序员，即现代法师。想要发挥出火球术的威力，还需探究和分析好法术的实现原理', logo: '/images/categories/codingwizard.png' },
      计划: { color: '#F3BD42', desc: '首先定下一个小目标，比如说赚他一个亿！', logo: '/images/categories/plan.jpg' },
    },
  },
}
