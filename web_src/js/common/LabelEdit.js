import $ from '../features/jquery.js'
import { initCompColorPicker } from './ColorPicker.js'

export async function initCompLabelEdit(selector) {
  if (!$(selector).length) return
  // Create label
  const $newLabelPanel = $('[jq-new-label-form]')
  $('[jq-new-label-button]').on('click', function () {
    $newLabelPanel.removeClass('hidden')
    $(this).addClass('hidden')
  })
  $('[jq-new-label-form] [jq-cancel]').on('click', () => {
    $newLabelPanel.addClass('hidden')
    $('[jq-new-label-button]').removeClass('hidden')
  })

  await initCompColorPicker()

  $('[jq-edit-label-button]')
    .off('click')
    .on('click', function () {
      $('[jq-edit-label-modal] :checkbox').prop('checked', true)

      $('.edit-label .color-picker').minicolors('value', $(this).data('color'))
      $('#label-modal-id').val($(this).data('id'))
      $('.edit-label .new-label-input').val($(this).data('title'))
      $('.edit-label .new-label-desc-input').val($(this).data('description'))
      $('.edit-label .color-picker').val($(this).data('color'))
      $('.edit-label .minicolors-swatch-color').css(
        'background-color',
        $(this).data('color')
      )
      $('[jq-edit-label-modal] [jq-cancel]').one('click', function () {
        $(this)
          .closest('[jq-edit-label-modal]')
          .children(':checkbox')
          .prop('checked', false)
      })
    })
}
