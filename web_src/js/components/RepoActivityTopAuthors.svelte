<script>
  import { onMount } from 'svelte'
  import SvelteBarGraph from './SvelteBarGraph.svelte'

  export let pageData = []

  let graphPoints = pageData.map((item) => ({
      value: item.commits,
      label: item.name
    })),
    graphWidth = pageData.length * 40,
    colors = {
      barColor: 'hsl(var(--pf))',
      textColor: 'hsl(var(--bc))',
      textAltColor: 'hsl(var(--pc))'
    }
</script>

<SvelteBarGraph points={graphPoints} showXAxis={true} showYAxis={false} showValues={true} width={graphWidth} barColor={colors.barColor} textColor={colors.textColor} textAltColor={colors.textAltColor} height={100} labelHeight={20}>
  <svelte:fragment slot="label" let:bar>
    {#each pageData as author, idx}
      <g v-for="(author, idx) in graphAuthors" key={author.position}>
        {#if bar.index === idx && author.home_link}
          <a href={author.home_link}>
            <image x={`${bar.midPoint - 10}px`} y={`${bar.yLabel}px`} height="20" width="20" href={author.avatar_link} />
          </a>
        {:else if bar.index === idx}
          <image x={`${bar.midPoint - 10}px`} y={`${bar.yLabel}px`} height="20" width="20" href={author.avatar_link} />
        {/if}
      </g>
    {/each}
  </svelte:fragment>
  <svelte:fragment slot="title" let:bar>
    {#each pageData as author, idx}
      <tspan key={author.position}>
        {#if bar.index === idx}
          {author.name}
        {/if}
      </tspan>
    {/each}
  </svelte:fragment>
</SvelteBarGraph>
