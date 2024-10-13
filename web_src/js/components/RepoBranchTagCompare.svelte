<script>
  import { beforeUpdate, onMount, tick } from 'svelte'
  import Svg from './Svg.svelte'
  import WrappedTypeahead from './WrappedTypeahead.svelte'
  import { pathEscapeSegments } from '../utils/url.js'

  export let base

  const data = {
    csrfToken: window.config.csrfToken,
    searchTerm: '',
    refNameText: '',

    mode: 'branches',

    isViewTag: false,
    isViewBranch: false,
    isViewTree: false,

    active: false,

    ...window.config.pageData.repoBranchTagCompare
  }

  let comboboxRef,
    searchRef,
    defaultSelectedValue,
    showNoResults = false

  $: if (base) {
    defaultSelectedValue = `${data.baseCompareName}:${data.baseBranch}`
  } else {
    defaultSelectedValue = `${data.headCompareName}:${data.headBranch}`
  }

  const initSearchData = []
  if (data.mode === 'tags') {
    if (base) {
      for (const item of Array.from(data.tags))
        initSearchData.push({
          type: 0,
          content: `${data.baseCompareName}:${item}`
        })
      if (data.notPullRequestCtxSameRepo) {
        for (const item of Array.from(data.headTags))
          initSearchData.push({
            type: 1,
            content: `${data.headCompareName}:${item}`
          })
      }
    } else {
      for (const item of Array.from(data.headTags))
        initSearchData.push({
          type: 0,
          content: `${data.headCompareName}:${item}`
        })
      if (data.notPullRequestCtxSameRepo) {
        for (const item of Array.from(data.tags))
          initSearchData.push({
            type: 1,
            content: `${data.baseCompareName}:${item}`
          })
      }
    }
    if (data.isOwnForkRepo) {
      for (const item of Array.from(data.ownForkRepoTags))
        initSearchData.push({
          type: 2,
          content: `${data.ownForkCompareName}:${item}`
        })
    }
    if (data.isRootRepo) {
      for (const item of Array.from(data.rootRepoTags))
        initSearchData.push({
          type: 3,
          content: `${data.rootRepoCompareName}:${item}`
        })
    }

    // if (!data.isViewTag) {
    //   defaultSelectedValue = null
    // }
  } else if (data.mode === 'branches') {
    if (base) {
      if (data.branches) {
        for (const item of Array.from(data.branches))
          initSearchData.push({
            type: 0,
            content: `${data.baseCompareName}:${item}`
          })
      }
      if (data.notPullRequestCtxSameRepo && data.headBranches) {
        for (const item of Array.from(data.headBranches))
          initSearchData.push({
            type: 1,
            content: `${data.headCompareName}:${item}`
          })
      }
    } else {
      if (data.headBranches) {
        for (const item of Array.from(data.headBranches))
          initSearchData.push({
            type: 0,
            content: `${data.headCompareName}:${item}`
          })
      }
      if (data.notPullRequestCtxSameRepo && data.branches) {
        for (const item of Array.from(data.branches))
          initSearchData.push({
            type: 1,
            content: `${data.baseCompareName}:${item}`
          })
      }
    }
    if (data.isOwnForkRepo && data.ownForkRepoBranches) {
      for (const item of Array.from(data.ownForkRepoBranches))
        initSearchData.push({
          type: 2,
          content: `${data.ownForkCompareName}:${item}`
        })
    }
    if (data.isRootRepo && data.rootRepoBranches) {
      for (const item of Array.from(data.rootRepoBranches))
        initSearchData.push({
          type: 3,
          content: `${data.rootRepoCompareName}:${item}`
        })
    }

    // if (!data.isViewBranch) {
    //   defaultSelectedValue = null
    // }
  }
  // $: if (data.viewType === 'tree') {
  //   data.isViewTree = true
  //   data.refNameText = data.commitIdShort
  // } else if (data.viewType === 'tag') {
  //   data.isViewTag = true
  //   data.refNameText = data.tagName
  // } else {
  //   data.isViewBranch = true
  //   data.refNameText = data.branchName
  // }

  const handleResults = (results) => {
    showNoResults = results.length === 0
  }

  const buildBaseUrl = (baseLink, name, type) => {
    if (type === 0) {
      return `${baseLink}/compare/${pathEscapeSegments(name)}${data.compareSeparator}${data.notPullRequestCtxSameRepo ? encodeURIComponent(data.headUserName) + '/' + encodeURIComponent(data.headRepoName) + ':' : ''}${pathEscapeSegments(data.headBranch)}`
    }

    return `${baseLink}/compare/${pathEscapeSegments(name)}${data.compareSeparator}${encodeURIComponent(data.headUserName)}/${encodeURIComponent(data.headRepoName)}:${pathEscapeSegments(data.headBranch)}`
  }

  const buildTargetUrl = (name, type) => {
    if (type === 0) {
      // "{{$.RepoLink}}/compare/{{PathEscapeSegments $.BaseBranch}}{{$.CompareSeparator}}{{if not $.PullRequestCtx.SameRepo}}{{PathEscape $.HeadUser.Name}}/{{PathEscape $.HeadRepo.Name}}:{{end}}{{PathEscapeSegments .}}"
      return `${data.repoLink}/compare/${pathEscapeSegments(data.baseBranch)}${data.compareSeparator}${data.notPullRequestCtxSameRepo ? encodeURIComponent(data.headUserName) + '/' + encodeURIComponent(data.headRepoName) + ':' : ''}${pathEscapeSegments(name)}`
    } else if (type === 1) {
      // {{$.RepoLink}}/compare/{{PathEscapeSegments $.BaseBranch}}{{$.CompareSeparator}}{{PathEscape $.BaseName}}/{{PathEscape $.Repository.Name}}:{{PathEscapeSegments .}}
      return `${data.repoLink}/compare/${pathEscapeSegments(data.baseBranch)}${data.compareSeparator}${encodeURIComponent(data.baseName)}/${encodeURIComponent(data.repositoryName)}:${pathEscapeSegments(name)}`
    } else if (type === 2) {
      // {{$.RepoLink}}/compare/{{PathEscapeSegments $.BaseBranch}}{{$.CompareSeparator}}{{PathEscape $.OwnForkRepo.OwnerName}}/{{PathEscape $.OwnForkRepo.Name}}:{{PathEscapeSegments .}}
      return `${data.repoLink}/compare/${pathEscapeSegments(data.baseBranch)}${data.compareSeparator}${encodeURIComponent(data.ownForkRepoOwnerName)}/${encodeURIComponent(data.ownForkRepoName)}:${pathEscapeSegments(name)}`
    } else if (type === 3) {
      // {{$.RepoLink}}/compare/{{PathEscapeSegments $.BaseBranch}}{{$.CompareSeparator}}{{PathEscape $.RootRepo.OwnerName}}/{{PathEscape $.RootRepo.Name}}:{{PathEscapeSegments .}}
      return `${data.repoLink}/compare/${pathEscapeSegments(data.baseBranch)}${data.compareSeparator}${encodeURIComponent(data.rootRepoOwnerName)}/${data.rootRepoName}:${pathEscapeSegments(name)}`
    }
  }

  const handleSelect = async ({ detail }) => {
    // console.log(detail)
    const { original } = detail
    const [_, name] = original.content.split(':')
    let url = ''
    if (!base) {
      url = buildTargetUrl(name, original.type)
    } else {
      switch (original.type) {
        case 0:
          url = buildBaseUrl(data.repoLink, name, original.type)
          break
        case 1:
          url = buildBaseUrl(data.headRepoLink, name, original.type)
          break
        case 2:
          url = buildBaseUrl(data.ownForkRepoLink, name, original.type)
          break
        case 3:
          url = buildBaseUrl(data.rootRepoLink, name, original.type)
          break
      }
    }
    // console.log(url)
    // return
    window.location.href = url
  }

  beforeUpdate(() => {
    // console.log(initSearchData)
  })

  onMount(() => {
    // console.log(data)
    // console.log(container, ref)
    // container.appendChild(ref)
  })
