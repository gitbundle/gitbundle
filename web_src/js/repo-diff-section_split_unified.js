import $ from './features/jquery.js'

function _init(selector) {
  if (!$(selector).length) return

  $(document)
    .off('click', '.blob-excerpt')
    .on('click', '.blob-excerpt', async ({ currentTarget }) => {
      const url = currentTarget.getAttribute('data-url')
      const query = currentTarget.getAttribute('data-query')
      const anchor = currentTarget.getAttribute('data-anchor')
      if (!url) return
      const blob = await $.get(`${url}?${query}&anchor=${anchor}`)
      currentTarget.closest('tr').outerHTML = blob
    })
}

_init('[jq-repo-diff-section_split_unified]')
