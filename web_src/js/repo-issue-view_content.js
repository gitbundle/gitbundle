import { createCommentEasyMDE, getAttachedEasyMDE } from './common/EasyMDE.js'
import { initCompImagePaste } from './common/ImagePaste.js'
import $ from './features/jquery.js'

//TODO: need a better solution @low
function initIssueCommentForm(selector) {
  // Change status
  const $forms = $(selector)
  if ($forms.length !== 1) {
    console.error('Found more than one `jq-issue-comment-form`')
    return
  }
  const $statusButton = $forms.find('[jq-status-button]')
  const $textarea = $forms.find('textarea')
  ;(async () => {
    await createCommentEasyMDE($textarea, {
      onChange: function (e) {
        const easyMDE = getAttachedEasyMDE($textarea)
        const value = easyMDE?.value()
        $statusButton.text(
          $statusButton.data(
            value.length === 0 ? 'status' : 'status-and-comment'
          )
        )
      }
    })
    initCompImagePaste($forms)
  })()

  // $(selector).find('#comment-form textarea').data('on-change', )
  // $('#comment-form textarea')
  //   .off('change')
  //   .on('change', function (e) {
  //     const easyMDE = getAttachedEasyMDE(this)
  //     const value = easyMDE?.value() || $(this).val()
  //     $statusButton.text(
  //       $statusButton.data(value.length === 0 ? 'status' : 'status-and-comment')
  //     )
  //   })
  $statusButton.off('click').on('click', () => {
    $forms.find('[jq-status]').val($statusButton.data('status-val'))
    $forms.trigger('submit')
  })
}

initIssueCommentForm('[jq-issue-comment-form]')
