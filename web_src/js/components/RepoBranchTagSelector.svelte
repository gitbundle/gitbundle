<script>
  import { beforeUpdate, onMount, tick } from 'svelte'
  import Svg from './Svg.svelte'
  import Typeahead from './WrappedTypeahead.svelte'
  import { pathEscapeSegments } from '../utils/url.js'

  export let pageData

  let comboboxRef,
    formRef,
    formActionUrl,
    searchRef,
    initSearchData = [],
    defaultSelectedValue = null,
    showNoResults = false,
    data = {
      csrfToken: window.config.csrfToken,
      searchTerm: '',
      refNameText: '',
      createTag: false,
      release: null,

      isViewTag: false,
      isViewBranch: false,
      isViewTree: false,

      active: false,

      ...pageData

      // the "data.defaultBranch" is ambiguous, it could be "branch name" or "tag name"
    }
  $: formActionUrl = `${data.repoLink}/branches/_new/${pathEscapeSegments(data.branchNameSubURL)}`
  $: {
    if (data.mode === 'tags') {
      initSearchData = data.tags
      if (!data.isViewTag) {
        defaultSelectedValue = null
      }
    } else if (data.mode === 'branches') {
      initSearchData = data.branches
      if (!data.isViewBranch) {
        defaultSelectedValue = null
      }
    }
  }
  $: if (data.viewType === 'tree') {
    data.isViewTree = true
    data.refNameText = data.commitIdShort
  } else if (data.viewType === 'tag') {
    data.isViewTag = true
    data.refNameText = data.tagName
    defaultSelectedValue = data.tagName
  } else {
    data.isViewBranch = true
    data.refNameText = data.branchName
    defaultSelectedValue = data.branchName
  }

  let showCreateNewBranch = !(data.disableCreateBranch || !data.searchTerm)
  const handleResults = (results) => {
    showCreateNewBranch = results.length === 0 && results.filter((item) => item.original.toLowerCase() === data.searchTerm.toLowerCase()).length === 0
    showNoResults = results.length === 0 && !showCreateNewBranch
  }

  const handleSelect = async (item) => {
    const url = data.mode === 'tags' ? data.tagURLPrefix + item.detail.original + data.tagURLSuffix : data.branchURLPrefix + item.detail.original + data.branchURLSuffix
    window.location.href = url
  }

  const createNewBranch = () => {
    if (!showCreateNewBranch) return
    formRef.submit()
  }

  beforeUpdate(() => {
    // console.log(showCreateNewBranch, data.searchTerm, data.active)
  })

  // onMount(() => {
  //   console.log(data)
  //   // console.log(container, ref)
  //   // container.appendChild(ref)
  // })
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
    class="btn-outline btn-sm btn items-center space-x-1 normal-case"
    on:click={() => {
      searchRef?.focus()
      data.active = true
    }}
  >
    {#if data.release}
      {data.textReleaseCompare}
    {:else}
      {#if data.isViewTag}
        <Svg name="octicon-tag" />
      {:else}
        <Svg name="octicon-git-branch" />
      {/if}
      <strong ref="dropdownRefName">{data.refNameText}</strong>
    {/if}
    <Svg name="octicon-triangle-down" size="14" />
  </label>
  <div class="card dropdown-content card-compact p-2">
    <div class="card-body !p-0">
      {#if data.showBranchesInDropdown}
        <div class="tabs flex-nowrap !border-b-0">
          <button
            class="tab flex-nowrap gap-1 text-xs"
            class:tab-active={data.mode === 'branches'}
            class:!text-primary={data.mode === 'branches'}
            on:click={() => {
              data.createTag = false
              data.mode = 'branches'
              searchRef?.focus()
            }}
          >
            <Svg name="octicon-git-branch" size={14} />
            {data.textBranches}
          </button>
          {#if !data.noTag}
            <button
              class="tab flex-nowrap gap-1 text-xs"
              class:tab-active={data.mode === 'tags'}
              class:!text-primary={data.mode === 'tags'}
              on:click={() => {
                data.createTag = true
                data.mode = 'tags'
                searchRef?.focus()
              }}
            >
              <Svg name="octicon-tag" size={14} />
              {data.textTags}
            </button>
          {/if}
        </div>
      {/if}

      <div data-repo-branch-tag>
        <Typeahead bind:value={data.searchTerm} bind:defaultSelectedValue bind:searchRef {handleSelect} {handleResults} showAllResultsOnOpen={true} data={initSearchData} placeholder={data.searchFieldPlaceholder} />
      </div>

      {#if showNoResults}
        <div class="text-center text-base-content/60">
          {data.noResults}
        </div>
      {/if}

      {#if showCreateNewBranch}
        <button class="text-sm text-primary" on:click={createNewBranch}>
          <div class="flex items-center" class:hidden={!data.createTag}>
            <Svg name="octicon-tag" size={14} />
            {@html data.textCreateTag.replace('%s', data.searchTerm)}
          </div>
          <div class="flex items-center" class:hidden={data.createTag}>
            <Svg name="octicon-git-branch" size={14} />
            {@html data.textCreateBranch.replace('%s', data.searchTerm)}
          </div>
          <div class="flex items-center text-xs">
            {#if data.isViewBranch}
              {data.textCreateBranchFrom.replace('%s', data.branchName)}
            {:else if data.isViewTag}
              {data.textCreateBranchFrom.replace('%s', data.tagName)}
            {:else}
              {data.textCreateBranchFrom.replace('%s', data.commitIdShort)}
            {/if}
          </div>
        </button>
        <form bind:this={formRef} action={formActionUrl} method="post">
          <input type="hidden" name="_csrf" value={data.csrfToken} />
          <input type="hidden" name="new_branch_name" bind:value={data.searchTerm} />
          <input type="hidden" name="create_tag" bind:value={data.createTag} />
          <input type="hidden" name="current_path" bind:value={data.treePath} />
        </form>
      {/if}
    </div>
  </div>
</main>
