<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Typeahead from './WrappedTypeahead.svelte'
  import Svg from './Svg.svelte'

  const { appSubUrl, assetUrlPrefix } = window.config

  let params = new URLSearchParams(window.location.search),
    page = Number(params.get('repo-search-page')) ?? 1,
    tab = params.get('repo-search-tab') ?? 'repos',
    archivedFilter = params.get('repo-search-archived') ?? 'unarchived',
    privateFilter = params.get('repo-search-private') ?? 'both',
    reposFilter = params.get('repo-search-filter') ?? 'all',
    searchQuery = params.get('repo-search-query') ?? '',
    repos = [],
    reposTotalCount = 0,
    finalPage = 1,
    isLoading = false,
    staticPrefix = assetUrlPrefix,
    counts = {},
    repoTypes = {
      all: {
        searchMode: ''
      },
      forks: {
        searchMode: 'fork'
      },
      mirrors: {
        searchMode: 'mirror'
      },
      sources: {
        searchMode: 'source'
      },
      collaborative: {
        searchMode: 'collaborative'
      }
    },
    textArchivedFilterTitles = {},
    textPrivateFilterTitles = {},
    subUrl = appSubUrl,
    pageData = window.config?.pageData?.dashboardRepoList ?? {},
    pageItems

  $: {
    pageItems = []
    for (let i = page; i <= finalPage; i++) {
      pageItems.push(i)
    }
  }

  $: repoTypeCount = counts[`${reposFilter}:${archivedFilter}:${privateFilter}`] ?? 0
  $: checkboxArchivedFilterTitle = textArchivedFilterTitles[archivedFilter] ?? ''
  $: checkboxArchivedFilterProps = {
    checked: archivedFilter === 'archived',
    indeterminate: archivedFilter === 'both'
  }
  $: checkboxPrivateFilterTitle = textPrivateFilterTitles[privateFilter] ?? ''
  $: checkboxPrivateFilterProps = {
    checked: privateFilter === 'private',
    indeterminate: privateFilter === 'both'
  }

  const showMoreReposLink = () => repos.length > 0 && repos.length < counts[`${reposFilter}:${archivedFilter}:${privateFilter}`]
  // prettier-ignore
  const searchURL = () =>
    `${subUrl}/repo/search?sort=updated&order=desc&uid=${
      pageData.uid
    }&team_id=${pageData.teamId ?? 0}&q=${searchQuery}&page=${page}&limit=${
      pageData.searchLimit
    }&mode=${repoTypes[reposFilter].searchMode}${
      reposFilter !== 'all' ? '&exclusive=1' : ''
    }${archivedFilter === 'archived' ? '&archived=true' : ''}${
      archivedFilter === 'unarchived' ? '&archived=false' : ''
    }${privateFilter === 'private' ? '&is_private=true' : ''}${
      privateFilter === 'public' ? '&is_private=false' : ''
    }`

  onMount(() => {
    changeReposFilter(reposFilter)
    textArchivedFilterTitles = {
      archived: pageData.textShowOnlyArchived,
      unarchived: pageData.textShowOnlyUnarchived,
      both: pageData.textShowBothArchivedUnarchived
    }

    textPrivateFilterTitles = {
      private: pageData.textShowOnlyPrivate,
      public: pageData.textShowOnlyPublic,
      both: pageData.textShowBothPrivatePublic
    }
  })

  const changeTab = (t) => {
    tab = t
    updateHistory()
  }

  const changeReposFilter = async (filter) => {
    reposFilter = filter
    repos = []
    page = 1
    counts[`${filter}:${archivedFilter}:${privateFilter}`] = 0
    await searchRepos()
  }

  const updateHistory = () => {
    const params = new URLSearchParams(window.location.search)

    if (tab === 'repos') {
      params.delete('repo-search-tab')
    } else {
      params.set('repo-search-tab', tab)
    }

    if (reposFilter === 'all') {
      params.delete('repo-search-filter')
    } else {
      params.set('repo-search-filter', reposFilter)
    }

    if (privateFilter === 'both') {
      params.delete('repo-search-private')
    } else {
      params.set('repo-search-private', privateFilter)
    }

    if (archivedFilter === 'unarchived') {
      params.delete('repo-search-archived')
    } else {
      params.set('repo-search-archived', archivedFilter)
    }

    if (searchQuery === '') {
      params.delete('repo-search-query')
    } else {
      params.set('repo-search-query', searchQuery)
    }

    if (page === 1) {
      params.delete('repo-search-page')
    } else {
      params.set('repo-search-page', `${page}`)
    }

    const queryString = params.toString()
    if (queryString) {
      window.history.replaceState({}, '', `?${queryString}`)
    } else {
      window.history.replaceState({}, '', window.location.pathname)
    }
  }

  const toggleArchivedFilter = async () => {
    if (archivedFilter === 'unarchived') {
      archivedFilter = 'archived'
    } else if (archivedFilter === 'archived') {
      archivedFilter = 'both'
    } else {
      // including both
      archivedFilter = 'unarchived'
    }
    page = 1
    repos = []
    counts[`${reposFilter}:${archivedFilter}:${privateFilter}`] = 0
    await searchRepos()
  }

  const togglePrivateFilter = async () => {
    if (privateFilter === 'both') {
      privateFilter = 'public'
    } else if (privateFilter === 'public') {
      privateFilter = 'private'
    } else {
      // including private
      privateFilter = 'both'
    }
    page = 1
    repos = []
    counts[`${reposFilter}:${archivedFilter}:${privateFilter}`] = 0
    await searchRepos()
  }

  const changePage = async (p) => {
    page = p
    if (page > finalPage) {
      page = finalPage
    }
    if (page < 1) {
      page = 1
      repos = []
    }
    counts[`${reposFilter}:${archivedFilter}:${privateFilter}`] = 0
    await searchRepos()
  }

  const searchRepos = async () => {
    isLoading = true

    const searchedMode = repoTypes[reposFilter].searchMode
    const searchedURL = searchURL()
    const searchedQuery = searchQuery

    let response, json
    try {
      if (!reposTotalCount) {
        const totalCountSearchURL = `${subUrl}/repo/search?count_only=1&uid=${pageData.uid}&team_id=${pageData.teamId ?? 0}&q=&page=1&mode=`
        response = await fetch(totalCountSearchURL)
        reposTotalCount = response.headers.get('X-Total-Count')
      }
      response = await fetch(searchedURL)
      json = await response.json()
    } catch {
      if (searchedURL === searchURL()) {
        isLoading = false
      }
      return
    }

    if (searchedURL === searchURL()) {
      repos.push(...json.data)
      const count = response.headers.get('X-Total-Count')
      if (searchedQuery === '' && searchedMode === '' && archivedFilter === 'both') {
        reposTotalCount = count
      }
      counts[`${reposFilter}:${archivedFilter}:${privateFilter}`] = count
      finalPage = Math.ceil(count / pageData.searchLimit)
      updateHistory()
      isLoading = false
    }
  }

  const repoIcon = (repo) => {
    if (repo.fork) {
      return 'octicon-repo-forked'
    } else if (repo.mirror) {
      return 'octicon-mirror'
    } else if (repo.template) {
      return `octicon-repo-template`
    } else if (repo.private) {
      return 'octicon-lock'
    } else if (repo.internal) {
      return 'octicon-repo'
    }
    return 'octicon-repo'
  }

  const handleType = async (e) => {
    e.preventDefault()
    searchQuery = e.detail ?? ''
    changeReposFilter(reposFilter)
  }

  beforeUpdate(() => {})
