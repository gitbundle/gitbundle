import $ from '../features/jquery.js'

export function initJqInputSearch(inputSelector, itemSelector) {
  $(inputSelector).on('input', function (e) {
    e.preventDefault()
    const keyword = $(this).val()
    // eslint-disable-next-line array-callback-return
    $(`${itemSelector}:not(:icontains(${keyword}))`).filter(function () {
      $(this).addClass('!hidden')
    })
    // eslint-disable-next-line array-callback-return
    $(`${itemSelector}:icontains(${keyword})`).filter(function () {
      $(this).removeClass('!hidden')
    })
  })

  return function () {
    $(inputSelector).val('')
    $(itemSelector).each((_, el) => $(el).removeClass('!hidden'))
  }
}
