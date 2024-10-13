import DashboardRepoList from '../components/DashboardRepoList.svelte'

function initDashboardRepoList() {
  const el = document.getElementById('dashboard-repo-list')
  if (!el) return

  el.innerHTML = ''
  new DashboardRepoList({
    target: el
  })
}

initDashboardRepoList()
