import $ from 'jquery'

$('label[jq-create-branch]').each((_, el) => {
  $(el)
    .off('click')
    .on('click', () => {
      const $form = $('form[jq-create-branch]')
      if ($form.length > 1) {
        console.error('got more than one form with `form[jq-create-branch]`')
        return
      }
      $form.find('span[jq-create-branch-from]').text($(el).data('branch-from'))
      $form.attr(
        'action',
        $form.data('base-action') + $(el).data('branch-from-urlcomponent')
      )
    })
})
