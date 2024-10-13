import { htmlEscape } from 'escape-goat'
import $ from 'jquery'
import { initJqDropdownClose } from '../common/dropdown.js'
import { initJqInputSearch } from '../common/search.js'
import { svg } from '../common/svg.js'

const { appSubUrl, csrfToken } = window.config

export function initRepoIssueTimeTracking() {
  $('[jq-issue-add-time] input')
    .off('keydown')
    .on('keydown', function (e) {
      if ((e.keyCode || e.key) === 13) {
        $(this).parent('[jq-issue-add-time]').trigger('submit')
      }
    })

  //TODO: currently do not have this class in dom
  // $(document).on('click', 'button.issue-delete-time', function () {
  //   const sel = `.issue-delete-time-modal[data-id="${$(this).data('id')}"]`
  //   $(sel)
  //     .modal({
  //       duration: 200,
  //       onApprove() {
  //         $(`${sel} form`).trigger('submit')
  //       }
  //     })
  //     .modal('show')
  // })
}

function updateDeadline(deadlineString) {
  $('[jq-deadline] [jq-invalid-date]').addClass('hidden')
  $('[jq-deadline] [jq-loader]')
    .removeClass('hidden')
    .siblings()
    .addClass('hidden')

  let realDeadline = null
  if (deadlineString !== '') {
    const newDate = Date.parse(deadlineString)

    if (Number.isNaN(newDate)) {
      $('[jq-deadline] [jq-loader]')
        .addClass('hidden')
        .siblings()
        .removeClass('hidden')
      $('[jq-deadline] [jq-invalid-date]').removeClass('hidden')
      return false
    }
    realDeadline = new Date(newDate)
  }

  $.ajax(`${$('[jq-deadline] [jq-form]').attr('action')}`, {
    data: JSON.stringify({
      due_date: realDeadline
    }),
    headers: {
      'X-Csrf-Token': csrfToken
    },
    contentType: 'application/json',
    type: 'POST',
    success() {
      window.location.reload()
    },
    error() {
      $('[jq-deadline] [jq-loader]')
        .addClass('hidden')
        .siblings()
        .removeClass('hidden')
      $('[jq-deadline] [jq-invalid-date]').removeClass('hidden')
    }
  })
}

export function initRepoIssueDue() {
  $('[jq-deadline] [jq-due-edit]').on('click', () =>
    $('[jq-deadline] [jq-form]').fadeToggle(150)
  )
  $('[jq-deadline] [jq-due-remove]').on('click', () => updateDeadline(''))
  // $(document).on('submit', '.issue-due-form', () => {
  //   updateDeadline($('#deadlineDate').val())
  //   return false
  // })
  $('[jq-deadline] [jq-form]').on('submit', function () {
    updateDeadline($(this).find('input[name=deadlineDate]').val())
    return false
  })
}

export function initRepoIssueList() {
  const repolink = $('#repolink').val()
  const repoId = $('#repoId').val()
  const crossRepoSearch = $('#crossRepoSearch').val()
  const tp = $('#type').val()
  let issueSearchUrl = `${appSubUrl}/${repolink}/issues/search?q={query}&type=${tp}`
  if (crossRepoSearch === 'true') {
    issueSearchUrl = `${appSubUrl}/issues/search?q={query}&priority_repo_id=${repoId}&type=${tp}`
  }
  //TODO: this should be double checked @critical
  // $('#new-dependency-drop-list').dropdown({
  //   apiSettings: {
  //     url: issueSearchUrl,
  //     onResponse(response) {
  //       const filteredResponse = { success: true, results: [] }
  //       const currIssueId = $('#new-dependency-drop-list').data('issue-id')
  //       // Parse the response from the api to work with our dropdown
  //       $.each(response, (_i, issue) => {
  //         // Don't list current issue in the dependency list.
  //         if (issue.id === currIssueId) {
  //           return
  //         }
  //         filteredResponse.results.push({
  //           name: `#${issue.number} ${htmlEscape(issue.title)}<div class="text small dont-break-out">${htmlEscape(issue.repository.full_name)}</div>`,
  //           value: issue.id
  //         })
  //       })
  //       return filteredResponse
  //     },
  //     cache: false
  //   },

  //   fullTextSearch: true
  // })

  function excludeLabel(item) {
    const href = $(item).attr('href')
    const id = $(item).data('label-id')

    const regStr = `labels=((?:-?[0-9]+%2c)*)(${id})((?:%2c-?[0-9]+)*)&`
    const newStr = 'labels=$1-$2$3&'

    window.location = href.replace(new RegExp(regStr), newStr)
  }

  $('[jq-label-filter-item]').each(function () {
    $(this).on('click', function (e) {
      console.log(e)
      if (e.altKey) {
        e.preventDefault()
        excludeLabel(this)
      }
    })
  })

  // NOTE: this has been handled by [jq-label-filter-item] @low
  // $('[jq-label-filter]').on('keydown', (e) => {
  //   console.log(e.altKey, e.keyCode)
  //   if (e.altKey && e.keyCode === 13) {
  //     const selectedItems = $('[jq-label-filter] .menu .item.selected') // TODO: this class selector is removed
  //     console.log(selectedItems)
  //     return
  //     if (selectedItems.length > 0) {
  //       excludeLabel($(selectedItems[0]))
  //     }
  //   }
  // })
}

