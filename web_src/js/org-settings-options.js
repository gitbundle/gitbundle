import $ from 'jquery'
import { initSimpleSvelteSelect } from './modules/svelte-simple-select.js'

export function _init(selector) {
  if (!$(selector).length) return

  $('#org_name').on('keyup', function () {
    const $prompt = $('#org-name-change-prompt')
    const $prompt_redirect = $('#org-name-change-redirect-prompt')
    if ($(this).val().toString().toLowerCase() !== $(this).data('org-name').toString().toLowerCase()) {
      $prompt.removeClass('hidden')
      $prompt_redirect.removeClass('hidden')
    } else {
      $prompt.addClass('hidden')
      $prompt_redirect.addClass('hidden')
    }
  })

  initSimpleSvelteSelect(`${selector} [svelte-org-plugins]`, 'orgPlugins')
}

_init('[jq-org-settings-options]')
