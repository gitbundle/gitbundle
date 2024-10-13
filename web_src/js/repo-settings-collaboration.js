import {
  initRepoSearchBox,
  initRepoSettingsCollaboration
} from './modules/repo-settings.js'
import { initSearchUserBox } from './modules/search-user.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoSettingsCollaboration(selector)
  initRepoSearchBox(`${selector} [svelte-repo-collab-team]`, 'repoCollabTeam')
  initSearchUserBox(`${selector} [svelte-repo-collab]`, 'repoCollab')
}

_init('[jq-repository-settings-collaboration]')
