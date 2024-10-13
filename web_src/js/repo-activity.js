import RepoActivityTopAuthors from './components/RepoActivityTopAuthors.svelte'
import $ from './features/jquery.js'

function _init(selector) {
  const $topAuthorsCharts = $(
    `${selector} [svelte-repo-activity-top-authors-chart]`
  )
  if (!$topAuthorsCharts.length) return

  const el = $topAuthorsCharts[0]
  const pageData = window.config.pageData.repoActivityTopAuthors
  new RepoActivityTopAuthors({
    target: el,
    props: { pageData }
  })
}

_init('[jq-repository-activity]')
