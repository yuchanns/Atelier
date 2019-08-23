module.exports = {
  mounted () {
    const aplayer = document.createElement('div')
    aplayer.setAttribute('id', 'aplayer')
    const footer = document.getElementsByTagName('footer')[0]
    footer.appendChild(aplayer)
    const APlayer = require('aplayer')
    const ap = new APlayer({
      container: document.getElementById('aplayer'),
      lrcType: 0,
      audio: [
        {
          name: 'ミック.ジャガーに微笑みを',
          artist: '中森明菜',
          url: 'https://m10.music.126.net/20190823232525/5c00ea7e121961636f9cf33b2abb0239/ymusic/2569/a5d1/f96f/a0b8dd43862ce39314195ebb795b071d.mp3',
          cover: 'https://p3.music.126.net/ue1gYMQCssgLa0ylxRD3rg==/625622116216240.jpg?param=300y300',
        },
        {
          name: 'Innocent World',
          artist: 'Vivienne',
          url: 'https://m10.music.126.net/20190823235034/301a7425815553a2df53991a0d317cef/ymusic/f994/eafd/f2f7/3d9872701d58fd97e02a48a1c1707062.mp3',
          cover: 'https://p3.music.126.net/MIpvFQfzLPzvipmFKbS5Jw==/6055010534567895.jpg?param=300y300',
        },
        {
          name: 'GATE OF STEINER',
          artist: '佐佐木惠梨',
          url: 'https://m10.music.126.net/20190823234003/18290492fde169169a66f3c3c855f349/ymusic/6ec2/1a26/a488/61d3432d3d5603f6055f16db25dfb5ff.mp3',
          cover: 'https://p3.music.126.net/z-KAhHeB8AnWi4bBOnpLsA==/3299634402086878.jpg?param=300y300',
        },
        {
          name: '千年の虹',
          artist: '阿兰',
          url: 'https://m10.music.126.net/20190823234424/bc65f1d49812f1eaee8131034558c7e0/ymusic/5416/df13/b7d3/63982b37da4ddb1e6d5db84edc62051b.mp3',
          cover: 'https://p3.music.126.net/wgbcjy8sIDTVQ4UKMjRtPw==/858718581317063.jpg?param=300y300',
        },
        {
          name: 'やさしくしないで',
          artist: '杉恵ゆりか',
          url: 'https://m10.music.126.net/20190823234745/ec152e7193bff20174af84e1ecbd52a9/ymusic/cb5e/2c6b/b7f2/8dbaa9966cad8e78d5659448f1ff4e5e.mp3',
          cover: 'https://p3.music.126.net/Vls3tRRSmLjYNX8_jJsjyw==/2532175280692526.jpg?param=300y300',
        },
      ],
      fixed: true,
    })
    setTimeout(function () {
      ap.play()
    }, 5000)
  },
}
