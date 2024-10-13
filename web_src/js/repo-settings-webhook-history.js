import $ from 'jquery'

const { csrfToken } = window.config

export function initTestDelivery(selector) {
  $(selector).on('click', function () {
    const $this = $(this)
    $this.addClass('loading disabled')
    $.post($this.data('link'), {
      _csrf: csrfToken
    }).done(
      setTimeout(() => {
        window.location.href = $this.data('redirect')
      }, 5000)
    )
  })
}

export function initHistoryToggle(selector) {
  $(`${selector} [jq-toggle-button]`).on('click', function (e) {
    $(this).closest(selector).find('[jq-toggle-content]').toggleClass('hidden')

    // $content
    //   .children('[jq-tabs]')
    //   .find('[data-tab]')
    //   .one('click', function () {

    //     const $tab = $(this).data('tab')
    //     $content
    //       .children(`[jq-content]`)
    //       .find(`[data-tab=${$tab}]`)
    //       .removeClass('hidden')
    //       .siblings()
    //       .addClass('hidden')
    //   })
  })

  $(`${selector} [jq-toggle-content] [jq-tabs]`)
    .find('[data-tab]')
    .on('click', function () {
      $(this).addClass('tab-active').siblings().removeClass('tab-active')
      const $tab = $(this).data('tab')
      $(this)
        .closest(`[jq-tabs]`)
        .siblings('[jq-contents]')
        .find(`[data-tab=${$tab}]`)
        .removeClass('hidden')
        .siblings()
        .addClass('hidden')

      try {
        const $j = $(this).closest(`[jq-tabs]`).siblings('[jq-contents]').find(`[data-tab=${$tab}]`).find(`[jq-json]`)
        const jsonObj = JSON.parse($j.text())
        $j.text(JSON.stringify(jsonObj, null, 2))
      } catch (error) {
        
      }
    })
}

function _init(selector) {
  if (!$(selector).length) return

  initTestDelivery('[jq-test-delivery]')
  initHistoryToggle('[jq-repository-settings-history]')
}

_init('[jq-repo-settings-webhook-history]')
