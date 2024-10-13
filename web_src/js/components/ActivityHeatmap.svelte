<script>
  import { onMount, beforeUpdate } from 'svelte'
  import CalendarHeatmap from './CalendarHeatmap.svelte'

  export let values
  export let locale

  function handleDayClick(e) {
    // Reset filter if same date is clicked
    const params = new URLSearchParams(document.location.search)
    const queryDate = params.get('date')
    // Timezone has to be stripped because toISOString() converts to UTC
    const clickedDate = new Date(e.detail.date - e.detail.date.getTimezoneOffset() * 60000).toISOString().substring(0, 10)
    if (queryDate && queryDate === clickedDate) {
      params.delete('date')
    } else {
      params.set('date', clickedDate)
    }
    params.delete('page')
    const newSearch = params.toString()
    window.location.search = newSearch.length ? `?${newSearch}` : ''
  }

  const rangeColor = [
    'hsla(var(--b2) / var(--tw-text-opacity, 1))',
    '#1e3a8a', // 'hsla(var(--pf) / var(--tw-text-opacity, 0.8))',
    '#1e40af', // 'hsla(var(--p) / var(--tw-text-opacity, 0.2))',
    '#1d4ed8', // 'hsla(var(--p) / var(--tw-text-opacity, 0.4))',
    '#2563eb', // 'hsla(var(--p) / var(--tw-text-opacity, 1))'
    '#3b82f6'
  ]
  const endDate = new Date()
  let sum = 0

  beforeUpdate(() => {
    sum = 0
    for (let i = 0; i < values.length; i++) {
      sum += values[i].count
    }
  })

  onMount(() => {
    // work around issue with first legend color being rendered twice and legend cut off
    // const legend = document.querySelector('.vch__external-legend-wrapper');
    // legend.setAttribute('viewBox', '12 0 80 10');
    // legend.style.marginRight = '-12px';
  })
</script>

<CalendarHeatmap {locale} noDataText={locale.no_contributions} tooltipUnit={locale.contributions} {endDate} {values} {rangeColor} on:day-click={handleDayClick}>
  <div slot="vch__legend-left">
    <span class="text-base-content">{sum}</span> contributions in the last 12 months
  </div>
</CalendarHeatmap>

<style lang="postcss">
  :global(.vch__wrapper) {
    @apply fill-current;
  }
  :global(.vch__legend) {
    @apply flex items-center justify-between gap-x-1;
  }
</style>
