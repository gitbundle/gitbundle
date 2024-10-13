<script>
  import SvgIcon from './Svg.svelte'
  import { beforeUpdate, onMount } from 'svelte'

  const csrfToken = window.config.csrfToken,
    mergeForm = window.config.pageData.pullRequestMergeForm

  let mergeTitleFieldValue = '',
    mergeMessageFieldValue = '',
    deleteBranchAfterMerge = false,
    autoMergeWhenSucceed = false,
    mergeStyleDetail = {
      // dummy only, these values will come from one of the mergeForm.mergeStyles
      hideMergeMessageTexts: false,
      textDoMerge: '',
      mergeTitleFieldText: '',
      mergeMessageFieldText: ''
    },
    showMergeStyleMenu = false,
    showActionForm = false,
    dropdownRef = null,
    dropdownOpen = false

  $: mergeButtonStyleClass = (() => {
    if (mergeForm.allOverridableChecksOk) return 'btn-primary'
    return autoMergeWhenSucceed ? 'btn-green' : 'btn-red'
  })()

  // created:
  let mergeStyleAllowedCount = mergeForm.mergeStyles.reduce((v, msd) => v + (msd.allowed ? 1 : 0), 0)
  let mergeStyle = mergeForm.mergeStyles.find((e) => e.allowed && e.name === mergeForm.defaultMergeStyle)?.name
  if (!mergeStyle) {
    mergeStyle = mergeForm.mergeStyles.find((e) => e.allowed)?.name
  }
  switchMergeStyle(mergeStyle, !mergeForm.canMergeNow)

  // DEBUG:
  // dropdownOpen = true

  // function hideMergeStyleMenu() {
  //   showMergeStyleMenu = false
  // }

  function toggleActionForm(show) {
    showActionForm = show
    if (!show) return
    deleteBranchAfterMerge = mergeForm.defaultDeleteBranchAfterMerge
    mergeTitleFieldValue = mergeStyleDetail.mergeTitleFieldText
    mergeMessageFieldValue = mergeStyleDetail.mergeMessageFieldText
  }

  function switchMergeStyle(name, autoMerge = false, e = null) {
    !e || e.stopPropagation()
    mergeStyle = name
    autoMergeWhenSucceed = autoMerge
    mergeStyleDetail = mergeForm.mergeStyles.find((e) => e.name === mergeStyle)
    dropdownOpen = false
  }

  beforeUpdate(() => {
    // console.log(dropdownOpen, mergeStyleDetail, mergeStyle)
  })

  onMount(() => {
    // console.log(mergeForm, mergeStyleDetail, mergeButtonStyleClass, mergeStyleAllowedCount, mergeStyle)
    // console.log(mergeStyleDetail, mergeForm.mergeStyles, mergeStyle)
  })
</script>

<svelte:window
  on:click={({ target }) => {
    if (!dropdownRef?.contains(target)) {
      dropdownOpen = false
    }
  }}
/>

