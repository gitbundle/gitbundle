import { renderCodeCopy } from './codecopy.js'
import { renderMermaid } from './mermaid.js'
import { initMarkupTasklist } from './tasklist.js'

// code that runs for all markup content
export function initMarkupContent() {
  renderMermaid()
  renderCodeCopy()
}

// code that only runs for comments
export function initCommentContent() {
  initMarkupTasklist()
}

initMarkupContent()
initCommentContent()
