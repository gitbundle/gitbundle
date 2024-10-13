import $ from './features/jquery.js'
import {
  initListSubmits,
  // initCommentTextarea,
  initRepoIssueDue,
  initRepoIssueTimeTracking,
  initRepoIssueWipToggle,
  initRepoPullRequestAllowMaintainerEdit,
  initSelectItem,
  updateIssuesMeta
} from './modules/repo-issue.js'
import { initTypeaheadIssueDependenciesDropdown } from './modules/typeahead-issue-dependencies-dropdown.js'

function _init(selector) {
  if (!$(selector).length) return

  // initCommentTextarea('[jq-comment-form]')

  // Init labels and assignees
  initListSubmits('jq-select-label', 'jq-labels')
  // initListSubmits('jq-select-assignee', 'jq-assignees')
  initListSubmits('jq-select-assignees-modify', 'jq-assignees')
  initListSubmits('jq-select-reviewers-modify', 'jq-assignees')

  // Milestone, Assignee, Project
  initSelectItem('jq-select-project', 'jq-projects')
  initSelectItem('jq-select-milestone', 'jq-milestones')
  // selectItem('.select-assignee', '#assignee_id')

  $('[jq-re-request-review]')
    .off('click')
    .on('click', function (e) {
      e.preventDefault()
      const url = $(this).data('update-url')
      const issueId = $(this).data('issue-id')
      const id = $(this).data('id')
      const isChecked = $(this).is('[jq-checked]')

      updateIssuesMeta(url, isChecked ? 'detach' : 'attach', issueId, id).then(
        () => window.location.reload()
      )
    })

  initRepoIssueWipToggle()
  initRepoIssueTimeTracking()
  initRepoIssueDue()

  initTypeaheadIssueDependenciesDropdown(
    '[svelte-typeahead-issue-dependencies-dropdown]'
  )

  initRepoPullRequestAllowMaintainerEdit()
}

_init('[jq-repo-issue-view_content-sidebar]')
