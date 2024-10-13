import RepoWikiSidebarNav from './components/RepoWikiSidebarNav.svelte'
import $ from './features/jquery.js'

function initRepoWikiSidebarNav(selector) {
  const $els = $(`${selector}`)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData.repoWikiSidebarNav
  new RepoWikiSidebarNav({
    target: el,
    props: { pageData }
  })
}

function _init(selector) {
  if (!$(selector).length) return

  initRepoWikiSidebarNav('[svelte-repo-wiki-sidebar-nav]')
}

_init('[jq-repository-wiki-view]')
