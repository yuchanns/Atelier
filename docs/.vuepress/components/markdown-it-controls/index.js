/* eslint-disable no-tabs */
module.exports = function controlsPlugin (md) {
  const fence = md.renderer.rules.fence
  md.renderer.rules.fence = (...args) => {
    const rawCode = fence(...args)
    const finalCode = rawCode
      .replace('<!--afterbegin-->', `<div style="position: relative;top: 10px;margin-left: 14px;z-index: 2;"><svg xmlns="http://www.w3.org/2000/svg" width="54" height="14" viewBox="0 0 54 14">
			<g fill="none" fillrule="evenodd" transform="translate(1 1)">
				<circle cx="6" cy="6" r="6" fill="#FF5F56" stroke="#E0443E" strokewidth=".5"></circle>
				<circle cx="26" cy="6" r="6" fill="#FFBD2E" stroke="#DEA123" strokewidth=".5"></circle>
				<circle cx="46" cy="6" r="6" fill="#27C93F" stroke="#1AAB29" strokewidth=".5"></circle>
			</g>
		</svg></div><!--afterbegin-->`)
    return finalCode
  }
}
