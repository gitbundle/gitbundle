<script>
  import { onMount } from 'svelte'
  import Svg from './Svg.svelte'
  import jquery from 'jquery'
  export let data = {}
  export let handleDismiss = (e) => {
    // console.log(jquery(e.target).closest(data.selector))
    // e.parentNode.removeChild(e.parentNode)
    jquery(e.target).closest(data.selector).remove()
  }

  onMount(() => {
    // console.log(data)
  })
</script>

{#if data.type === 'info'}
  <div class={`alert mb-4 gap-0 rounded-lg border !border-primary bg-primary/10 font-bold text-primary ${data.extraClass}`} class:opacity-80={data.opacity}>
    <div>
      {#if data.icon}
        <Svg name={data.iconName ?? 'fa-circle-exclamation'} size="18" className="h-[18px] w-[18px]" />
      {/if}
      <div>
        {#if data.title}
          <h3 class="font-bold">{@html data.title}</h3>
          <div class="text-xs">{@html data.desc}</div>
        {:else}
          <div>{@html data.desc}</div>
        {/if}
      </div>
    </div>
  </div>
{:else if data.type === 'success'}
  <div class={`alert mb-4 gap-0 rounded-lg border !border-green-600 bg-green-600/10 font-bold text-green-600 ${data.extraClass}`} class:opacity-80={data.opacity}>
    <div class="flex flex-1 items-center gap-2">
      {#if data.icon}
        <Svg name={data.iconName ?? 'fa-circle-check-regular'} size="18" className="h-[18px] w-[18px]" />
      {/if}
      <div>
        {#if data.title}
          <h3 class="font-bold">{@html data.title}</h3>
          <div class="text-xs">{@html data.desc}</div>
        {:else}
          <div>{@html data.desc}</div>
        {/if}
      </div>
    </div>
    {#if data.close}
      <button class="flex-none cursor-pointer" on:click={handleDismiss}>
        <Svg name="fa-xmark-solid" size="18" className="h-[18px] w-[18px]" />
      </button>
    {/if}
  </div>
{:else if data.type === 'warning'}
  <div class={`alert mb-4 gap-0 rounded-lg border !border-yellow-500 bg-yellow-500/10 font-bold text-yellow-500 ${data.extraClass}`} class:opacity-80={data.opacity}>
    <div class="flex flex-1 items-center gap-2">
      {#if data.icon}
        <Svg name={data.iconName ?? 'fa-triangle-exclamation'} size="18" className="h-[18px] w-[18px]" />
      {/if}
      <div>
        {#if data.title}
          <h3 class="font-bold">{@html data.title}</h3>
          <div class="text-xs">{@html data.desc}</div>
        {:else}
          <div>{@html data.desc}</div>
        {/if}
      </div>
    </div>
    {#if data.close}
      <button class="flex-none cursor-pointer" on:click={handleDismiss}>
        <Svg name="fa-xmark-solid" size="18" className="h-[18px] w-[18px]" />
      </button>
    {/if}
  </div>
{:else if data.type === 'error'}
  <div class={`alert mb-4 gap-0 rounded-lg border !border-red-600 bg-red-600/10 font-bold text-red-600 ${data.extraClass}`} class:opacity-80={data.opacity}>
    <div class="flex flex-1 items-center gap-2">
      {#if data.icon}
        <Svg name={data.iconName ?? 'fa-circle-xmark-regular'} size="18" className="h-[18px] w-[18px]" />
      {/if}
      <div>
        {#if data.title}
          <h3 class="font-bold">{@html data.title}</h3>
          <div class="text-xs">{@html data.desc}</div>
        {:else}
          <div>{@html data.desc}</div>
        {/if}
      </div>
    </div>
    {#if data.close}
      <button class="flex-none cursor-pointer" on:click={handleDismiss}>
        <Svg name="fa-xmark-solid" size="18" className="h-[18px] w-[18px]" />
      </button>
    {/if}
  </div>
{:else}
  <div class={`alert mb-4 gap-0 rounded-lg border bg-base-300 font-bold ${data.extraClass}`} class:opacity-80={data.opacity}>
    <div>
      {#if data.icon}
        <Svg name={data.iconName ?? 'fa-circle-exclamation'} size="18" className="h-[18px] w-[18px]" />
      {/if}
      {#if data.title}
        <h3 class="font-bold">{@html data.title}</h3>
        <div class="text-xs">{@html data.desc}</div>
      {:else}
        <div>{@html data.desc}</div>
      {/if}
    </div>
  </div>
{/if}
