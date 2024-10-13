import { initRepoCodeView } from './modules/repo-code.js'
import { initUnicodeEscapeButton } from './modules/repo-unicode-escape.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoCodeView()
  initUnicodeEscapeButton()
}

_init('[jq-repo-home]')
