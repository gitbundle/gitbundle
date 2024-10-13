import $ from './features/jquery.js'
import initImageDiff from './modules/imagediff.js'

function _init(selector) {
  if (!$(selector).length) return

  $(document)
    .off('click', '.image-diff .tabs .tab')
    .on('click', '.image-diff .tabs .tab', function () {
      $(this).addClass('tab-active').siblings().removeClass('tab-active')

      // console.log($(this).data('tab'))
      $(this)
        .closest('.image-diff')
        .find('[jq-contents]')
        .find(`[data-tab=${$(this).data('tab')}]`)
        .removeClass('hidden')
        .siblings()
        .addClass('hidden')
    })

  initImageDiff()
}

_init('[jq-repo-diff-image_diff]')
