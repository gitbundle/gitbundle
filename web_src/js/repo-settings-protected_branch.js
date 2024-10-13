import $ from 'jquery'
import SearchUserStatic from './components/SearchUserStatic.svelte'
import { initRepoSettingBranches } from './modules/repo-settings.js'

export function initRepoSearchUsers(selector, container, dataKey) {
  const $els = $(`${selector} ${container}`)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData[dataKey]
  new SearchUserStatic({
    target: el,
    props: { pageData }
  })
}

function _init(selector) {
  if (!$(selector).length) return

  initRepoSettingBranches(selector)
  initRepoSearchUsers(selector, '[svelte-repo-search-users]', 'repoSearchUsers')
  initRepoSearchUsers(selector, '[svelte-repo-search-teams]', 'repoSearchTeams')
  initRepoSearchUsers(
    selector,
    '[svelte-repo-search-merge_whitelist_users]',
    'repoSearchMergeWhitelistUsers'
  )
  initRepoSearchUsers(
    selector,
    '[svelte-repo-search-merge_whitelist_teams]',
    'repoSearchMergeWhitelistTeams'
  )
  initRepoSearchUsers(
    selector,
    '[svelte-repo-search-approvals_whitelist_users]',
    'repoSearchApprovalsWhitelistUsers'
  )
  initRepoSearchUsers(
    selector,
    '[svelte-repo-search-approvals_whitelist_teams]',
    'repoSearchApprovalsWhitelistTeams'
  )
}

_init('[jq-repository-settings-protected-branches]')
