import $ from 'jquery'
import { initJqDropdownClose } from './common/dropdown.js'
import { initJqInputSearch } from './common/search.js'

const { csrfToken } = window.config

function _init(selector) {
  if (!$(`[${selector}]`).length) return

  initJqDropdownClose(`[${selector}] [jq-dropdown]`, null)

  const _ = initJqInputSearch(
    `[${selector}] [jq-input-search]`,
    `[${selector}] [jq-select-item]`
  )

  const $tabs = $(`[${selector}] [jq-tabs]`)
  const $menus = $(`[${selector}] [jq-menu]`)
  const $branchName = $(`[${selector}] [jq-branch-name]`)

  $tabs
    .children()
    .off('click')
    .on('click', function () {
      $(this).addClass('tab-active').siblings().removeClass('tab-active')
      $menus
        .children(`[jq-${$(this).data('target')}]`)
        .removeClass('!hidden')
        .siblings()
        .addClass('!hidden')
    })

  $menus.children().each((_, el) => {
    $(el)
      .children()
      .off('click')
      .on('click', function () {
        const selectedValue = $(this).data('id')
        $(`[${selector}] input[name=${$(el).data('input-name')}]`).val(
          selectedValue
        )
        if ($(el).is('[jq-new-issue]')) {
          $branchName.text($(this).data('name'))
          return
        }

        const editMode = $(`[${selector}] input[name=edit_mode]`).val()
        if (editMode === 'true') {
          $.post(
            $(`[${selector}] [jq-issueref-form]`).attr('action'),
            { _csrf: csrfToken, ref: selectedValue },
            () => window.location.reload()
          )
        } else if (editMode === '') {
          $branchName.text(selectedValue)
        }
        $(this).not('[jq-no-checked]').toggleClass('checked')
      })
  })

  // const $selectBranch = $('.ui.select-branch')
  // const $branchMenu = $selectBranch.find('.reference-list-menu')
  // const $isNewIssue = $branchMenu.hasClass('new-issue')
  // $branchMenu.find('.item:not(.no-select)').click(function () {
  //   const selectedValue = $(this).data('id')
  //   const editMode = $('#editing_mode').val()
  //   $($(this).data('id-selector')).val(selectedValue)
  //   if ($isNewIssue) {
  //     $selectBranch.find('.ui .branch-name').text($(this).data('name'))
  //     return
  //   }
  //
  //   if (editMode === 'true') {
  //     const form = $('#update_issueref_form')
  //     $.post(form.attr('action'), { _csrf: csrfToken, ref: selectedValue }, () => window.location.reload())
  //   } else if (editMode === '') {
  //     $selectBranch.find('.ui .branch-name').text(selectedValue)
  //   }
  // })
  // $selectBranch.find('.reference.column').on('click', function () {
  //   $selectBranch.find('.scrolling.reference-list-menu').css('display', 'none')
  //   $selectBranch.find('.reference .text').removeClass('black')
  //   $($(this).data('target')).css('display', 'block')
  //   $(this).find('.text').addClass('black')
  //   return false
  // })
}

_init('jq-branch-selector')
