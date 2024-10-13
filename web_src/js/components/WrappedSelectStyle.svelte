<style lang="postcss">
  :global(div.svelte-select) {
    @apply flex items-center !border-base-content/10;
  }
  :global(div.svelte-select > .svelte-select-list) {
    @apply !rounded-lg !bg-base-200/70 !p-2 !backdrop-blur-md;
  }

  :global(div.svelte-select > .svelte-select-list > .list-item) {
    @apply flex h-fit;
  }
  :global(div.svelte-select > .svelte-select-list > .list-item > .item),
  :global(div.svelte-select > .svelte-select-list > .list-item > .item.first) {
    @apply h-fit !rounded-lg py-1;
  }

  :global(div.svelte-select > .svelte-select-list > .list-item > .item.hover) {
    @apply !bg-base-content/10;
  }

  :global(div.svelte-select > .svelte-select-list > .list-item > .item.active) {
    @apply !bg-base-content/10;
  }

  :global(div.svelte-select > .value-container) {
    @apply h-full flex-wrap items-center !py-0.5;
  }
  :global(div.svelte-select > .value-container > .selected-item) {
    @apply flex h-fit flex-wrap items-center;
  }
  :global(div.svelte-select > .value-container > .multi-item) {
    @apply !flex !h-full !items-center !rounded-lg !bg-base-content/10 !outline-mako-800;
  }
  :global(div.svelte-select > .value-container > input) {
    @apply !m-0 !h-full !border-none !border-base-content/10 !p-0 placeholder:!text-gray-500;
  }
  :global(div.svelte-select > .value-container > .multi-item > .multi-item-clear) {
    --clear-icon-color: hsl(var(--bc));
  }
</style>