export function initRepoIssueCommentDelete() {
  // Delete comment
  $(document)
    .off('click', '[jq-delete-comment]')
    .on('click', '[jq-delete-comment]', function () {
      const $this = $(this)
      if (window.confirm($this.data('locale'))) {
        $.post($this.data('url'), {
          _csrf: csrfToken
        }).done(() => {
          const $conversationHolder = $this.closest('[jq-conversation-holder]')

          // Check if this was a pending comment.
          if ($conversationHolder.find('[jq-pending-label]').length) {
            const $counter = $('[jq-review-box] [jq-review-comments-counter]')
            let num =
              parseInt($counter.attr('data-pending-comment-number')) - 1 || 0
            num = Math.max(num, 0)
            $counter.attr('data-pending-comment-number', num)
            $counter.text(num)
          }

          $(`#${$this.data('comment-id')}`).remove()
          if (
            $conversationHolder.length &&
            !$conversationHolder.find('[jq-comment]').length
          ) {
            const path = $conversationHolder.data('path')
            const side = $conversationHolder.data('side')
            const idx = $conversationHolder.data('idx')
            const lineType = $conversationHolder.closest('tr').data('line-type')
            if (lineType === 'same') {
              $(
                `[data-path="${path}"] a.add-code-comment[data-idx="${idx}"]`
              ).removeClass('invisible')
            } else {
              $(
                `[data-path="${path}"] a.add-code-comment[data-side="${side}"][data-idx="${idx}"]`
              ).removeClass('invisible')
            }
            $conversationHolder.remove()
          }
        })
      }
      return false
    })
}

export function initRepoPullRequestReview() {
  $(document)
    .off('click', '[jq-show-outdated]')
    .on('click', '[jq-show-outdated]', function (e) {
      e.preventDefault()
      const id = $(this).data('comment')
      if (!id) return
      $(this).addClass('hidden')
      $(`[jq-code-comments-${id}]`).removeClass('hidden')
      $(`[jq-code-preview-${id}]`).removeClass('hidden')
      $(`[jq-hide-outdated-${id}]`).removeClass('hidden')
    })

  $(document)
    .off('click', '[jq-hide-outdated]')
    .on('click', '[jq-hide-outdated]', function (e) {
      e.preventDefault()
      const id = $(this).data('comment')
      if (!id) return
      $(this).addClass('hidden')
      $(`[jq-code-comments-${id}]`).addClass('hidden')
      $(`[jq-code-preview-${id}]`).addClass('hidden')
      $(`[jq-show-outdated-${id}]`).removeClass('hidden')
    })
}

export function initRepoPullRequestUpdate() {
  initJqDropdownClose('[jq-update-buttons]')

  // Pull Request update button
  $('[jq-update-buttons] [jq-button]')
    .off('click')
    .on('click', function (e) {
      e.preventDefault()
      const $this = $(this)
      $this.addClass('loading')
      const redirect = $this.data('redirect')
      $.post($this.data('do'), {
        _csrf: csrfToken
      }).done((data) => {
        if (data.redirect) {
          window.location.href = data.redirect
        } else if (redirect) {
          window.location.href = redirect
        } else {
          window.location.reload()
        }
      })
    })

  $('[jq-update-buttons] [jq-menu-item]')
    .off('click')
    .on('click', function (e) {
      const $url = $(this).data('do')
      const $pullUpdateButton = $(this)
        .closest('[jq-update-buttons]')
        .find('[jq-button]')
      if ($url) {
        $pullUpdateButton.text($(this).text())
        $pullUpdateButton.data('do', $url) //NOTE: use `attr` will made thing complicated
      }
      $(this).closest('[jq-update-buttons]').removeClass('dropdown-open')
      $(this)
        .addClass('active')
        .closest('li')
        .siblings()
        .find('[jq-menu-item]')
        .removeClass('active')
    })
}

