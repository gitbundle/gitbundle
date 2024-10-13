import $ from 'jquery'
import { initUnicodeEscapeButton } from './modules/repo-unicode-escape.js'

function _init(selector) {
  if (!$(selector).length) return

  initUnicodeEscapeButton()
}

_init('[jq-repo-settings-lfs_file]')
