const aplayer = require('./aplayer')
const getyear = require('./getyear')
const bsz = require('./busuanzi')

export default {
  mixins: [
    aplayer,
    getyear,
    bsz,
  ],
}
