// bootstrap module must be the first one to be imported, it handles webpack lazy-loading and global errors
import './utils/bootstrap.js'

import './features/clipboard.js'
import './features/jquery.js'

import './markup/content.js'

import './modules/admin-common.js'
import './modules/common-global.js'
import './modules/dropzone.js'
import './modules/easymde.js'
import './modules/tablesort.js'

import './modules/dashboard-repo-list.js'
import './modules/global-async-button.js'
import './modules/global-navbar.js'
import './modules/global-simple-alert.js'
import './modules/global-simple-modal.js'
import './modules/heatmap.js'
import './modules/pull-request-merge-form.js'
import './modules/repo-branch-tag-compare.js'
import './modules/repo-branch-tag-selector.js'
import './modules/repo-issue.js'

import './modules/serviceworker.js'
