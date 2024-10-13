import { initJqDropdownClose } from './common/dropdown.js'
import {
  initRepoIssueCommentDelete,
  initRepoIssueReferenceIssue
} from './modules/repo-issue.js'
import { initRepoIssueCommentEdit } from './modules/repo-legacy.js'

function _init(selector) {
  if (!$(selector).length) return

  initJqDropdownClose(selector)
  initRepoIssueCommentEdit(selector)
  initRepoIssueReferenceIssue()
  initRepoIssueCommentDelete()
}

_init('[jq-dropdown-context-menu]')
