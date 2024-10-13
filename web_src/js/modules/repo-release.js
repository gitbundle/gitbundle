import { createCommentEasyMDE } from '../common/EasyMDE.js'
import { initEasyMDEImagePaste } from '../common/ImagePaste.js'
import { initCompMarkupContentPreviewTab } from '../common/MarkupContentPreview.js'
import $ from '../features/jquery.js'
import attachTribute from './tribute.js'

export function initRepoRelease() {
  $(document)
    .off('click', '[jq-remove-rel-attach]')
    .on('click', '[jq-remove-rel-attach]', function () {
      const uuid = $(this).data('uuid')
      const id = $(this).data('id')
      $(`input[name='attachment-del-${uuid}']`).attr('value', true)
      $(`#attachment-${id}`).hide()
    })
}

export function initRepoReleaseEditor() {
  const $editor = $('[jq-repository-new-release] [jq-content-editor]')
  if ($editor.length === 0) {
    return false
  }

  ;(async () => {
    const $textarea = $editor.find('textarea')
    await attachTribute($textarea.get(), { mentions: false, emoji: true })
    const $files = $editor.parent().find('.files')
    const easyMDE = await createCommentEasyMDE($textarea)
    initCompMarkupContentPreviewTab($editor)
    const dropzone = $editor.parent().find('.dropzone')[0]
    initEasyMDEImagePaste(easyMDE, dropzone, $files)
  })()
}
