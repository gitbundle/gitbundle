import $ from './features/jquery.js'

function _init(selector) {
  if (!$(selector).length) return

  // Milestones
  $('#clear-date')
    .off('click')
    .on('click', () => {
      $('#deadline').val('')
      return false
    })
}

_init('[jq-repository-new-milestone]')
