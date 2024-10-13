import { initNewRepo } from './modules/repo-legacy.js'
import { initRepoMigration } from './modules/repo-migration.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoMigration()
  initNewRepo(selector)
}

_init('[jq-repo-new-migrate]')
