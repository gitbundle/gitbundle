import $ from '../features/jquery.js'

export async function createColorPicker($els) {
  if (!$els || !$els.length) return

  await Promise.all([
    import(/* webpackChunkName: "minicolors" */ '@claviska/jquery-minicolors'),
    import(
      /* webpackChunkName: "minicolors" */ '@claviska/jquery-minicolors/jquery.minicolors.css'
    )
  ])

  $els.minicolors()
}

export async function initCompColorPicker() {
  createColorPicker($('.color-picker'))

  $('.precolors [jq-color]').on('click', function () {
    const color_hex = $(this).data('color-hex')
    $('.color-picker').val(color_hex)
    $('.minicolors-swatch-color').css('background-color', color_hex)
  })
}
