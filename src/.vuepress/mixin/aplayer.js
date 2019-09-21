module.exports = {
  data () {
    return {
      audio: [
        {
          name: 'Beautiful Dreams',
          artist: 'Vivienne',
          url: 'https://music.163.com/song/media/outer/url?id=744728.mp3',
          cover: 'https://p3.music.126.net/QN94YMkikhvQN4132VnQVw==/109951163516679812.jpg?param=300y300',
        },
        {
          name: 'Different Kind of Love',
          artist: 'Vivienne',
          url: 'https://music.163.com/song/media/outer/url?id=450041327.mp3',
          cover: 'https://p3.music.126.net/kcGhHJs0wk6Q0MIICCwvTw==/18503681185518261.jpg?param=300y300',
        },
        {
          name: 'ミック.ジャガーに微笑みを',
          artist: '中森明菜',
          url: 'https://music.163.com/song/media/outer/url?id=4910560.mp3',
          cover: 'https://p3.music.126.net/ue1gYMQCssgLa0ylxRD3rg==/625622116216240.jpg?param=300y300',
        },
        {
          name: 'Innocent World',
          artist: 'Vivienne',
          url: 'https://music.163.com/song/media/outer/url?id=28578496.mp3',
          cover: 'https://p3.music.126.net/MIpvFQfzLPzvipmFKbS5Jw==/6055010534567895.jpg?param=300y300',
        },
        {
          name: 'GATE OF STEINER',
          artist: '佐佐木惠梨',
          url: 'https://music.163.com/song/media/outer/url?id=442562811.mp3',
          cover: 'https://p3.music.126.net/z-KAhHeB8AnWi4bBOnpLsA==/3299634402086878.jpg?param=300y300',
        },
        // {
        //   name: '千年の虹',
        //   artist: '阿兰',
        //   url: 'https://music.163.com/song/media/outer/url?id=22817060.mp3',
        //   cover: 'https://p3.music.126.net/wgbcjy8sIDTVQ4UKMjRtPw==/858718581317063.jpg?param=300y300',
        // },
        {
          name: 'やさしくしないで',
          artist: '杉恵ゆりか',
          url: 'https://music.163.com/song/media/outer/url?id=29764416.mp3',
          cover: 'https://p3.music.126.net/Vls3tRRSmLjYNX8_jJsjyw==/2532175280692526.jpg?param=300y300',
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
