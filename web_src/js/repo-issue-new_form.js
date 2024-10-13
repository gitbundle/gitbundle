import { initListSubmits, initSelectItem } from './modules/repo-issue.js'

function _init(selector) {
  if (!$(selector).length) return

  // initCommentTextarea('[jq-comment-form]')

  // Init labels and assignees
  initListSubmits('jq-select-label', 'jq-labels')
  initListSubmits('jq-select-assignee', 'jq-assignees')

  // Milestone, Project
  initSelectItem('jq-select-project', 'jq-projects')
  initSelectItem('jq-select-milestone', 'jq-milestones')
}

_init('[jq-comment-form]')
