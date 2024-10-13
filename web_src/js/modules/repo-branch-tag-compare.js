import RepoBranchTagCompare from '../components/RepoBranchTagCompare.svelte'

function initRepoBranchTagCompare(pageData, containerId) {
  const container = document.getElementById(containerId)
  if (!container) return

  const base = !!container.getAttribute('data-base')
  new RepoBranchTagCompare({
    target: container,
    props: { base }
  })
}

initRepoBranchTagCompare(
  window.config.pageData.repoBranchTagCompare,
  'repo-branch-tag-compare-base'
)
initRepoBranchTagCompare(
  window.config.pageData.repoBranchTagCompare,
  'repo-branch-tag-compare-target'
)
