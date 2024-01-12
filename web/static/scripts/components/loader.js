
export default (enable = false) => {
  const {body} = document
  const loadingElement = `
      <div id="loader">
        <span></span>
        <p>Loading...</p>
      </div>
  `
  const loaderDom = body.querySelector('#loader')

  !enable && loaderDom ? loaderDom.remove():body.insertAdjacentHTML('beforeend', loadingElement)
}
