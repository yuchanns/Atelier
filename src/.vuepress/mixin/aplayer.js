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
      audio: [
        {
          name: 'Cyberangel',
          artist: 'Hanser',
          url: '/assets/music/Cyberangle.flac',
          cover: '/assets/music/Cyberangle.gif',
          lrc: '/assets/music/Cyberangle.lrc',
        },
        {
          name: 'Gray Zone',
          artist: 'Stack',
          url: '/assets/music/GrayZone.flac',
          cover: '/assets/music/GrayZone.jpg',
          lrc: '/assets/music/GrayZone.lrc',
        },
        {
          name: 'Bon☆Dance',
          artist: 'Stack',
          url: '/assets/music/Bon☆Dance.flac',
          cover: '/assets/music/Bon☆Dance.jpg',
          lrc: '/assets/music/Bon☆Dance.lrc',
        },
        {
          name: 'Starlight Dance Floor',
          artist: 'Ark Brown',
          url: '/assets/music/StarlightDanceFloor.mp3',
          cover: '/assets/music/StarlightDanceFloor.jpg',
        },
        {
          name: 'Girls in the Mirror',
          artist: 'Ark Brown',
          url: '/assets/music/GirlsintheMirror.flac',
          cover: '/assets/music/GirlsintheMirror.jpg',
        },
      ],
      fixed: true,
    })
    setTimeout(function () {
      ap.play()
    }, 3000)
  },
}
