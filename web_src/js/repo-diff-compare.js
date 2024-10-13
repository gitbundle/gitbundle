import $ from 'jquery'
import { getAttachedEasyMDE } from './common/EasyMDE.js'
import { initUnicodeEscapeButton } from './modules/repo-unicode-escape.js'

function _init() {
  if ($('[jq-compare-pull]').length === 0) {
    return
  }

  // Pull request
  const $repoComparePull = $('[jq-compare-pull]')
  if ($repoComparePull.length > 0) {
    // show pull request form
    $repoComparePull
      .find('[jq-show-form-button]')
      .off('click')
      .on('click', function (e) {
        e.preventDefault()
        // eslint-disable-next-line jquery/no-hide
        $(this).parent().hide()

        const $form = $repoComparePull.find('[jq-pullrequest-form]')
        const easyMDE = getAttachedEasyMDE($form.find('textarea.edit_area'))
        $form.toggleClass('hidden')
        easyMDE.codemirror.refresh()
      })
  }

  initUnicodeEscapeButton()
}

_init()
