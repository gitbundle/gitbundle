import { initCompColorPicker } from './common/ColorPicker.js'
import $ from './features/jquery.js'
import { initRepoProject } from './modules/repo-projects.js'

function _init(selector) {
  if (!$(selector).length) return

  initCompColorPicker()
  initRepoProject()
}

_init('[jq-repository-view-project]')
