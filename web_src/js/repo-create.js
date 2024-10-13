import $ from 'jquery'
import RepoIssueLabelSelect from './components/RepoIssueLabelSelect.svelte'
import { initNewRepo } from './modules/repo-legacy.js'
import { initRepoTemplateSelect } from './modules/repo-template.js'
import { initSimpleSvelteSelect } from './modules/svelte-simple-select.js'

export function initIssueLabelSelect(selector) {
  const $els = $(selector)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData.repoIssueLabel
  new RepoIssueLabelSelect({
    target: el,
    props: { pageData }
  })
}

function _init(selector) {
  if (!$(selector).length) return

  initNewRepo(selector)
  initRepoTemplateSelect(`${selector} [svelte-repo-template-select]`)
  initIssueLabelSelect(`${selector} [svelte-repo-issue-label]`)

  initSimpleSvelteSelect(`${selector} [svelte-repo-gitignore]`, 'repoGitignore')
  initSimpleSvelteSelect(`${selector} [svelte-repo-license]`, 'repoLicense')
  initSimpleSvelteSelect(`${selector} [svelte-repo-readmes]`, 'repoReadmes')
  initSimpleSvelteSelect(
    `${selector} [svelte-repo-trust_model]`,
    'repoTrustModel'
  )
}

_init('[jq-repository-new-repo]')
