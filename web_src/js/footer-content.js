import $ from './features/jquery.js'

function initQrcodeToggle(selector) {
  $(`${selector} [jq-toggle]`).hover(
    function () {
      $(this).children('[jq-qrcode]').removeClass('hidden')
    },
    function () {
      $(this).children('[jq-qrcode]').addClass('hidden')
    }
  )
}

function _init(selector) {
  const $els = $(`${selector}`)
  if (!$els.length) return

  initQrcodeToggle(selector)
}

_init('[jq-footer-content]')
