module.exports = {
  methods: {
    push () {
      const rmNode = document.getElementById('baidupush')
      if (rmNode !== null) {
        rmNode.parentNode.removeChild(rmNode)
      }
      const bp = document.createElement('script')
      bp.setAttribute('id', 'baidupush')
      bp.setAttribute('src', 'https://zz.bdstatic.com/linksubmit/push.js')
      const s = document.getElementsByTagName('script')[0]
      s.parentNode.insertBefore(bp, s)
    },
  },
  watch: {
    '$route': 'push',
  },
  mounted () {
    this.push()
  },
}
