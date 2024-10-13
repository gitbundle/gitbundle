import { initCompMarkupContentPreviewTab } from './common/MarkupContentPreview.js'
import { initGlobalEnterQuickSubmit } from './modules/common-global.js'

function _init(selector) {
  if (!$(selector).length) return

  initCompMarkupContentPreviewTab(selector)
  initGlobalEnterQuickSubmit()
}

_init('[jq-comment-tab]')
