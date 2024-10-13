<script>
  import { onMount, beforeUpdate } from 'svelte'
  import MultiSelect from 'svelte-multiselect'
  import Svg from './Svg.svelte'
  import { initRepoGraphGit } from '../modules/repo-graph.js'
  import WrappedMultiSelectStyle from './WrappedMultiSelectStyle.svelte'

  export let pageData

  const { graphContainer, url, params, updateGraph, dropdownSelected, initJqEvent } = initRepoGraphGit()

  const handleAdd = ({ detail: { option } }) => {
    if (option.Name === '...flow-hide-pr-refs') {
      params.set('hide-pr-refs', true)
    } else {
      params.append('branch', option.Name)
    }
    updateGraph()
  }
  const handleRemove = ({ detail: { option } }) => {
    if (option.Name === '...flow-hide-pr-refs') {
      params.delete('hide-pr-refs')
    } else {
      const branches = params.getAll('branch')
      params.delete('branch')
      for (const branch of branches) {
        if (branch !== option.Name) {
          params.append('branch', branch)
        }
      }
    }
    updateGraph()
  }
  const handleRemoveAll = (e) => {
    params.delete('branch')
    params.delete('hide-pr-refs')
    updateGraph()
  }

  beforeUpdate(() => {})

  onMount(() => {
    initJqEvent()
  })
</script>

<div class="input-group w-fit">
  <MultiSelect options={pageData.AllRefs} placeholder={pageData.CommitGraphSelect} on:add={handleAdd} on:remove={handleRemove} on:removeAll={handleRemoveAll}>
    <div let:option slot="selected" class="flex items-center truncate py-2" data-value={option.Name}>
      {#if option.RefGroup === 'pull'}
        <Svg name="octicon-git-pull-request" size="16" />
      {:else if option.RefGroup === 'tags'}
        <Svg name="octicon-tag" size="16" />
      {:else if option.RefGroup === 'remotes'}
        <Svg name="octicon-cross-reference" size="16" />
      {:else if option.RefGroup === 'heads'}
        <Svg name="octicon-git-branch" size="16" />
      {:else if option.RefGroup === ''}
        <Svg name="octicon-eye-closed" size="16" />
      {/if}

      {#if option.RefGroup === 'pull'}
        <span title={option.ShortName}>#{option.ShortName}</span>
      {:else}
        <span title={option.ShortName}>{option.ShortName}</span>
      {/if}
    </div>
    <div let:option slot="option" class="flex items-center truncate" data-value={option.Name}>
      {#if option.RefGroup === 'pull'}
        <Svg name="octicon-git-pull-request" size="16" />
      {:else if option.RefGroup === 'tags'}
        <Svg name="octicon-tag" size="16" />
      {:else if option.RefGroup === 'remotes'}
        <Svg name="octicon-cross-reference" size="16" />
      {:else if option.RefGroup === 'heads'}
        <Svg name="octicon-git-branch" size="16" />
      {:else if option.RefGroup === ''}
        <Svg name="octicon-eye-closed" size="16" />
      {/if}

      {#if option.RefGroup === 'pull'}
        <span title={option.ShortName}>#{option.ShortName}</span>
      {:else}
        <span title={option.ShortName}>{option.ShortName}</span>
      {/if}
    </div>
  </MultiSelect>
  <button id="flow-color-monochrome" class={`${pageData.Mode === 'monochrome' ? 'btn-active' : ''} btn-outline btn-sm btn`} title={pageData.CommitGraphMonochrome}><Svg name="material-invert-colors" size="16" className="mr-1" />{pageData.CommitGraphMonochrome}</button>
  <button id="flow-color-colored" class={`${pageData.Mode !== 'monochrome' ? 'btn-active' : ''} btn-outline btn-sm btn`} title={pageData.CommitGraphColor}><Svg name="material-palette" size="16" className="mr-1" />{pageData.CommitGraphColor}</button>
</div>
<WrappedMultiSelectStyle />
