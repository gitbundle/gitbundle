import $ from 'jquery'

export function initAdminEmails(selector) {
  if (!$(selector).length) return

  function linkEmailAction(e) {
    const $this = $(this)
    $('#form-uid').val($this.data('uid'))
    $('#form-email').val($this.data('email'))
    $('#form-primary').val($this.data('primary'))
    $('#form-activate').val($this.data('activate'))
    $('#change-email-modal').prop('checked', true)
    e.preventDefault()
  }
  $('[jq-link-email-action]').on('click', linkEmailAction)
}
