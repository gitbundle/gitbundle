import PullRequestMergeForm from '../components/PullRequestMergeForm.svelte'

function initPullRequestMergeForm() {
  const el = document.querySelector('[svelte-pull-request-merge-form]')
  if (!el) return

  // const pageData = window.config.pageData.pullRequestMergeForm
  new PullRequestMergeForm({
    target: el
    // props: { pageData }
  })
}

initPullRequestMergeForm()
