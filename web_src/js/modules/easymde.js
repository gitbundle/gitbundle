import { createCommentEasyMDE } from '../common/EasyMDE.js'
import { initCompImagePaste } from '../common/ImagePaste.js'
import $ from '../features/jquery.js'

export function initCommentTextarea(formSelector) {
  const $forms = $(formSelector)
  if ($forms.length !== 1) return

  const $textarea = $forms.find('textarea:not(.review-textarea)')
  if ($textarea.length !== 1)
    return // const easyMDE = getAttachedEasyMDE($textarea)
    // console.log($textarea, easyMDE)
    // if (easyMDE) return
  ;(async () => {
    await createCommentEasyMDE($textarea)
    initCompImagePaste($forms)
  })()
}

initCommentTextarea('[jq-comment-form]')