{#if mergeForm.hasPendingPullRequestMerge}
  <div class="cu-links-primary mb-4 gap-2 break-words rounded-lg border !border-blue-600/40 p-2 text-blue-600">
    {@html mergeForm.hasPendingPullRequestMergeTip}
  </div>
{/if}
{#if showActionForm}
  <form action={mergeForm.baseLink + '/merge'} method="post">
    <input type="hidden" name="_csrf" value={csrfToken} />
    <input type="hidden" name="head_commit_id" bind:value={mergeForm.pullHeadCommitID} />
    <input type="hidden" name="merge_when_checks_succeed" bind:value={autoMergeWhenSucceed} />
    {#if !mergeStyleDetail.hideMergeMessageTexts}
      <div class="mb-3 space-y-2">
        <input type="text" name="merge_title_field" bind:value={mergeTitleFieldValue} class="input-bordered input w-full" />
        <textarea name="merge_message_field" rows="5" placeholder={mergeForm.mergeMessageFieldPlaceHolder} bind:value={mergeMessageFieldValue} class="textarea-bordered textarea w-full" />
      </div>
    {/if}
    <div class="items-cente flex gap-x-2">
      <button type="submit" name="do" value={mergeStyle} class={`btn ${mergeButtonStyleClass}`}>
        {mergeStyleDetail.textDoMerge}
        {#if autoMergeWhenSucceed}
          {mergeForm.textAutoMergeButtonWhenSucceed}
        {/if}
      </button>
      <button class="btn" on:click={() => toggleActionForm(false)}>
        {mergeForm.textCancel}
      </button>
      {#if mergeForm.isPullBranchDeletable && !autoMergeWhenSucceed}
        <label class="flex items-center gap-x-1">
          <input name="delete_branch_after_merge" type="checkbox" bind:value={deleteBranchAfterMerge} class="checkbox checkbox-sm" />
          {mergeForm.textDeleteBranch}
        </label>
      {/if}
    </div>
  </form>
{/if}
{#if !showActionForm}
  <div class="flex items-center">
    <div bind:this={dropdownRef} class="dropdown" class:dropdown-open={dropdownOpen}>
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <div class="btn-group">
        <span
          class="btn gap-x-1 {mergeForm.emptyCommit ? 'btn-gray' : mergeForm.allOverridableChecksOk ? 'btn-primary' : 'btn-red'}"
          on:click={() => {
            toggleActionForm(true)
            showMergeStyleMenu = !showMergeStyleMenu
          }}
        >
          <SvgIcon name="octicon-git-merge" />
          <span class="button-text">
            {mergeStyleDetail.textDoMerge}
            {#if autoMergeWhenSucceed}
              {mergeForm.textAutoMergeButtonWhenSucceed}
            {/if}
          </span>
        </span>
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label
          tabindex="0"
          class="btn border-0 !border-l !border-l-primary-content/10 {mergeForm.emptyCommit ? 'btn-gray' : mergeForm.allOverridableChecksOk ? 'btn-primary' : 'btn-red'}"
          on:click={() => {
            dropdownOpen = true
          }}><SvgIcon name="octicon-triangle-down" /></label
        >
      </div>
      <!-- && showMergeStyleMenu -->
      {#if mergeStyleAllowedCount > 1}
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <ul tabindex="0" class="dropdown-content menu p-2 shadow" class:hidden={!dropdownOpen}>
          {#each mergeForm.mergeStyles as msd}
            <!-- if can merge now, show one action "merge now", and an action "auto merge when succeed" -->
            {#if msd.allowed && mergeForm.canMergeNow}
              <li>
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <div class="flex flex-col gap-y-0" class:active={mergeStyle === msd.name} on:click={() => switchMergeStyle(msd.name)}>
                  <span class="w-full">{msd.textDoMerge}</span>
                  {#if !msd.hideAutoMerge}
                    <div class="flex items-center gap-x-1 text-xs" on:click={(e) => switchMergeStyle(msd.name, true, e)}>
                      <SvgIcon name="octicon-clock" size={14} className="!w-3 !h-3" />
                      {mergeForm.textAutoMergeWhenSucceed}
                    </div>
                  {/if}
                </div>
              </li>
              <!-- if can NOT merge now, only show one action "auto merge when succeed" -->
            {:else if msd.allowed && !mergeForm.canMergeNow && !msd.hideAutoMerge}
              <li>
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <div class="flex flex-col gap-y-0" class:active={mergeStyle === msd.name} on:click={() => switchMergeStyle(msd.name, true)}>
                  <span class="w-full">{msd.textDoMerge}</span>
                  <span class="w-full">{mergeForm.textAutoMergeButtonWhenSucceed}</span>
                </div>
              </li>
            {/if}
          {/each}
        </ul>
      {/if}
    </div>
    {#if mergeForm.hasPendingPullRequestMerge}
      <form action={mergeForm.baseLink + '/cancel_auto_merge'} method="post" class="ml-4">
        <input type="hidden" name="_csrf" value={csrfToken} />
        <button class="btn">
          {mergeForm.textAutoMergeCancelSchedule}
        </button>
      </form>
    {/if}
  </div>
{/if}
