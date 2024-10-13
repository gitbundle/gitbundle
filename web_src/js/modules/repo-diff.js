import { initCompMarkupContentPreviewTab } from '../common/MarkupContentPreview.js'
import $ from '../features/jquery.js'
import {
  countAndUpdateViewedFiles,
  initViewedCheckboxListenerFor
} from './pull-view-file.js'
import { initRepoIssueContentHistory } from './repo-issue-content.js'

export function assignMenuAttributes(menu) {
  const id = Math.floor(Math.random() * Math.floor(1000000))
  menu.attr('data-write', menu.attr('data-write') + id)
  menu.attr('data-preview', menu.attr('data-preview') + id)
  menu.children().each(function () {
    const tab = $(this).attr('data-tab') + id
    $(this).attr('data-tab', tab)
  })
  menu.parent().find("*[data-tab='write']").attr('data-tab', `write${id}`)
  menu.parent().find("*[data-tab='preview']").attr('data-tab', `preview${id}`)
  initCompMarkupContentPreviewTab(menu.parent('form'))
  return id
}

// Will be called when the show more (files) button has been pressed
function onShowMoreFiles() {
  initRepoIssueContentHistory() // TODO: need double check this
  initViewedCheckboxListenerFor()
  countAndUpdateViewedFiles()
}

export function initRepoDiffShowMore() {
  $('#diff-files, #diff-file-boxes')
    .off('click', '#diff-show-more-files, #diff-show-more-files-stats')
    .on('click', '#diff-show-more-files, #diff-show-more-files-stats', (e) => {
      e.preventDefault()

      if ($(e.target).hasClass('btn-disabled')) {
        return
      }
      $('#diff-show-more-files, #diff-show-more-files-stats').addClass(
        'btn-disabled'
      )

      const url = $('#diff-show-more-files, #diff-show-more-files-stats').data(
        'href'
      )
      $.ajax({
        type: 'GET',
        url
      })
        .done((resp) => {
          if (!resp) {
            $('#diff-show-more-files, #diff-show-more-files-stats').removeClass(
              'btn-disabled'
            )
            return
          }
          $('#diff-too-many-files-stats').remove()
          $('#diff-files').append($(resp).find('#diff-files li'))
          $('#diff-incomplete').replaceWith(
            $(resp).find('#diff-file-boxes').children()
          )
          onShowMoreFiles()
        })
        .fail(() => {
          $('#diff-show-more-files, #diff-show-more-files-stats').removeClass(
            'btn-disabled'
          )
        })
    })
  $(document)
    .off('click', 'a.diff-show-more-button')
    .on('click', 'a.diff-show-more-button', (e) => {
      e.preventDefault()
      const $target = $(e.target)
      if ($target.hasClass('btn-disabled')) {
        return
      }
      $target.addClass('btn-disabled')

      const url = $target.data('href')
      $.ajax({
        type: 'GET',
        url
      })
        .done((resp) => {
          console.log(resp)
          if (!resp) {
            $target.removeClass('btn-disabled')
            return
          }

          $target
            .parent()
            .replaceWith(
              $(resp)
                .find('#diff-file-boxes .diff-file-body .file-body')
                .children()
            )
          onShowMoreFiles()
        })
        .fail(() => {
          $target.removeClass('btn-disabled')
        })
    })
}
