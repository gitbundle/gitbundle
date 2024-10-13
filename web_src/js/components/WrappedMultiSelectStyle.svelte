<style lang="postcss">
  :global(div.multiselect) {
    /* top-level wrapper div */
    /* @apply input-bordered input input-sm; */
    @apply !rounded-l-lg !rounded-r-none !border !border-base-content/10;
    @apply !bg-base-200 !text-base-content;
  }
  :global(div.multiselect.open) {
  }
  :global(div.multiselect.disabled) {
    @apply !bg-gray-700 !text-gray-600;
  }
  :global(div.multiselect > ul.selected) {
    /* selected list */
  }
  :global(div.multiselect > ul.selected > li) {
    /* selected list items */
    @apply !bg-base-content/10;
  }
  :global(div.multiselect button) {
    /* target all buttons in this component */
    @apply hover:!text-primary;
  }
  :global(div.multiselect > ul.selected > input) {
    @apply placeholder:!text-gray-500;
  }
  :global(div.multiselect > ul.selected > li button, button.remove-all) {
    /* buttons to remove a single or all selected options at once */
  }
  :global(div.multiselect > input[autocomplete]) {
    /* input inside the top-level wrapper div */
  }
  :global(div.multiselect > ul.options) {
    @apply !mt-1 !max-h-52 !w-full !rounded-lg !bg-base-200/70 !p-2 !shadow !backdrop-blur-md;
  }
  :global(div.multiselect > ul.options > li) {
    /* dropdown list items */
    @apply !flex !h-5 !items-center rounded-lg !p-4 !leading-5;
  }
  :global(div.multiselect > ul > li span) {
    @apply !bg-transparent !px-2;
  }
  :global(div.multiselect > ul.options > span) {
    @apply !rounded-lg !bg-gray-800 !text-gray-600;
  }
  :global(div.multiselect > ul.options > li.selected) {
    /* selected options in the dropdown list */
    @apply bg-base-content/10;
  }
  :global(div.multiselect > ul.options > li:not(.selected):hover) {
    /* unselected but hovered options in the dropdown list */
    @apply bg-base-content/10;
  }
  :global(div.multiselect > ul.options > li.active) {
    @apply bg-base-content/10;
  }
  :global(div.multiselect > ul.options > li.disabled) {
    @apply !bg-gray-700 !text-gray-600;
  }
</style>
