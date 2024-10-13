import {
  initRepoPullRequestMergeInstruction,
  initRepoPullRequestUpdate
} from './modules/repo-issue.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoPullRequestUpdate()
  initRepoPullRequestMergeInstruction()
}

_init('[jq-repo-issue-view_content-pull]')
