<script>
  import { onMount, beforeUpdate, afterUpdate } from 'svelte'
  import Svg from './Svg.svelte'
  import SimpleAlert from './GlobalSimpleAlert.svelte'

  const searchParams = new URLSearchParams(window.location.search)

  const { appSubUrl, assetUrlPrefix, csrfToken } = window.config
  const navbar = window.config?.pageData?.globalNavbar ?? {}
  const options = {
    openDropdownCreate: false,
    openDropdownAdmin: false,
    showMigrateModal: false,
    migrateModalID: `migrate-modal-` + Math.random(),
    migrateModalError: ''
  }

  let comboboxRef,
    isSearchBoxFocused = false,
    value = '',
    searchRef = null,
    visitedUsers = [],
    visitedRepos = [],
    toggleSignInSvg,
    toggleSignUpSvg

  if (navbar.PageIsExplore) {
    value = searchParams.get('q') ?? ''
  }
  const getVisitedHistories = async () => {
    const response = await fetch(`${appSubUrl}/${navbar.SignedUserName}/recent`)
    if (response.ok) {
      const res = await response.json()
      visitedUsers = []
      visitedRepos = []
      res.forEach((ele) => {
        if (ele.visitType === 'user') {
          visitedUsers.push(ele)
        } else if (ele.visitType === 'repo') {
          visitedRepos.push(ele)
        }
      })
      // console.log(visitedUsers, visitedRepos)
    }
  }
  const closeMigrateModal = async (e) => {
    options.openDropdownCreate = false
    options.showMigrateModal = false
    options.migrateModalError = ''
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

  onMount(() => {
    // console.log(navbar)
    // console.log(value, searchRef)
  })
</script>

<svelte:window
  on:keydown={(e) => {
    switch (e.key) {
      // TODO: need fix other input element input /
      // case '/':
      //   if (searchRef?.contains(e.target)) {
      //     searchRef?.focus()
      //     e.preventDefault()
      //   }
      //   break
      case 'Escape':
        searchRef?.blur()
        break
    }
  }}
  on:click={({ target }) => {
    if (!comboboxRef?.contains(target)) {
      options.openDropdownCreate = false
      options.openDropdownAdmin = false
    }
  }}
/>

<div class="flex-1 items-center gap-2">
  <a href={`${appSubUrl}/`} aria-label={navbar.textAriaLabel}>
    <img class="h-10 w-10" src={`${assetUrlPrefix}/img/logo.svg`} alt={navbar.textLogoAlt} aria-hidden="true" />
  </a>

  {#if navbar.IsSigned}
    <div class="dropdown" class:w-full={value !== '' || isSearchBoxFocused}>
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <label tabindex="0" class="input-bordered input flex w-full items-center !border-base-content/10 !bg-transparent bg-base-200 !px-0 hover:bg-base-300">
        <span class="flex h-full items-center px-2">
          <Svg name="octicon-search" size="16" className="h-4 w-4 text-base-content/60" />
        </span>

        <input
          class="w-full !bg-transparent focus:border-none focus:outline-none"
          bind:value
          bind:this={searchRef}
          on:focus|preventDefault={() => {
            getVisitedHistories()
            isSearchBoxFocused = true
          }}
          on:blur|preventDefault={() => {
            isSearchBoxFocused = false
          }}
        />

        <span class="flex h-full items-center justify-center pr-2">
          {#if !isSearchBoxFocused && value === ''}
            <kbd class="kbd kbd-sm text-base-content/60">/</kbd>
          {:else if value !== ''}
            <button
              class="h-full cursor-pointer px-2 hover:bg-mako-800"
              on:click={() => {
                isSearchBoxFocused = true
                value = ''
                searchParams.delete('q')
                const queryString = searchParams.toString()
                if (queryString) {
                  window.history.replaceState({}, '', `?${queryString}`)
                } else {
                  window.history.replaceState({}, '', window.location.pathname)
                }
                window.location.reload()
              }}
              type="button"
            >
              <Svg name="fa-circle-xmark-solid" size="16" className="h-4 w-4 text-base-content/60 fill-current" />
            </button>
          {/if}
        </span>
      </label>
      <ul class="cu-menu dropdown-content !max-h-96 w-full p-2" class:!hidden={value === '' && !visitedUsers.length && !visitedRepos.length}>
        {#if value !== ''}
          <li>
            <div class="flex items-center justify-between">
              <a href={`${appSubUrl}/explore/repos?q=${value}`} class="link-hover link flex items-center gap-2 hover:link-primary">
                <Svg name="octicon-search" size="16" className="h-4 w-4 text-base-content/60 fill-current" />
                {value}
              </a>
              <span class="text-gray-500">Search all of GitBundle</span>
            </div>
          </li>
          <div class="divider my-0" />
        {/if}
        {#if visitedUsers.length > 0}
          <li class="menu-title">
            <span>Owners</span>
          </li>
          {#each visitedUsers as item}
            <li>
              <div class="flex items-center justify-between">
                <a href={`${appSubUrl}/${item.content.name}`} class="link-hover link flex items-center gap-2 hover:link-primary">
                  <img class="h-4 w-4 rounded-full" src={item.content.avatar} alt={item.content.name} />
                  {item.content.name}
                </a>
                <span class="text-gray-500">Jump to</span>
              </div>
            </li>
          {/each}
          <div class="divider my-0" />
        {/if}
        {#if visitedRepos.length > 0}
          <li class="menu-title">
            <span>Repositories</span>
          </li>
          {#each visitedRepos as item}
            <li>
              <div class="flex items-center justify-between">
                <a href={`${appSubUrl}/${item.content.name}`} class="link-hover link flex items-center gap-2 hover:link-primary">
                  <Svg name={repoIcon(item.content)} size="16" className="h-4 w-4 text-base-content/60 fill-current" />
                  {item.content.name}
                </a>
                <span class="text-gray-500">Jump to</span>
              </div>
            </li>
          {/each}
        {/if}
      </ul>
    </div>
  {/if}

  <!-- {{ template "custom/extra_links" . }} -->
</div>

<div class="flex-none gap-4" bind:this={comboboxRef}>
  <ul class="cu-menu cu-menu-horizontal rounded-box p-1" class:!hidden={value !== '' || isSearchBoxFocused}>
    {#if navbar.IsSigned && navbar.MustChangePassword}
      <!-- /* No links */ -->
    {:else if navbar.IsSigned}
      {#if !navbar.UnitIssuesGlobalDisabled}
        <li>
          <a class="tooltip !tooltip-bottom" data-tip={navbar.textIssues} class:active={navbar.PageIsIssues} href={`${appSubUrl}/issues`}>
            <Svg name="octicon-issue-opened" />
          </a>
        </li>
      {/if}
      {#if !navbar.UnitPullsGlobalDisabled}
        <li>
          <a class="tooltip !tooltip-bottom" data-tip={navbar.textPullRequests} class:active={navbar.PageIsPulls} href={`${appSubUrl}/pulls`}>
            <Svg name="octicon-git-pull-request" />
          </a>
        </li>
      {/if}
      {#if !(navbar.UnitIssuesGlobalDisabled && navbar.UnitPullsGlobalDisabled)}
        {#if navbar.ShowMilestonesDashboardPage}
          <li>
            <a class="tooltip !tooltip-bottom" data-tip={navbar.textMilestones} class:active={navbar.PageIsMilestonesDashboard} href={`${appSubUrl}/milestones?q=${value}`}>
              <Svg name="octicon-milestone" />
            </a>
          </li>
        {/if}
      {/if}
      <!-- <li>
        <a class:active={navbar.PageIsExplore} href={`${appSubUrl}/explore/repos`}>{navbar.textExplore}</a>
      </li> -->
    {:else if navbar.IsLandingPageOrganizations}
      <!-- <li>
        <a class:active={navbar.PageIsExplore} href={`${appSubUrl}/explore/organizations`}>{navbar.textExplore}</a>
      </li> -->
    {:else}
      <!-- <li>
        <a class:active={navbar.PageIsExplore} href={`${appSubUrl}/explore/repos`}>{navbar.textExplore}</a>
      </li> -->
    {/if}
  </ul>

  {#if navbar.IsSigned && navbar.MustChangePassword}
    <div class="dropdown dropdown-end">
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <label tabindex="0" class="btn-ghost tooltip !tooltip-bottom btn flex w-fit" data-tip={navbar.textUserProfileAndMore}>
        <div class="avatar">
          <div class="w-8 rounded-full">
            {@html navbar.avatarSignedUser}
          </div>
        </div>
        <Svg name="octicon-triangle-down" />
      </label>
      <ul class="dropdown-content menu rounded-box menu-compact mt-3 whitespace-nowrap bg-base-200 p-2 shadow">
        <li>
          <a href data-url={`${appSubUrl}/user/logout" data-redirect="{{ AppSubUrl }}/`}>
            <Svg name="octicon-sign-out" />
            {navbar.textSignOut}<!-- Sign Out -->
          </a>
        </li>
      </ul>
    </div>
  {:else if navbar.IsSigned}
    {#if navbar.ActiveStopwatch}
      <div class="dropdown dropdown-end">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <label tabindex="0" class="btn-ghost tooltip !tooltip-bottom btn flex w-fit" data-tip={navbar.textActiveStopwatch}>
          <div class="indicator">
            <Svg name="octicon-stopwatch" />
            <span class="badge !badge-primary !badge-xs indicator-item" />
          </div>
        </label>
        <div class="card dropdown-content card-compact mt-3 whitespace-nowrap bg-base-200 shadow">
          <div class="card-body flex flex-row items-center">
            <a class="flex flex-row items-center" href={navbar.ActiveStopwatch.IssueLink}>
              <span class="text-primary"><Svg name="octicon-issue-opened" /></span>
              <span class="link-primary link">{navbar.ActiveStopwatch.RepoSlug}#{navbar.ActiveStopwatch.IssueIndex}</span>
              <span class="badge !badge-primary !badge-sm" data-seconds={navbar.ActiveStopwatchSeconds}>
                {#if navbar.ActiveStopwatch}
                  {navbar.ActiveStopwatchSeconds}
                {/if}
              </span>
            </a>
            <form class="btn-ghost btn-active btn-xs btn" method="POST" action={`${navbar.ActiveStopwatch.IssueLink}/times/stopwatch/toggle`}>
              <input type="hidden" name="_csrf" value={csrfToken} />
              <button class="tooltip !tooltip-bottom" data-tip={navbar.textRepoStopTracking}>
                <Svg name="octicon-square-fill" />
              </button>
            </form>
            <form class="btn-ghost btn-active btn-xs btn" method="POST" action={`${navbar.ActiveStopwatch.IssueLink}/times/stopwatch/cancel`}>
              <input type="hidden" name="_csrf" value={csrfToken} />
              <button class="tooltip !tooltip-bottom" data-tip={navbar.textRepoCancelTracking}>
                <Svg name="octicon-trash" />
              </button>
            </form>
          </div>
        </div>
      </div>
    {/if}

    <a href={`${appSubUrl}/notifications`} class="!btn-ghost tooltip !tooltip-bottom btn flex" data-tip={navbar.textNotifications}>
      <div class="indicator">
        <Svg name="octicon-bell" size="16" />
        <span class:hidden={!navbar.NotificationUnreadCount} class="badge !badge-primary !badge-xs indicator-item">{navbar.NotificationUnreadCount}</span>
      </div>
    </a>

    <div class="dropdown dropdown-end" class:dropdown-open={options.openDropdownCreate}>
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <label tabindex="0" class="!btn-ghost tooltip !tooltip-bottom btn flex w-fit" data-tip={navbar.textCreateNew}>
        <button type="button" class="indicator" on:click={() => (options.openDropdownCreate = true)}>
          <Svg name="octicon-plus" size="16" />
          <Svg name="octicon-triangle-down" size="16" />
        </button>
      </label>
      <ul class="dropdown-content menu rounded-box menu-compact mt-3 whitespace-nowrap bg-base-200 p-2 shadow">
        <li>
          <a href={`${appSubUrl}/repo/create`}>
            <span><Svg name="octicon-plus" /></span>
            {navbar.textNewRepo}
          </a>
        </li>
        {#if !navbar.DisableMigrations}
          <li>
            <a href={`${appSubUrl}/repo/migrate`}>
              <span><Svg name="octicon-repo-push" /></span>
              {navbar.textNewMigrate}
            </a>
          </li>
        {/if}
        {#if navbar.SignedUserCanCreateOrganization}
          <li>
            <a href={`${appSubUrl}/org/create`}>
              <span><Svg name="octicon-organization" /></span>
              {navbar.textNewOrg}
            </a>
          </li>
        {/if}
      </ul>
    </div>

    <div class="dropdown dropdown-end" class:dropdown-open={options.openDropdownAdmin}>
      <!-- svelte-ignore a11y-label-has-associated-control -->
      <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <label tabindex="0" class="!btn-ghost tooltip !tooltip-bottom btn flex w-fit" data-tip={navbar.textUserProfileAndMore} on:click={() => (options.openDropdownAdmin = true)}>
        <div class="avatar">
          <div class="w-8 rounded-full">
            {@html navbar.avatarSignedUser}
          </div>
        </div>
        <Svg name="octicon-triangle-down" />
      </label>
      <ul class="dropdown-content menu rounded-box menu-compact mt-3 whitespace-nowrap bg-base-200 p-2 shadow">
        <!-- <li class="menu-title justify-center flex-row space-x-1">
              <span class="p-0">{text signed_in_as}</span>
              <span class="font-bold p-0">{{ .SignedUser.Name }}</span>
            </li>
            <div class="divider m-0"></div> -->
        <li>
          <a href={navbar.SignedUserHomeLink}>
            <Svg name="octicon-person" />
            {navbar.textYourProfile}<!-- Your profile -->
          </a>
        </li>
        {#if !navbar.DisableStars}
          <li>
            <a href={`${navbar.SignedUserHomeLink}?tab=stars`}>
              <Svg name="octicon-star" />
              {navbar.textYourStarred}
            </a>
          </li>
        {/if}
        <li>
          <a class:active={navbar.PageIsUserSettings} href={`${appSubUrl}/user/settings`}>
            <Svg name="octicon-gear" />
            {navbar.textYourSettings}<!-- Your settings -->
          </a>
        </li>
        <li>
          <a target="_blank" rel="noopener noreferrer" href="https://docs.gitbundle.com">
            <Svg name="octicon-question" />
            {navbar.textHelp}<!-- Help -->
          </a>
        </li>

        {#if navbar.IsAdmin}
          <div class="divider !m-0" />

          <li>
            <a class:active={navbar.PageIsAdmin} href={`${appSubUrl}/admin`}>
              <Svg name="octicon-server" />
              {navbar.textAdminPanel}<!-- Admin Panel -->
            </a>
          </li>
        {/if}

        <div class="divider !m-0" />
        <li>
          <a href={`${appSubUrl}/user/logout`}>
            <Svg name="octicon-sign-out" />
            {navbar.textSignOut}<!-- Sign Out -->
          </a>
        </li>
      </ul>
    </div>
  {:else}
    <!-- <a
      class="btn-ghost btn normal-case"
      target="_blank"
      rel="noopener noreferrer"
      href="https://docs.gitbundle.com">{navbar.textHelp}</a
    > -->
    {#if navbar.ShowRegistrationButton}
      <a class:btn-active={navbar.PageIsSignUp} class="btn-outline-primary btn-outline btn-sm btn rounded-full normal-case" href={`${appSubUrl}/user/sign_up`} on:mouseenter={() => (toggleSignUpSvg = true)} on:mouseleave={() => (toggleSignUpSvg = false)}>
        {navbar.textSignUp}
        <Svg name="octicon-arrow-right" className={`animate-fade-right ${toggleSignUpSvg ? '' : 'hidden'}`} />
        <Svg name="octicon-chevron-right" className={`animate-fade-right ${toggleSignUpSvg ? 'hidden' : ''}`} />
      </a>
    {/if}
    <a class:btn-active={navbar.PageIsSignIn} class="btn-outline-primary btn-outline btn-sm btn rounded-full normal-case" href={`${appSubUrl}/user/login${!navbar.PageIsSignIn ? '?redirect_to=' + navbar.CurrentURL : ''}`} on:mouseenter={() => (toggleSignInSvg = true)} on:mouseleave={() => (toggleSignInSvg = false)}>
      {navbar.textSignIn}
      <Svg name="octicon-arrow-right" className={`animate-fade-right ${toggleSignInSvg ? '' : 'hidden'}`} />
      <Svg name="octicon-chevron-right" className={`animate-fade-right ${toggleSignInSvg ? 'hidden' : ''}`} />
    </a>
  {/if}
</div>
