import { initRepoMigrationStatusChecker } from './modules/repo-migrate.js'

function _init(selector) {
  if (!$(selector).length) return

  initRepoMigrationStatusChecker()
}

_init('[jq-repo-migrating]')
