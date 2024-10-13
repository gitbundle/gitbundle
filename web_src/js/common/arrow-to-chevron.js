import $ from '../features/jquery.js'

export function initArrowToChevron(selector) {
  $(selector).on("mouseenter", function(e) {
    $(this).children('.octicon-arrow-right').removeClass('hidden')
    $(this).children(".octicon-chevron-right").addClass('hidden')
  }).on("mouseleave", function(e) {
    $(this).children('.octicon-arrow-right').addClass('hidden')
    $(this).children(".octicon-chevron-right").removeClass('hidden')
  })
}