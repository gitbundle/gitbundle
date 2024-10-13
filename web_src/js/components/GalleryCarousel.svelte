<script>
  import { onMount } from 'svelte'
  export let images
  export let index
  export let idPrefix

  const { assetUrlPrefix } = window.config

  onMount(async () => {
    // console.log(images)
  })
</script>

{#each images as image, j}
  <div id={`${idPrefix}-${index}-${j}`} class="carousel-item relative w-full">
    <div class="flex items-center">
      <img src={`${assetUrlPrefix}/${image.src}`} alt={image.alt} />
    </div>
    <div class="absolute left-5 right-5 top-1/2 flex -translate-y-1/2 transform justify-between">
      <a href={`#${idPrefix}-${index}-${j - 1 >= 0 ? j - 1 : j - 1 + images.length}`} class="slide-button btn-circle btn">❮</a>
      <a href={`#${idPrefix}-${index}-${(j + 1) % images.length}`} class="slide-button btn-circle btn">❯</a>
    </div>
  </div>
{/each}

<style lang="postcss">
  .slide-button {
    @apply border-none bg-base-content/5 !font-normal !text-base-content/10 hover:bg-base-content/40 hover:!text-primary-content;
  }
</style>
