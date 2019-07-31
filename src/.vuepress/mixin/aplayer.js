module.exports = {
  mounted () {
    const aplayer = document.createElement('div')
    aplayer.setAttribute('id', 'aplayer')
    const footer = document.getElementsByTagName('footer')[0]
    footer.appendChild(aplayer)
    const APlayer = require('aplayer')
    // eslint-disable-next-line no-unused-vars
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
      fixed: true,
    })
    setTimeout(function () {
      ap.play()
    }, 3000)
  },
}
