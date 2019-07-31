module.exports = {
  mounted () {
    const bsz = document.createElement('div')
    bsz.setAttribute('id', 'aplayer')
    const footer = document.getElementsByTagName('footer')[0]
    footer.appendChild(bsz)
    const APlayer = require('aplayer')
    // eslint-disable-next-line no-unused-vars
    setTimeout(function () {
      const ap = new APlayer({
        container: document.getElementById('aplayer'),
        lrcType: 3,
        audio: [{
          name: 'Cyberangel',
          artist: 'Hanser',
          url: 'https://www.yuchanns.xyz/assets/music/Cyberangle.flac',
          cover: 'https://avatars3.githubusercontent.com/u/25029451',
          lrc: 'https://www.yuchanns.xyz/assets/music/Cyberangle.lrc',
        }],
        mini: true,
        fixed: true,
        autoplay: true,
      })
    }, 1000)
  },
}
