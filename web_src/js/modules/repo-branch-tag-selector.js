import RepoBranchTagSelector from '../components/RepoBranchTagSelector.svelte'

function initRepoBranchTagSelector() {
  const els = document.querySelectorAll('.repo-branch-tag-selector')
  els.forEach((el, index) => {
    const pageData = window.config.pageData.branchDropdownDataList[index]
    new RepoBranchTagSelector({
      target: el,
      props: { pageData }
    })
  })
}

initRepoBranchTagSelector()
