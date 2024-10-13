import { initJqDropdownClose } from './common/dropdown.js'
import {
  initRepoIssueBranchSelect,
  initRepoIssueTitleEdit
} from './modules/repo-issue.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoIssueTitleEdit()
  initRepoIssueBranchSelect()
  initJqDropdownClose('[jq-branch-select-dropdown]')
}

_init('[jq-repo-issue-view_title]')