export function initRepoPullRequestMergeInstruction() {
  $('.show-instruction').on('click', () => {
    $('[jq-instruct-content]').toggle()
  })
}

export function initRepoPullRequestAllowMaintainerEdit() {
  const $el = $('[jq-allow-edits-from-maintainers]')
  if (!$el.length) return
  $el.attr('data-variation', 'inverted tiny')

  const $checkbox = $el.children(':checkbox')
  const promptTip = $el.attr('data-prompt-tip')
  const promptError = $el.attr('data-prompt-error')
  $el.popup({ content: promptTip })
  $el.off('change').on('change', (e) => {
    const checked = $checkbox.prop('checked')
    let url = $el.attr('data-url')
    url += '/set_allow_maintainer_edit'
    $checkbox.attr('disabled', true)
    $.ajax({
      url,
      type: 'POST',
      data: { _csrf: csrfToken, allow_maintainer_edit: checked },
      error: () => {
        $el.popup({
          content: promptError,
          onHidden: () => {
            // the error popup should be shown only once, then we restore the popup to the default message
            $el.popup({ content: promptTip })
          }
        })
        $el.popup('show')
      },
      complete: () => {
        $checkbox.removeAttr('disabled')
      }
    })
  })
}

export function initRepoIssueWipTitle() {
  $('.title_wip_desc > a').on('click', (e) => {
    e.preventDefault()

    const $issueTitle = $('#issue_title')
    $issueTitle.focus()
    const value = $issueTitle.val().trim().toUpperCase()

    const wipPrefixes = $('.title_wip_desc').data('wip-prefixes')
    for (const prefix of wipPrefixes) {
      if (value.startsWith(prefix.toUpperCase())) {
        return
      }
    }

    $issueTitle.val(`${wipPrefixes[0]} ${$issueTitle.val()}`)
  })
}

export async function updateIssuesMeta(url, action, issueIds, elementId) {
  return $.ajax({
    type: 'POST',
    url,
    data: {
      _csrf: csrfToken,
      action,
      issue_ids: issueIds,
      id: elementId
    }
  })
}

function initRepoIssueComments() {
  // if ($('.repository.view.issue .timeline').length === 0) return

  // NOTE: this exists because we dont want the browser record this form action
  $(document).on('click', (event) => {
    const urlTarget = $(':target')
    if (urlTarget.length === 0) return

    const urlTargetId = urlTarget.attr('id')
    if (!urlTargetId) return
    if (!/^(issue|pull)(comment)?-\d+$/.test(urlTargetId)) return

    const $target = $(event.target)

    if ($target.closest(`#${urlTargetId}`).length === 0) {
      const scrollPosition = $(window).scrollTop()
      window.location.hash = ''
      $(window).scrollTop(scrollPosition)
      window.history.pushState(null, null, ' ')
    }
  })
}

export function initRepoIssueReferenceIssue() {
  // Reference issue
  $(document)
    .off('click', '[jq-reference-issue]')
    .on('click', '[jq-reference-issue]', function (event) {
      const $this = $(this)
      // $this.closest('.dropdown').find('.menu').toggle('visible')

      const content = $(`#comment-${$this.data('target')}`).text()
      const poster = $this.data('poster-username')
      const reference = $this.data('reference')
      const $modal = $($this.data('modal'))
      $modal
        .parent()
        .find('textarea[name="content"]')
        .val(`${content}\n\n_Originally posted by @${poster} in ${reference}_`)
      $modal
        .parent()
        .find('[jq-cancel]')
        .one('click', function () {
          $modal.removeAttr('checked')
        })
      $modal.attr('checked', true)

      event.preventDefault()
    })
}

export function initRepoIssueWipToggle() {
  // Toggle WIP
  //'.toggle-wip a, .toggle-wip button'
  $('[jq-toggle-wip]')
    .off('click')
    .on('click', async function (e) {
      e.preventDefault()
      const title = $(this).data('title')
      const wipPrefix = $(this).data('wip-prefix')
      const updateUrl = $(this).data('update-url')
      await $.post(updateUrl, {
        _csrf: csrfToken,
        title: title?.startsWith(wipPrefix)
          ? title.slice(wipPrefix.length).trim()
          : `${wipPrefix.trim()} ${title}`
      })
      window.location.reload()
    })
}

