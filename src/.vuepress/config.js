const markdownItFootnote = require('markdown-it-footnote')

module.exports = {
  title: 'yuchanns\'Atelier',

  description: 'Coding to die',

  theme: 'hermit',

  head: [
    ['link', { rel: 'shortcut icon', href: '/favicon.ico' }],
    ['link', { rel: 'manifest', href: '/manifest.json' }],
    ['meta', { name: 'keywords', content: 'golang,go,yuchanns,Atelier' }],
    ['meta', { name: 'google-site-verification', content: 'K1weNq67k6Udp1q4Jr7AZahSmfHcaYXBnYHHCxPgh_I' }],
    ['meta', { name: 'baidu-site-verification', content: 'o24pKRlUdn' }],
  ],

  locales: {
    '/': {
      lang: 'en-US'
    }
  },

  plugins: [
    ['mathjax', {
      target: 'svg',
      macros: {
        '*': '\\times',
      },
    }],
    ['vuepress-plugin-zooming', {
      selector: '.content img',
    }],
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
    }],
    ['flowchart']
  ],

  markdown: {
    lineNumbers: false,
    extendMarkdown: md => {
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
    lang: 'en-US',

    author: {
      name: 'yuchanns',
      url: 'https://yuchanns.org'
    },

    nav: [
      { text: 'Awesome', link: '/awesome/' },
      { text: 'About', link: '/about/' },
      { text: 'YuC\'s', link: 'http://yuc.wiki/' }
    ],

    sns: [
      // { type: 'codepen', url: 'https://codepen.io/' },
      // { type: 'facebook', url: 'https://www.facebook.com' },
      { type: 'github', url: 'https://github.com/yuchanns' },
      // { type: 'gitlab', url: 'https://gitlab.com/yuchanns' },
      // { type: 'instagram', url: 'https://www.instagram.com/' },
      // { type: 'linkedin', url: 'https://www.linkedin.com/' },
      // { type: 'slack', url: 'https://slack.com/' },
      // { type: 'stackoverflow', url: 'https://stackoverflow.com/users/10038512/yuchanns?tab=profile' },
      // { type: 'telegram', url: 'https://telegram.org/' },
      // { type: 'twitch', url: 'https://www.twitch.tv/' },
      { type: 'twitter', url: 'https://twitter.com/Airamusume' },
      // { type: 'youtube', url: 'https://www.youtube.com/' },
      { type: 'email', url: 'mailto:airamusume@gmail.com' },
      // { type: 'dribbble', url: 'https://dribbble.com/' },
      // { type: 'behance', url: 'https://www.behance.net/' },
      // { type: 'freepik', url: 'https://www.freepik.com/' },
      // { type: 'adobestock', url: 'https://stock.adobe.com/' },
      // { type: 'shutterstock', url: 'https://www.shutterstock.com/' },
      // { type: '123rf', url: 'https://www.123rf.com/' },
      // { type: 'dreamstime', url: 'https://www.dreamstime.com/' },
      // { type: 'paypal', url: 'https://www.paypal.com/' }
    ],

    pagination: {
      lengthPerPage: 10
    },

    comment: {
      service: 'vssue',
      owner: 'yuchanns',
      repo: 'Atelier',
      clientId: '779fd70f4ac32b536176',
      clientSecret: '3ca68e15fd0ff7e5f0fd2679e7aa2dc56fb54009'
    },

    feed: {
      canonical_base: 'https://yuchanns.org'
    }
  }
}