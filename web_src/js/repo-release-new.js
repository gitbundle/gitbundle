import $ from './features/jquery.js'
import {
  initRepoRelease,
  initRepoReleaseEditor
} from './modules/repo-release.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoRelease()
  initRepoReleaseEditor()
}

_init('[jq-repository-new-release]')