export function initRepoIssueTitleEdit() {
  // Edit issue title
  const $issueTitle = $('[jq-issue-title]')
  const $editInput = $('[jq-edit-title-input] input')

  const editTitleToggle = function () {
    $issueTitle.toggleClass('hidden')
    $('[jq-not-in-edit]').toggleClass('hidden')
    $('[jq-edit-title-input]').toggleClass('hidden')
    $('[jq-pull-desc]').toggleClass('hidden')
    $('[jq-pull-desc-edit]').toggleClass('hidden')
    $('[jq-in-edit]').toggleClass('hidden')
    $('[jq-issue-title-wrapper]').toggleClass('edit-active')
    $editInput.focus()
    return false
  }

  $('[jq-edit-title]').off('click').on('click', editTitleToggle)
  $('[jq-cancel-edit-title]').off('click').on('click', editTitleToggle)
  $('[jq-save-edit-title]')
    .off('click')
    .on('click', editTitleToggle)
    .on('click', function () {
      const pullrequest_targetbranch_change = function (update_url) {
        const targetBranch = $('[jq-pull-target-branch]').data('branch')
        const $branchTarget = $('[jq-branch-target]')
        if (targetBranch === $branchTarget.text()) {
          return false
        }
        $.post(update_url, {
          _csrf: csrfToken,
          target_branch: targetBranch
        })
          .done((data) => {
            $branchTarget.text(data.base_branch)
          })
          .always(() => {
            window.location.reload()
          })
      }

      const pullrequest_target_update_url = $(this).data('target-update-url')
      if (
        $editInput.val().length === 0 ||
        $editInput.val() === $issueTitle.text()
      ) {
        $editInput.val($issueTitle.text())
        pullrequest_targetbranch_change(pullrequest_target_update_url)
      } else {
        $.post(
          $(this).data('update-url'),
          {
            _csrf: csrfToken,
            title: $editInput.val()
          },
          (data) => {
            $editInput.val(data.title)
            $issueTitle.text(data.title)
            pullrequest_targetbranch_change(pullrequest_target_update_url)
            window.location.reload()
          }
        )
      }
      return false
    })
}

export function initRepoIssueBranchSelect() {
  const changeBranchSelect = function () {
    const selectionTextField = $('[jq-pull-target-branch]')

    const baseName = selectionTextField.data('basename')
    const branchNameNew = $(this).data('branch')
    const branchNameOld = selectionTextField.data('branch')

    // Replace branch name to keep translation from HTML template
    selectionTextField.html(
      selectionTextField
        .html()
        .replace(`${baseName}:${branchNameOld}`, `${baseName}:${branchNameNew}`)
    )
    selectionTextField.data('branch', branchNameNew) // update branch name in setting
    $(this).addClass('checked').siblings().removeClass('checked')
  }
  $('[jq-branch-select] > li').off('click').on('click', changeBranchSelect)
}

