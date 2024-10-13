import { initInstall } from './modules/install.js'

function _init(selector) {
  if (!$(selector).length) return

  initInstall()
}

_init('[jq-page-install]')
