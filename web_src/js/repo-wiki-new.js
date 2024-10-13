import $ from './features/jquery.js'
import { initRepoWikiForm } from './modules/repo-wiki.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoWikiForm(selector)
}

_init('[jq-repository-wiki-new]')
