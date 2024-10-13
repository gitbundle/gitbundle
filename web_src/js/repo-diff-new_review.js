import { createCommentEasyMDE, getAttachedEasyMDE } from './common/EasyMDE.js'
import { initCompImagePaste } from './common/ImagePaste.js'
import $ from './features/jquery.js'

function _init(selector) {
  if (!$(selector).length) return

  $(document)
    .off('click', '[jq-review-box] [jq-open-button]')
    .on('click', '[jq-review-box] [jq-open-button]', function () {
      $(this)
        .closest('[jq-review-box]')
        .addClass('dropdown-open')
        .children('[jq-dropdown-content]')
        .removeClass('hidden')
    })

  $(document)
    .off('click', '[jq-review-box] [jq-close-button]')
    .on('click', '[jq-review-box] [jq-close-button]', function () {
      $(this)
        .closest('[jq-review-box]')
        .removeClass('dropdown-open')
        .children('[jq-dropdown-content]')
        .addClass('hidden')
    })

  const $reviewBox = $('[jq-review-box]')
  const $textarea = $reviewBox.find('textarea')
  if ($reviewBox.length === 1 && $textarea.length === 1) {
    if (!getAttachedEasyMDE($textarea)) {
      ;(async () => {
        // the editor's height is too large in some cases, and the panel cannot be scrolled with page now because there is `.repository .diff-detail-box.sticky { position: sticky; }`
        // the temporary solution is to make the editor's height smaller (about 4 lines). GitHub also only show 4 lines for default. We can improve the UI (including Dropzone area) in future
        // EasyMDE's options can not handle minHeight & maxHeight together correctly, we have to set max-height for .CodeMirror-scroll in CSS.
        await createCommentEasyMDE($textarea, {
          minHeight: '80px'
        })
        initCompImagePaste($reviewBox)
      })()
    }
  }
}

_init('[jq-review-box]')
