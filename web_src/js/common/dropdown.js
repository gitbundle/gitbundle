import $ from 'jquery'

export function initJqDropdownClose(selector, onClose) {
  $(`${selector} > label`).on('click', function (e) {
    // e.preventDefault()
    $(this).closest(selector).addClass('dropdown-open')
  })

  $(document).on('click', function (e) {
    if ($(selector).hasClass('dropdown-open')) {
      if (!$(e.target).closest(`${selector}`).length) {
        onClose && onClose()
        $(`${selector}`).removeClass('dropdown-open')
      }
    }
  })
}
