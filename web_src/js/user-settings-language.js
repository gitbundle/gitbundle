import $ from 'jquery'
import { initSimpleSvelteSelect } from './modules/svelte-simple-select.js'

function _init(selector) {
  if (!$(selector).length) return

  initSimpleSvelteSelect(
    `${selector} [svelte-user-language-select]`,
    'userLanguage'
  )
}

_init('[jq-user-settings-appearance]')
