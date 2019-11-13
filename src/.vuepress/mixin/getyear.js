module.exports = {
  mounted () {
    const year = new Date().getFullYear()
    document.querySelector('#year').innerHTML = '© ' + year
  },
}
