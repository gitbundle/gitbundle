<script>
  import { onMount, beforeUpdate } from 'svelte'
  import Typeahead from './WrappedTypeahead.svelte'
  import Svg from './Svg.svelte'

  export let pageData

  const extract = (item) => {
    return item.Name
  }

  const handleSelect = ({
    detail: {
      original: { SubURL }
    }
  }) => {
    // console.log(pageData.RepoLink, SubURL)
    window.location.href = `${pageData.RepoLink}/wiki/${SubURL}`
  }

  onMount(() => {
    // console.log(pageData)
  })
</script>

<div class="rounded-lg border">
  <div class="w-full rounded-t-lg bg-base-200 px-4 py-2">
    {pageData.RepoWikiPage}:
    <strong>{pageData.Title}</strong>
  </div>
  <div class="divide-y border-t px-4 pt-4" title={pageData.textFilter}>
    <Typeahead placeholder={pageData.RepoWikiFilterPage} {handleSelect} {extract} data={pageData.Pages} showAllResultsOnOpen={true} defaultSelectedValue={pageData.CurrentTitle} className="!input !input-sm !pr-8">
      <svelte:fragment slot="typeahead-right">
        <Svg name="octicon-filter" size="16" className="m-2 h-4 w-4 text-base-content/60" />
      </svelte:fragment>
      <svelte:fragment slot="typeahead-no-results" let:value>
        <div class="text-center">{value}{pageData.RepoPullsNoResults}</div>
      </svelte:fragment>
      <svelte:fragment let:result>
        <a class:active={result.Name === pageData.CurrentTitle} href={`${pageData.RepoLink}/wiki/${result.SubURL}`}>
          {result.Name}
        </a>
      </svelte:fragment>
    </Typeahead>
  </div>
</div>

<style lang="postcss">
  :global([svelte-repo-wiki-sidebar-nav] [data-svelte-typeahead] .svelte-typeahead-list) {
    @apply mb-4;
  }
  :global([svelte-repo-wiki-sidebar-nav] [data-svelte-typeahead] .svelte-typeahead-list .selected) {
    @apply !bg-primary;
  }
</style>