// List submits
export function initListSubmits(selector, listSelector) {
  const $list = $(`[${selector}] [${listSelector}]`)
  const $noSelect = $list.find('[jq-no-select]')
  const $menu = $(`[${selector}] [jq-select-menu]`)
  const $items = $(`[${selector}] [jq-select-item]`)
  let hasUpdateAction = $menu.data('action') === 'update'
  const items = {}

  const onClose = () => {
    hasUpdateAction = $menu.data('action') === 'update' // Update the var
    if (hasUpdateAction) {
      // TODO: Add batch functionality and make this 1 network request.
      ;(async function () {
        for (const [elementId, item] of Object.entries(items)) {
          await updateIssuesMeta(
            item['update-url'],
            item.action,
            item['issue-id'],
            elementId
          )
        }
        window.location.reload()
      })()
    }
  }

  initJqDropdownClose(`[${selector}] [jq-dropdown]`, onClose)
  const clearSearch = initJqInputSearch(
    `[${selector}] [jq-input-search]`,
    `[${selector}] [jq-select-item]`
  )

  $items.off('click').on('click', function (e) {
    e.preventDefault()
    // NOTE: class ban-change => attr jq-ban-change
    if ($(this).is('[jq-ban-change]')) {
      return false
    }

    hasUpdateAction = $menu.data('action') === 'update' // Update the var
    if ($(this).hasClass('checked')) {
      $(this).removeClass('checked')
      $(this).find('[jq-octicon-check]').addClass('invisible')
      if (hasUpdateAction) {
        if (!($(this).data('id') in items)) {
          items[$(this).data('id')] = {
            'update-url': $menu.data('update-url'),
            action: 'detach',
            'issue-id': $menu.data('issue-id')
          }
        } else {
          delete items[$(this).data('id')]
        }
      }
    } else {
      $(this).addClass('checked')
      $(this).find('[jq-octicon-check]').removeClass('invisible')
      if (hasUpdateAction) {
        if (!($(this).data('id') in items)) {
          items[$(this).data('id')] = {
            'update-url': $menu.data('update-url'),
            action: 'attach',
            'issue-id': $menu.data('issue-id')
          }
        } else {
          delete items[$(this).data('id')]
        }
      }
    }

    // TODO: Which thing should be done for choosing review requests
    // to make chosen items be shown on time here?
    if (
      selector === 'jq-select-reviewers-modify' ||
      selector === 'jq-select-assignees-modify'
    ) {
      return false
    }

    const listIds = []
    $items.each((_, el) => {
      if ($(el).hasClass('checked')) {
        listIds.push($(el).data('id'))
        $list
          .children(`[jq-${$(el).data('id-selector')}]`)
          .removeClass('hidden')
      } else {
        $list.children(`[jq-${$(el).data('id-selector')}]`).addClass('hidden')
      }
    })
    if (listIds.length === 0) {
      $noSelect.removeClass('hidden')
    } else {
      $noSelect.addClass('hidden')
    }
    $(`[${selector}] input[name=${$menu.data('input-name')}]`).val(
      listIds.join(',')
    )
    return false
  })
  $(`[${selector}] [jq-no-select-item]`)
    .off('click')
    .on('click', function (e) {
      e.preventDefault()
      if (hasUpdateAction) {
        updateIssuesMeta(
          $menu.data('update-url'),
          'clear',
          $menu.data('issue-id'),
          ''
        ).then(() => window.location.reload())
      }

      $items.each(function () {
        $(this).removeClass('checked').removeClass('!hidden')
        $(this).find('[jq-octicon-check]').addClass('invisible')
      })
      clearSearch()

      if (
        selector === 'select-reviewers-modify' ||
        selector === 'select-assignees-modify'
      ) {
        return false
      }

      $list
        .children(':not([jq-no-select])')
        .each((_, el) => $(el).addClass('hidden'))
      $noSelect.removeClass('hidden')
      $(`[${selector}] input[name=${$menu.data('input-name')}]`).val('')
    })
}

export function initSelectItem(selector, listSelector) {
  initJqDropdownClose(`[${selector}] [jq-dropdown]`, null)
  const clearSearch = initJqInputSearch(
    `[${selector}] [jq-input-search]`,
    `[${selector}] [jq-select-item]`
  )

  const $menu = $(`[${selector}] [jq-select-menu]`)
  const $items = $(`[${selector}] [jq-select-item]`)
  const $list = $(`[${selector}] [${listSelector}]`)
  const hasUpdateAction = $menu.data('action') === 'update'

  $items.off('click').on('click', async function (e) {
    e.preventDefault()
    $items.each((_, el) => $(el).removeClass('checked'))
    $(this).addClass('checked')

    if (hasUpdateAction) {
      updateIssuesMeta(
        $menu.data('update-url'),
        '',
        $menu.data('issue-id'),
        $(this).data('id')
      ).then(() => window.location.reload())
    }

    let icon = ''
    if (selector === 'jq-select-milestone') {
      icon = await svg('octicon-milestone', 18)
    } else if (selector === 'jq-select-project') {
      icon = await svg('octicon-project', 18)
    }
    //  else if (selector === 'jq-select-assignee') {
    //   //TODO: currently `jq-select-item` does not have `data-avatar`
    //   icon = `<img class="ui avatar image mr-3" src=${$(this).data('avatar')}>`
    // }

    $list.find('[jq-selected]').html(`
        <a class="flex items-center gap-x-1" href=${$(this).data('href')}>
          ${icon}
          ${htmlEscape($(this).text())}
        </a>
      `)

    $list.children(`[jq-no-select]`).addClass('hidden')
    $(`[${selector}] input[name=${$menu.data('input-name')}]`).val(
      $(this).data('id')
    )
  })
  $(`[${selector}] [jq-no-select-item]`)
    .off('click')
    .on('click', function (e) {
      e.preventDefault()
      if (hasUpdateAction) {
        updateIssuesMeta(
          $menu.data('update-url'),
          '',
          $menu.data('issue-id'),
          $(this).data('id')
        ).then(() => window.location.reload())
      }

      $items.each((_, el) => $(el).removeClass('checked'))
      clearSearch()

      $list.find('[jq-selected]').html('')
      $list.find('[jq-no-select]').removeClass('hidden')
      $(`[${selector}] input[name=${$menu.data('input-name')}]`).val('')
    })
}

initRepoIssueComments()
