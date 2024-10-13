import $ from './features/jquery.js'
import { initRepoIssueList, updateIssuesMeta } from './modules/repo-issue.js'

export function initCommonIssue() {
  $('[jq-issue-list] [jq-issue-checkbox]')
    .off('click')
    .on('click', () => {
      const numChecked = $('[jq-issue-list] [jq-issue-checkbox]:checked').length
      if (numChecked > 0) {
        $('[jq-issue-filters]').addClass('hidden')
        $('[jq-issue-actions]').removeClass('hidden')
      } else {
        $('[jq-issue-filters]').removeClass('hidden')
        $('[jq-issue-actions]').addClass('hidden')
      }
    })

  $('[jq-issue-action]')
    .off('click')
    .on('click', async function () {
      let action = $(this).data('action')
      let elementId = $(this).data('element-id')
      const url = $(this).data('url')
      const issueIDs = $('[jq-issue-list] [jq-issue-checkbox]:checked')
        .map((_, el) => {
          return $(el).data('issue-id')
        })
        .get()
        .join(',')
      if (!elementId && url.slice(-9) === '/assignee') {
        elementId = ''
        action = 'clear'
      }
      console.log(url, action, issueIDs, elementId)
      // return
      updateIssuesMeta(url, action, issueIDs, elementId).then(() => {
        // NOTICE: This reset of checkbox state targets Firefox caching behaviour, as the
        // checkboxes stay checked after reload
        if (action === 'close' || action === 'open') {
          // uncheck all checkboxes
          $('[jq-issue-list] [jq-issue-checkbox]').each((_, e) => {
            e.checked = false
          })
        }
        window.location.reload()
      })
    })

  // NOTICE: This event trigger targets Firefox caching behaviour, as the checkboxes stay
  // checked after reload trigger ckecked event, if checkboxes are checked on load
  $('[jq-issue-list] [jq-issue-checkbox]:checked')
    .first()
    .each((_, e) => {
      e.checked = false
      $(e).off('click').trigger('click')
    })
}

function _init(selector) {
  if (!$(selector).length) return

  initCommonIssue()
  initRepoIssueList()
}

_init('[jq-repo-issue-list]')
