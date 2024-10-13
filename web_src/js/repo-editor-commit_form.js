import $ from './features/jquery.js'

function _init(selector) {
  if (!$(selector).length) return

  // $('.js-quick-pull-choice-option').on('change', function () {
  //   if ($(this).val() === 'commit-to-new-branch') {
  //     $('.quick-pull-branch-name').show()
  //     $('.quick-pull-branch-name input').prop('required', true)
  //   } else {
  //     $('.quick-pull-branch-name').hide()
  //     $('.quick-pull-branch-name input').prop('required', false)
  //   }
  //   $('[jq-commit-button]').text($(this).data('button-text'))
  // })

  $('[jq-quick-pull-choice-option] input')
    .off('change')
    .on('change', function (e) {
      if ($(this).val() === 'commit-to-new-branch') {
        $(this)
          .parent()
          .find('[jq-quick-pull-branch-name]')
          .removeClass('hidden')
        $(this)
          .parent()
          .find('[jq-quick-pull-branch-name] input')
          .prop('required', true)
      } else {
        $(this)
          .closest('[jq-form-control]')
          .siblings('[jq-form-control]')
          .find('[jq-quick-pull-choice-option] [jq-quick-pull-branch-name]')
          .addClass('hidden')
        $(this)
          .closest('[jq-form-control]')
          .siblings('[jq-form-control]')
          .find(
            '[jq-quick-pull-choice-option] [jq-quick-pull-branch-name] input'
          )
          .prop('required', false)
      }
      $('[jq-commit-button]').text($(this).data('button-text'))

      // if ($(this).find(':checked').length) {
      //   $(this).parent().find('[jq-quick-pull-branch-name]').removeClass('hidden')
      //   $(this)
      //     .parent()
      //     .siblings()
      //     .find('[jq-quick-pull-choice-option] [jq-quick-pull-branch-name]')
      //     .addClass('hidden')
      // }
    })
}

_init('[jq-commit-form-wrapper]')
