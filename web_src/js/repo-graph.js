import RepoGraphDropdown from './components/RepoGraphDropdown.svelte'

function initRepoGraphDropdown() {
  const el = document.querySelector('[svelte-repo-graph-dropdown]')
  if (!el) return
  const pageData = window.config.pageData.repoGraph
  new RepoGraphDropdown({
    target: el,
    props: { pageData }
  })
}

function _init(selector) {
  if (!$(selector).length) return
  initRepoGraphDropdown()
}

_init('[jq-repo-graph]')
