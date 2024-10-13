import $ from 'jquery'
import SearchUserStatic from './components/SearchUserStatic.svelte'

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

  initRepoSearchUsers(
    selector,
    '[svelte-repo-search-allowlist_users]',
    'repoSearchAllowlistUsers'
  )

  initRepoSearchUsers(
    selector,
    '[svelte-repo-search-allowlist_teams]',
    'repoSearchAllowlistTeams'
  )
}

_init('[jq-repository-settings-tags]')