</script>

<div>
  {#if 'isOrganization' in pageData && !pageData.isOrganization}
    <div class="tabs tabs-boxed">
      <button class="tab w-1/2" class:tab-active={tab === 'repos'} on:click={() => changeTab('repos')}>
        <span class="whitespace-nowrap">
          {pageData.textRepository}
          <span class:hidden={reposTotalCount <= 0} class="badge badge-primary badge-sm">
            {reposTotalCount}
          </span>
        </span>
      </button>
      <button class="tab w-1/2" class:tab-active={tab === 'organizations'} on:click={() => changeTab('organizations')}>
        <span class="whitespace-nowrap">
          {pageData.textOrganization}
          <span class:hidden={pageData.organizationsTotalCount <= 0} class="badge badge-primary badge-sm">
            {pageData.organizationsTotalCount}
          </span>
        </span>
      </button>
    </div>
  {/if}
  <div class:hidden={tab !== 'repos'} class="flex flex-col">
    <div class="p-2" title={pageData.textFilter} class:loading={isLoading}>
      <Typeahead placeholder={pageData.textSearchRepos} {handleType} handleClear={handleType} value={searchQuery} className="!px-8 !rounded-lg !input">
        <span slot="typeahead-left">
          <Svg name="octicon-search" size="16" className="mx-2 my-4 h-4 w-4 text-base-content/60" />
        </span>
        <div slot="typeahead-right" class="dropdown dropdown-bottom dropdown-end mx-1 my-3" title={pageData.textFilter} class:loading={isLoading}>
          <button tabindex="0" class="btn-square btn-xs btn text-base-content/60">
            <Svg name="octicon-filter" size={16} />
          </button>
          <ul class="dropdown-content menu rounded-box whitespace-nowrap bg-base-200 p-2 shadow">
            <li>
              <button title={checkboxArchivedFilterTitle} on:click={() => toggleArchivedFilter()}>
                <label class="flex items-center space-x-2">
                  <input type="checkbox" class="checkbox checkbox-sm" bind:checked={checkboxArchivedFilterProps.checked} bind:indeterminate={checkboxArchivedFilterProps.indeterminate} />
                  <Svg name="octicon-archive" size={16} />
                  {pageData.textShowArchived}
                </label>
              </button>
            </li>
            <li>
              <button title={checkboxPrivateFilterTitle} on:click={() => togglePrivateFilter()}>
                <label class="flex items-center space-x-2">
                  <input type="checkbox" class="checkbox checkbox-sm" bind:checked={checkboxPrivateFilterProps.checked} bind:indeterminate={checkboxPrivateFilterProps.indeterminate} />
                  <Svg name="octicon-lock" size={16} />
                  {pageData.textShowPrivate}
                </label>
              </button>
            </li>
          </ul>
        </div>
      </Typeahead>
    </div>
    {#if !isLoading}
      <!-- <Loading />
    {:else} -->
      <div class="tabs w-full flex-nowrap overflow-x-auto">
        <button class="cu-tab-lifted cu-tab px-1 text-xs" class:tab-active={reposFilter === 'all'} on:click={() => changeReposFilter('all')}>
          <span class="whitespace-nowrap">
            {pageData.textAll}
            <span class:hidden={reposFilter !== 'all'} class="badge badge-primary badge-xs">
              {repoTypeCount}
            </span>
          </span>
        </button>
        <button class="cu-tab-lifted cu-tab px-1 text-xs" class:tab-active={reposFilter === 'sources'} on:click={() => changeReposFilter('sources')}>
          <span class="whitespace-nowrap">
            {pageData.textSources}
            <span class:hidden={reposFilter !== 'sources'} class="badge badge-primary badge-xs">
              {repoTypeCount}
            </span>
          </span>
        </button>
        <button class="cu-tab-lifted cu-tab px-1 text-xs" class:tab-active={reposFilter === 'forks'} on:click={() => changeReposFilter('forks')}>
          <span class="whitespace-nowrap">
            {pageData.textForks}
            <span class:hidden={reposFilter !== 'forks'} class="badge badge-primary badge-xs">
              {repoTypeCount}
            </span>
          </span>
        </button>
        {#if !pageData.isMirrorsEnabled}
          <button class="cu-tab-lifted cu-tab px-1 text-xs" class:tab-active={reposFilter === 'mirrors'} on:click={() => changeReposFilter('mirrors')}>
            <span class="whitespace-nowrap">
              {pageData.textMirrors}
              <span class:hidden={reposFilter !== 'mirrors'} class="badge badge-primary badge-xs">
                {repoTypeCount}
              </span>
            </span>
          </button>
        {/if}
        <button class="cu-tab-lifted cu-tab px-1 text-xs" class:tab-active={reposFilter === 'collaborative'} on:click={() => changeReposFilter('collaborative')}>
          <span class="whitespace-nowrap">
            {pageData.textCollaborative}
            <span class:hidden={reposFilter !== 'collaborative'} class="badge badge-primary badge-xs">
              {repoTypeCount}
            </span>
          </span>
        </button>
      </div>
      {#if repos.length}
        <ul class="divide-y">
          {#each repos as repo}
            <li class="cu-links-primary-content flex px-2 py-4 leading-8 hover:bg-base-content/10" class:private={repo.private || repo.internal}>
              <div class="flex flex-col gap-2">
                <div class="flex flex-1">
                  <a class="flex flex-1 items-center gap-1" href={repo.html_url}>
                    {#if repo.avatar_url !== ''}
                      <img class="h-5 w-5 rounded-full" src={repo.avatar_url} alt={repo.full_name} />
                    {:else}
                      <Svg name={repoIcon(repo)} size={16} />
                    {/if}
                    <span class="line-clamp-1">{repo.full_name}</span>
                    {#if repo.archived}
                      <Svg name="octicon-archive" size={16} />
                    {/if}
                  </a>
                </div>
                {#if repo.description !== ''}
                  <p class="line-clamp-3 text-sm text-base-content/80">
                    {repo.description}
                  </p>
                {/if}
                <div class="flex items-center gap-4 text-sm text-base-content/80">
                  {#if pageData.isStarsEnabled}
                    <span class="flex items-center gap-1">
                      <Svg name="octicon-star" size={16} />
                      <span> {repo.stars_count}</span>
                    </span>
                  {/if}
                  {#if repo.language_stats}
                    <div class="flex items-center gap-1">
                      <span class="badge badge-xs" style={`background-color: ${repo.language_stats[0].color}`} />
                      {repo.language_stats[0].language}
                    </div>
                  {/if}
                </div>
              </div>
            </li>
          {/each}
        </ul>
        {#if showMoreReposLink()}
          <button class="btn-outline btn-ghost btn-block btn-sm btn mt-2" class:disabled={page >= finalPage} on:click={() => changePage(page + 1)}>
            <span class="capitalize">{pageData.textLoadMore}</span>
            <Svg name="octicon-arrow-down" size={16} />
          </button>
        {/if}
      {/if}
    {/if}
  </div>
  {#if !pageData.isOrganization}
    <div class:hidden={tab !== 'organizations'} class="flex flex-col">
      {#if pageData.organizations}
        <ul class="flex flex-col divide-y">
          {#each pageData.organizations as org}
            <li class="cu-links-primary-content flex justify-between px-2 py-4 leading-8 hover:bg-base-content/10" key={org.name}>
              <a class="flex flex-1 items-center gap-1" href={subUrl + '/' + encodeURIComponent(org.name)}>
                <Svg name="octicon-organization" size={16} />
                <strong class="line-clamp-1">
                  {org.name}
                </strong>
              </a>
              <span class="flex items-center gap-1">
                {org.num_repos}
                <Svg name="octicon-repo" size={16} />
              </span>
            </li>
          {/each}
        </ul>
      {/if}
    </div>
  {/if}
</div>