</script>

<svelte:window
  on:click={({ target }) => {
    if (!comboboxRef?.contains(target)) {
      data.active = false
    }
  }}
/>

<main bind:this={comboboxRef} class="dropdown" class:dropdown-open={data.active}>
  <!-- svelte-ignore a11y-label-has-associated-control -->
  <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <label
    tabindex="0"
    class="btn-outline btn-sm btn"
    on:click={() => {
      searchRef?.focus()
      data.active = true
    }}
  >
    {#if base}
      <span class="text"
        >{#if data.pageIsComparePull}{data.textPullsCompareBase}{:else}{data.textCompareCompareBase}{/if}:
        {data.baseCompareName}:{data.baseBranch}</span
      >
    {:else}
      <span class="text"
        >{#if data.pageIsComparePull}{data.textPullsCompareCompare}{:else}{data.textCompareCompareHead}{/if}:
        {data.headCompareName}:{data.headBranch}</span
      >
    {/if}
    <Svg name="octicon-triangle-down" size={14} />
  </label>
  <div class="card dropdown-content card-compact p-2">
    <div class="card-body !p-0">
      <div class="tabs flex-nowrap !border-b-0">
        <button
          class="tab flex-nowrap gap-1 text-xs"
          class:tab-active={data.mode === 'branches'}
          class:!text-primary={data.mode === 'branches'}
          on:click={() => {
            data.mode = 'branches'
            searchRef?.focus()
          }}
        >
          <Svg name="octicon-git-branch" size={14} />
          {data.textBranches}
        </button>
        <button
          class="tab flex-nowrap gap-1 text-xs"
          class:tab-active={data.mode === 'tags'}
          class:!text-primary={data.mode === 'tags'}
          on:click={() => {
            data.mode = 'tags'
            searchRef?.focus()
          }}
        >
          <Svg name="octicon-tag" size={14} />
          {data.textTags}
        </button>
      </div>
    </div>

    <div data-repo-branch-tag>
      <WrappedTypeahead bind:value={data.searchTerm} bind:defaultSelectedValue bind:searchRef {handleSelect} {handleResults} showAllResultsOnOpen={true} data={initSearchData} placeholder={data.searchFieldPlaceholder} />
    </div>

    {#if showNoResults}
      <div class="text-center text-base-content/60">
        {data.noResults}
      </div>
    {/if}
  </div>
</main>
