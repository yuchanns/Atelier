module.exports = {
  data () {
    return {
      audio: [
        {
          name: 'オープニング',
          artist: '片岡真央',
          url: 'http://music.163.com/song/media/outer/url?id=1309814525.mp3',
          cover: 'http://p1.music.126.net/sHMTWM6Z0zXlqz98Wdbl3A==/109951163555243264.jpg?param=130y130',
        },
        {
          name: 'リンクの記憶「英傑 ミファー」',
          artist: '岩田恭明',
          url: 'http://music.163.com/song/media/outer/url?id=1309814566.mp3',
          cover: 'http://p1.music.126.net/sHMTWM6Z0zXlqz98Wdbl3A==/109951163555243264.jpg?param=130y130',
        },
        {
          name: '再会 リーバル',
          artist: '岩田恭明',
          url: 'http://music.163.com/song/media/outer/url?id=1309814590.mp3',
          cover: 'http://p1.music.126.net/sHMTWM6Z0zXlqz98Wdbl3A==/109951163555243264.jpg?param=130y130',
        },
        {
          name: '太陽の集落~バルバレ',
          artist: '鈴木まり香',
          url: 'http://music.163.com/song/media/outer/url?id=28830127.mp3',
          cover: 'http://p1.music.126.net/HlVPKh7Y-RY5mpXpKZPKPA==/5976945209054550.jpg?param=130y130',
        },
        {
          name: '新大陸への礎(いしずえ) ～ 調査拠点アステラ',
          artist: '牧野忠義',
          url: 'http://music.163.com/song/media/outer/url?id=537854168.mp3',
          cover: 'http://p2.music.126.net/EB2CtT2MJMfkrtoKg_rhJg==/109951163156648575.jpg?param=130y130',
        },
      ],
    }
  },
  mounted () {
    const aplayer = document.createElement('div')
    aplayer.setAttribute('id', 'aplayer')
    const footer = document.getElementsByTagName('footer')[0]
    footer.appendChild(aplayer)
    const APlayer = require('aplayer')
    let _ = require('lodash')
    const ap = new APlayer({
      container: document.getElementById('aplayer'),
      lrcType: 0,
      audio: _.shuffle(this.audio),
      fixed: true,
    })
    setTimeout(function () {
      ap.play()
    }, 5000)
  },
}
