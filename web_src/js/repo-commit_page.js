import { initJqDropdownClose } from './common/dropdown.js'
import $ from './features/jquery.js'

function _init(selector) {
  if (!$(selector).length) return

  initJqDropdownClose('[jq-dropdown-commit-menu]')

  $(document)
    .off('click', '[jq-cherry-pick]')
    .on('click', '[jq-cherry-pick]', function (e) {
      const $checkbox = $($(this).data('modal'))
      $checkbox.attr('checked', true)

      const $modal = $checkbox.siblings()
      $modal.find('[jq-header]').text($(this).data('modal-cherry-pick-header'))
      $modal.find('[jq-label]').text($(this).data('modal-cherry-pick-content'))
      $modal.find('[jq-submit]').text($(this).data('modal-cherry-pick-submit'))
      $modal
        .find('input[name=cherry-pick-type]')
        .val($(this).data('modal-cherry-pick-type'))
      $modal.removeClass('hidden')

      $modal
        .find('[jq-cancel]')
        .one('click', () => $checkbox.removeAttr('checked'))
    })

  const $checkbox = $('#cherry-pick-modal')
  if ($checkbox.attr('checked') === undefined || !$checkbox.attr('checked')) {
    $checkbox.removeAttr('checked')
    $checkbox.removeProp('checked')
    $checkbox.siblings().addClass('hidden')
  }
}

_init('[jq-repo-commit_page]')
