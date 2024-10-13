import $ from './features/jquery.js'
import { initArrowToChevron } from './common/arrow-to-chevron.js'

function _init(selector) {
  const $els = $(`${selector}`)
  if (!$els.length) return

  initArrowToChevron(`${selector} [jq-signin]`)
}

_init('[jq-navbar]')
