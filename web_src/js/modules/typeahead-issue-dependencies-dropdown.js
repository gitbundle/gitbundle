import TypeaheadDropdown from '../components/TypeaheadIssueDependenciesDropdown.svelte'

export function initTypeaheadIssueDependenciesDropdown(selector) {
  const els = document.querySelectorAll(selector)
  if (!els.length) return
  if (els.length > 1) {
    console.error(`found more than one container "${selector}"`)
    return
  }
  const container = els[0]

  const pageData = {
    crossRepoSearch: container.getAttribute('data-cross-repo-search'),
    action: container.getAttribute('data-action'),
    placeholder: container.getAttribute('data-placeholder'),
    issueID: container.getAttribute('data-issue-id'),
    repolink: container.getAttribute('data-repolink'),
    repoid: container.getAttribute('data-repoid'),
    issueType: container.getAttribute('data-issue-type'),
    exists: container.getAttribute('data-exists')
  }
  new TypeaheadDropdown({
    target: container,
    props: { pageData }
  })
}
