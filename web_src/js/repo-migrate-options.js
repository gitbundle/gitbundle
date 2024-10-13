import $ from 'jquery'

function _init(selector) {
  if (!$(selector).length) return

  const $lfs = $(selector).find('#lfs')
  const $lfsSettings = $(selector).find('#lfs_settings')
  const $lfsEndpoint = $(selector).find('#lfs_endpoint')

  function setLFSSettingsVisibility() {
    const visible = $lfs.is(':checked')
    $lfsSettings.toggle(visible)
    $lfsEndpoint.addClass('hidden')
  }
  setLFSSettingsVisibility()

  $(`${selector} #lfs_settings_show`).on('click', () => {
    $lfsEndpoint.removeClass('hidden')
    return false
  })
  $lfs.on('change', setLFSSettingsVisibility)
}

_init('[jq-repo-migrate-options]')
