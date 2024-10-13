import { initRepoPullRequestReview } from './modules/repo-issue.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoPullRequestReview()
}

_init('[jq-repo-issue-view_content-comments]')
