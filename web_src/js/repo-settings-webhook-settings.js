import $ from 'jquery'

export function initWebHookSettings(selector) {
  if (!$(selector).length) return

  $('[jq-events] input').on('change', function () {
    if ($(this).is(':checked')) {
      $('[jq-events-fields]').removeClass('hidden')
    }
  })
  $('[jq-non-events] input').on('change', function () {
    if ($(this).is(':checked')) {
      $('[jq-events-fields]').addClass('hidden')
    }
  })
}

initWebHookSettings('[jq-webhook-settings]')
