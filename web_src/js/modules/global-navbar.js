import GlobalNavbar from '../components/GlobalNavbar.svelte'

function initGlobalNavbar() {
  const el = document.getElementById('global-navbar')
  if (!el) return

  el.innerHTML = ''
  new GlobalNavbar({
    target: el
  })
}

initGlobalNavbar()
