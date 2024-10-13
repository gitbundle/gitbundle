<script>
  export let fullScreen
</script>

<div class={`flex items-center justify-center ${fullScreen ? 'h-screen' : 'max-h-screen'} w-full`}>
  <div class={`flex flex-col items-center rounded-lg px-5 py-2`}>
    <div class="loader-dots relative mt-2 block h-5 w-20">
      <div class={`absolute top-0 h-3 w-3 rounded-full bg-primary`} />
      <div class={`absolute top-0 h-3 w-3 rounded-full bg-primary`} />
      <div class={`absolute top-0 h-3 w-3 rounded-full bg-primary`} />
      <div class={`absolute top-0 h-3 w-3 rounded-full bg-primary`} />
    </div>
    <!-- {#if fullScreen} -->
    <!--   <div class="mt-2 text-center text-xs font-light text-base-content/60"> -->
    <!--     Please wait... -->
    <!--   </div> -->
    <!-- {/if} -->
  </div>
</div>
