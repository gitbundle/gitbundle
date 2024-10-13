import $ from './features/jquery.js'

$.each($('[jq-error-message]'), (_, el) =>
  $(el)
    .off('click')
    .on('click', (e) => $(el).parent().remove())
)
