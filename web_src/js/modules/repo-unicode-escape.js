import $ from '../features/jquery.js'

export function initUnicodeEscapeButton() {
  $(document)
    .off('click', 'a.escape-button')
    .on('click', 'a.escape-button', (e) => {
      e.preventDefault()
      $(e.target)
        .parents('.file-content, .non-diff-file-content')
        .find('.file-code, .file-view')
        .addClass('unicode-escaped')
      $(e.target).addClass('hidden')
      $(e.target).siblings('a.unescape-button').removeClass('hidden')
    })
  $(document)
    .off('click', 'a.unescape-button')
    .on('click', 'a.unescape-button', (e) => {
      e.preventDefault()
      $(e.target)
        .parents('.file-content, .non-diff-file-content')
        .find('.file-code, .file-view')
        .removeClass('unicode-escaped')
      $(e.target).addClass('hidden')
      $(e.target).siblings('a.escape-button').removeClass('hidden')
    })
  $(document).on('click', 'a.toggle-escape-button', (e) => {
    e.preventDefault()
    const fileContent = $(e.target).parents(
      '.file-content, .non-diff-file-content'
    )
    const fileView = fileContent.find('.file-code, .file-view')
    if (fileView.hasClass('unicode-escaped')) {
      fileView.removeClass('unicode-escaped')
      fileContent.find('a.unescape-button').addClass('hidden')
      fileContent.find('a.escape-button').removeClass('hidden')
    } else {
      fileView.addClass('unicode-escaped')
      fileContent.find('a.unescape-button').removeClass('hidden')
      fileContent.find('a.escape-button').addClass('hidden')
    }
  })
}
