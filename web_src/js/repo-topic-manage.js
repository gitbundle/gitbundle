import RepoTopicModal from './components/RepoTopicModal.svelte'
import $ from './features/jquery.js'

function initRepoTopic(selector) {
  const $els = $(selector)
  if (!$els.length) return

  const el = $els[0]
  const pageData = window.config.pageData.repoTopics
  new RepoTopicModal({
    target: el,
    props: { pageData }
  })
}

initRepoTopic('[svelte-repo-topics-manager]')
