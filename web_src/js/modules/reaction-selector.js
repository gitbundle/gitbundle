import $ from '../features/jquery.js'

const { csrfToken } = window.config

export function initReactionSelector(selector) {
  $(selector)
    .find('[jq-item]')
    .off('click')
    .on('click', function (e) {
      e.preventDefault()
      if ($(this).hasClass('btn-disabled')) return
      const actionURL = $(this).closest(selector).data('action-url')
      const action = $(this).hasClass('btn-primary') ? 'unreact' : 'react'
      const url = `${actionURL}/${action}`
      $.ajax({
        type: 'POST',
        url,
        data: {
          _csrf: csrfToken,
          content: $(this).data('content')
        }
      }).done((resp) => {
        const $temp = $(this).closest('[jq-content]').find('[jq-temp]')
        let $react = $temp.find('[jq-reactions]')
        if (resp && (resp.html || resp.empty)) {
          if (
            (!resp.empty || resp.html === '') &&
            $react.length > 0 &&
            action === 'unreact'
          ) {
            $react.html('').addClass('hidden')
          }

          if ($react.length) {
            if (!resp.empty) {
              $react.html(resp.html).removeClass('hidden')
            }
          } else {
            $react = $(resp.html)
            $temp.append($react)
          }

          if (!resp.empty) {
            // NOTE: should double check this for updated dom
            initReactionSelector($react)
          }
        }
      })
    })
}
