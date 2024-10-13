import $ from 'jquery'

export function initRepoSettingsOptions(selector) {
  if (!$(selector).length) return

  // Enable or select internal/external wiki system and issue tracker.
  $('[jq-enable-system]').on('change', function () {
    if (this.checked) {
      $($(this).data('target')).removeClass('disabled')
      if (!$(this).data('context'))
        $($(this).data('context')).addClass('disabled')
    } else {
      $($(this).data('target')).addClass('disabled')
      if (!$(this).data('context'))
        $($(this).data('context')).removeClass('disabled')
    }
  })
  $('[jq-enable-system-radio]').on('change', function () {
    if (this.value === 'false') {
      $($(this).data('target')).addClass('disabled')
      if (typeof $(this).data('context') !== 'undefined')
        $($(this).data('context')).removeClass('disabled')
    } else if (this.value === 'true') {
      $($(this).data('target')).removeClass('disabled')
      if (typeof $(this).data('context') !== 'undefined')
        $($(this).data('context')).addClass('disabled')
    }
  })
  const $trackerIssueStyleRadios = $('[jq-tracker-issue-style]')
  $trackerIssueStyleRadios.on('change input', () => {
    const checkedVal = $trackerIssueStyleRadios.filter(':checked').val()
    $('#tracker-issue-style-regex-box').toggleClass(
      'disabled',
      checkedVal !== 'regexp'
    )
  })
}

initRepoSettingsOptions('[jq-repository-settings-options]')
