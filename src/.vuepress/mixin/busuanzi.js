module.exports = {
  mounted () {
    // const bsz = document.createElement('span')
    // bsz.setAttribute('id', 'busuanzi_container_site_pv')
    // bsz.innerHTML = '你是第<span id="busuanzi_value_site_pv"></span>位访问者'
    // const footer = document.getElementsByTagName('footer')[0]
    // footer.appendChild(bsz)
    document.getElementById('present').innerHTML = new Date().getFullYear()
    require('busuanzi/bsz.pure.mini')
  },
}
