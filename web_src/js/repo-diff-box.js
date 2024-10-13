import {
  createCommentEasyMDE,
  validateTextareaNonEmpty
} from './common/EasyMDE.js'
import { attachTribute } from './common/tribute.js'
import $ from './features/jquery.js'
import { invertFileFolding } from './modules/file-fold.js'
import { initViewedCheckboxListenerFor } from './modules/pull-view-file.js'
import { initReactionSelector } from './modules/reaction-selector.js'
import {
  assignMenuAttributes,
  initRepoDiffShowMore
} from './modules/repo-diff.js'
import { initRepoPullRequestReview } from './modules/repo-issue.js'
import { initUnicodeEscapeButton } from './modules/repo-unicode-escape.js'

const { csrfToken } = window.config

function init() {
  $('.diff-container [jq-toggle-diff-stats]')
    .off('click')
    .on('click', function (e) {
      $(this)
        .closest('.diff-container')
        .find('.diff-stats')
        .toggleClass('hidden')
    })

  $(document)
    .off('click', '.fold-file')
    .on('click', '.fold-file', ({ currentTarget }) => {
      invertFileFolding(currentTarget.closest('.file-content'), currentTarget)
    })
}

function initRepoDiffConversationForm() {
  $(document)
    .off('submit', '[jq-conversation-holder] form')
    .on('submit', '[jq-conversation-holder] form', async (e) => {
      e.preventDefault()

      const form = $(e.target)
      const $textArea = form.find('textarea')
      if (!validateTextareaNonEmpty($textArea)) {
        return
      }

      const newConversationHolder = $(
        await $.post(form.attr('action'), form.serialize())
      )
      const { path, side, idx } = newConversationHolder.data()

      // TODO: this need double check for form submit
      // initPopup(newConversationHolder.find('.tooltip'))
      form
        .closest('[jq-conversation-holder]')
        .replaceWith(newConversationHolder)
      if (form.closest('tr').data('line-type') === 'same') {
        $(
          `[data-path="${path}"] a.add-code-comment[data-idx="${idx}"]`
        ).addClass('invisible')
      } else {
        $(
          `[data-path="${path}"] a.add-code-comment[data-side="${side}"][data-idx="${idx}"]`
        ).addClass('invisible')
      }
      // TODO: this need double check for form submit
      // newConversationHolder.find('.dropdown').dropdown()
      initReactionSelector(newConversationHolder)
    })

  $(document)
    .off('click', '[jq-resolve-conversation]')
    .on('click', '[jq-resolve-conversation]', async function (e) {
      e.preventDefault()
      const comment_id = $(this).data('comment-id')
      const origin = $(this).data('origin')
      const action = $(this).data('action')
      const url = $(this).data('update-url')

      const data = await $.post(url, {
        _csrf: csrfToken,
        origin,
        action,
        comment_id
      })

      if ($(this).closest('[jq-conversation-holder]').length) {
        const conversation = $(data)
        $(this).closest('[jq-conversation-holder]').replaceWith(conversation)
        initReactionSelector(conversation)
      } else {
        window.location.reload()
      }
    })
}

function initRepoDiffConversationNav() {
  // Previous/Next code review conversation
  $(document)
    .off('click', '[jq-previous-conversation]')
    .on('click', '[jq-previous-conversation]', (e) => {
      const $conversation = $(e.currentTarget).closest(
        '[jq-comment-code-cloud]'
      )
      const $conversations = $('[jq-comment-code-cloud]:not(.hidden)')
      const index = $conversations.index($conversation)
      const previousIndex = index > 0 ? index - 1 : $conversations.length - 1
      const $previousConversation = $conversations.eq(previousIndex)
      const anchor = $previousConversation
        .find('[jq-comment]')
        .first()
        .attr('id')
      window.location.href = `#${anchor}`
    })
  $(document)
    .off('click', '[jq-next-conversation]')
    .on('click', '[jq-next-conversation]', (e) => {
      const $conversation = $(e.currentTarget).closest(
        '[jq-comment-code-cloud]'
      )
      const $conversations = $('[jq-comment-code-cloud]:not(.hidden)')
      const index = $conversations.index($conversation)
      const nextIndex = index < $conversations.length - 1 ? index + 1 : 0
      const $nextConversation = $conversations.eq(nextIndex)
      const anchor = $nextConversation.find('[jq-comment]').first().attr('id')
      window.location.href = `#${anchor}`
    })
}

function initRepoDiffFileViewToggle() {
  $('[jq-file-view-toggle]')
    .off('click')
    .on('click', function () {
      $(this).addClass('btn-active').siblings().removeClass('btn-active')
      const $target = $($(this).data('toggle-selector'))
      $target.removeClass('hidden').siblings().addClass('hidden')
    })
}

function _init(selector) {
  if (!$(selector).length) return

  $(document)
    .off('click', 'a.add-code-comment')
    .on('click', 'a.add-code-comment', async function (e) {
      if ($(e.target).hasClass('btn-add-single')) return // https://github.com/go-gitea/gitea/issues/4745
      e.preventDefault()

      const isSplit = $(this).closest('.code-diff').hasClass('code-diff-split')
      const side = $(this).data('side')
      const idx = $(this).data('idx')
      const path = $(this).closest('[data-path]').data('path')
      const tr = $(this).closest('tr')
      const lineType = tr.data('line-type')

      let ntr = tr.next()
      if (!ntr.hasClass('add-comment')) {
        ntr = $(`
      <tr class="add-comment" data-line-type="${lineType}">
        ${
          isSplit
            ? `<td class="lines-num"></td>
          <td class="lines-escape"></td>
          <td class="lines-type-marker"></td>
          <td class="add-comment-left"></td>
          <td class="lines-num"></td>
          <td class="lines-escape"></td>
          <td class="lines-type-marker"></td>
          <td class="add-comment-right"></td>`
            : `<td class="lines-num"></td>
          <td class="lines-num"></td>
          <td class="lines-escape"></td>
          <td class="add-comment-left add-comment-right" colspan="2"></td>`
        }
      </tr>`)
        tr.after(ntr)
      }

      const td = ntr.find(`.add-comment-${side}`)
      let commentCloud = td.find('[jq-comment-code-cloud]')
      if (
        commentCloud.length === 0 &&
        !ntr.find('button[name="is_review"]').length
      ) {
        const data = await $.get(
          $(this).closest('[data-new-comment-url]').data('new-comment-url')
        )
        td.html(data)
        commentCloud = td.find('[jq-comment-code-cloud]')
        assignMenuAttributes(commentCloud.find('[jq-tabs]'))
        td.find("input[name='line']").val(idx)
        td.find("input[name='side']").val(
          side === 'left' ? 'previous' : 'proposed'
        )
        td.find("input[name='path']").val(path)
        const $textarea = commentCloud.find('textarea')
        await attachTribute($textarea.get(), { mentions: true, emoji: true })
        const easyMDE = await createCommentEasyMDE($textarea, {
          minHeight: '80px'
        })
        $textarea.focus()
        easyMDE.codemirror.focus()
      }
    })

  init()
  initRepoDiffConversationForm()
  initRepoDiffConversationNav()
  initRepoPullRequestReview()
  initRepoDiffFileViewToggle()
  initUnicodeEscapeButton()
  initViewedCheckboxListenerFor()
  initRepoDiffShowMore()
}

_init('[jq-repo-diff-box]')
