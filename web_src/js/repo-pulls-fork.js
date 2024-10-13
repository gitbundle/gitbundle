import RepoOwnerSelect from './components/RepoOwnerSelect.svelte'
import $ from './features/jquery.js'

export function _init(selector) {
  if (!$(selector).length) return

  const $els = $(`${selector} [svelte-repo-owner-select]`)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData.repoOwnerSelect
  new RepoOwnerSelect({
    target: el,
    props: { pageData }
  })
}

_init('[jq-repository-new-fork]')
