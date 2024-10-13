import { createCommentEasyMDE, getAttachedEasyMDE } from './common/EasyMDE.js'
import { attachTribute } from './common/tribute.js'
import $ from './features/jquery.js'
import { assignMenuAttributes } from './modules/repo-diff.js'

function initRepoPullRequestReview() {
  const issueCommentPrefix = '#issuecomment-'
  if (
    window.location.hash &&
    window.location.hash.startsWith(issueCommentPrefix)
  ) {
    const commentDiv = $(window.location.hash)
    const issueCommentId = window.location.hash.slice(issueCommentPrefix.length)
    if (commentDiv) {
      // get the name of the parent id
      if (commentDiv.closest(`div[jq-code-comments-${issueCommentId}]`)) {
        $(`[jq-show-outdated-${issueCommentId}]`).addClass('hidden')
        $(`[jq-code-comments-${issueCommentId}]`).removeClass('hidden')
        $(`[jq-code-preview-${issueCommentId}]`).removeClass('hidden')
        $(`[jq-hide-outdated-${issueCommentId}]`).removeClass('hidden')
        commentDiv[0].scrollIntoView()
      }
    }
  }

  // $('.btn-review')
  //   .on('click', function (e) {
  //     e.preventDefault()
  //     $(this).closest('.dropdown').find('.menu').toggle('visible')
  //   })
  //   .closest('.dropdown')
  //   .find('.close')
  //   .on('click', function (e) {
  //     e.preventDefault()
  //     $(this).closest('.menu').toggle('visible')
  //   })
}

function initRepoDiffReviewButton() {
  const $reviewBox = $('[jq-review-box]')
  const $counter = $reviewBox.find('[jq-review-comments-counter]')
  $(document)
    .off('click', 'button[name="is_review"]')
    .on('click', 'button[name="is_review"]', (e) => {
      console.log(e)
      const $form = $(e.target).closest('form')
      $form.append('<input type="hidden" name="is_review" value="true">')
      // Watch for the form's submit event.
      $form.on('submit', () => {
        const num =
          parseInt($counter.attr('data-pending-comment-number')) + 1 || 1
        $counter.attr('data-pending-comment-number', num)
        $counter.text(num)
        // TODO: currently do not have a class named `pulse`
        // Force the browser to reflow the DOM. This is to ensure that the browser replay the animation
        // $reviewBox.removeClass('pulse')
        // $reviewBox.width()
        // $reviewBox.addClass('pulse')
      })
    })
}

function _init(selector) {
  if (!$(selector).length) return

  // Cancel inline code comment
  $(document)
    .off('click', '[jq-cancel-code-comment]')
    .on('click', '[jq-cancel-code-comment]', (e) => {
      const form = $(e.currentTarget).closest('form')
      if (form.length > 0 && form.is('[jq-comment-form]')) {
        form
          .addClass('hidden')
          .closest('[jq-comment-code-cloud]')
          .find('button[jq-comment-form-reply]')
          .removeClass('hidden')
      } else {
        form.closest('[jq-comment-code-cloud]').remove()
      }
    })

  $(document)
    .off('click', 'button[jq-comment-form-reply]')
    .on('click', 'button[jq-comment-form-reply]', async function (e) {
      e.preventDefault()

      $(this).addClass('hidden')
      const form = $(this)
        .closest('[jq-comment-code-cloud]')
        .find('[jq-comment-form]')
      form.removeClass('hidden')
      const $textarea = form.find('textarea')
      let easyMDE = getAttachedEasyMDE($textarea)
      if (!easyMDE) {
        await attachTribute($textarea.get(), { mentions: true, emoji: true })
        easyMDE = await createCommentEasyMDE($textarea, {
          minHeight: '80px'
        })
      }
      $textarea.focus()
      easyMDE.codemirror.focus()
      assignMenuAttributes(form.find('[jq-tabs]'))
    })

  initRepoDiffReviewButton()
  initRepoPullRequestReview()
}

_init('[jq-repo-diff-comment_form]')
