<script context="module" lang="ts">
  export interface Value {
    date: Date | string
    count: number
  }

  export interface Activity {
    count: number
    colorIndex: number
  }

  export type Activities = Map<string, Activity>

  export interface CalendarItem {
    date: Date
    count?: number
    colorIndex: number
  }

  export type Calendar = CalendarItem[][]

  export interface Month {
    value: number
    index: number
  }

  export interface Locale {
    months: string[]
    days: string[]
    on: string
    less: string
    more: string
  }

  export type TooltipFormatter = (item: CalendarItem, unit: string) => string

  export class Heatmap {
    static readonly DEFAULT_RANGE_COLOR_LIGHT = ['#ebedf0', '#dae2ef', '#c0ddf9', '#73b3f3', '#3886e1', '#17459e']
    //#1F1F22 #224367 #0E65B7 #0076E0 #1893FF #00B7FF
    static readonly DEFAULT_RANGE_COLOR_DARK = ['#1F1F22', '#224367', '#0E65B7', '#0076E0', '#1893FF', '#00B7FF']
    // other color candidates
    // static readonly DEFAULT_RANGE_COLOR_LIGHT = [ '#ebedf0', '#9be9a8', '#40c463', '#30a14e', '#216e39' ];
    // static readonly DEFAULT_RANGE_COLOR_DARK  = [ '#161b22', '#0e4429', '#006d32', '#26a641', '#39d353' ];
    // static readonly DEFAULT_RANGE_COLOR_DARK    = [ '#011526', '#012E40', '#025959', '#02735E', '#038C65' ];
    // static readonly DEFAULT_RANGE_COLOR_DARK    = [ '#161b22', '#015958', '#008F8C', '#0CABA8', '#0FC2C0' ];
    // static readonly DEFAULT_RANGE_COLOR_DARK    = [ '#012030', '#13678A', '#45C4B0', '#9AEBA3', '#DAFDBA' ];
    static readonly DEFAULT_LOCALE: Locale = {
      months: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
      days: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'],
      on: 'on',
      less: 'Less',
      more: 'More'
    }
    static readonly DEFAULT_TOOLTIP_UNIT = 'contributions'
    static readonly DAYS_IN_ONE_YEAR = 365
    static readonly DAYS_IN_WEEK = 7
    static readonly SQUARE_SIZE = 10

    startDate: Date
    endDate: Date
    max: number

    private _values: Value[]
    private _firstFullWeekOfMonths?: Month[]
    private _activities?: Activities
    private _calendar?: Calendar

    constructor(endDate: Date | string, values: Value[], max?: number) {
      this.endDate = this.parseDate(endDate)
      this.max = max || Math.ceil((Math.max(...values.map((day) => day.count)) / 5) * 4)
      this.startDate = this.shiftDate(endDate, -Heatmap.DAYS_IN_ONE_YEAR)
      this._values = values
    }

    set values(v: Value[]) {
      this.max = Math.ceil((Math.max(...v.map((day) => day.count)) / 5) * 4)
      this._values = v
      this._firstFullWeekOfMonths = undefined
      this._calendar = undefined
      this._activities = undefined
    }

    get values(): Value[] {
      return this._values
    }

    get activities(): Activities {
      if (!this._activities) {
        this._activities = new Map()
        for (let i = 0, len = this.values.length; i < len; i++) {
          this._activities.set(this.keyDayParser(this.values[i].date), {
            count: this.values[i].count,
            colorIndex: this.getColorIndex(this.values[i].count)
          })
        }
      }
      return this._activities
    }

    get weekCount() {
      return this.getDaysCount() / Heatmap.DAYS_IN_WEEK
    }

    get calendar() {
      if (!this._calendar) {
        let date = this.shiftDate(this.startDate, -this.getCountEmptyDaysAtStart())
        date = new Date(date.getFullYear(), date.getMonth(), date.getDate())
        this._calendar = new Array(this.weekCount)
        for (let i = 0, len = this._calendar.length; i < len; i++) {
          this._calendar[i] = new Array(Heatmap.DAYS_IN_WEEK)
          for (let j = 0; j < Heatmap.DAYS_IN_WEEK; j++) {
            const dayValues = this.activities.get(this.keyDayParser(date))
            this._calendar![i][j] = {
              date: new Date(date.valueOf()),
              count: dayValues ? dayValues.count : undefined,
              colorIndex: dayValues ? dayValues.colorIndex : 0
            }
            date.setDate(date.getDate() + 1)
          }
        }
      }
      return this._calendar
    }

    get firstFullWeekOfMonths(): Month[] {
      if (!this._firstFullWeekOfMonths) {
        const cal = this.calendar
        this._firstFullWeekOfMonths = []
        for (let index = 1, len = cal.length; index < len; index++) {
          const lastWeek = cal[index - 1][0].date,
            currentWeek = cal[index][0].date
          if (lastWeek.getFullYear() < currentWeek.getFullYear() || lastWeek.getMonth() < currentWeek.getMonth()) {
            this._firstFullWeekOfMonths.push({
              value: currentWeek.getMonth(),
              index
            })
          }
        }
      }
      return this._firstFullWeekOfMonths
    }

    getColorIndex(count?: number) {
      if (count === null || count === undefined) {
        return 0
      } else if (count <= 0) {
        return 1
      } else if (count >= this.max) {
        return 5
      } else {
        return Math.ceil(((count * 100) / this.max) * 0.03) + 1
      }
    }

    getCountEmptyDaysAtStart() {
      return this.startDate.getDay()
    }

    getCountEmptyDaysAtEnd() {
      return Heatmap.DAYS_IN_WEEK - 1 - this.endDate.getDay()
    }

    getDaysCount() {
      return Heatmap.DAYS_IN_ONE_YEAR + 1 + this.getCountEmptyDaysAtStart() + this.getCountEmptyDaysAtEnd()
    }

    private shiftDate(date: Date | string, numDays: number) {
      const newDate = new Date(date)
      newDate.setDate(newDate.getDate() + numDays)
      return newDate
    }

    private parseDate(entry: Date | string) {
      return entry instanceof Date ? entry : new Date(entry)
    }

    private keyDayParser(date: Date | string) {
      const day = this.parseDate(date)
      return String(day.getFullYear()) + String(day.getMonth()).padStart(2, '0') + String(day.getDate()).padStart(2, '0')
    }
  }
</script>

<script lang="ts">
  import { onMount, onDestroy, createEventDispatcher } from 'svelte'
  import tippy, { createSingleton, CreateSingletonInstance, Instance } from 'tippy.js'
  import 'tippy.js/dist/tippy.css'
  import 'tippy.js/dist/svg-arrow.css'

  export let endDate: Date // required
  export let max: number // Number
  export let rangeColor: string[] // Array
  export let values: Value[] // Array, required
  export let locale: Partial<Locale> // Object
  export let tooltip: Boolean = true // Boolean
  export let tooltipUnit: string = Heatmap.DEFAULT_TOOLTIP_UNIT // String
  export let tooltipFormatter: TooltipFormatter // Function
  export let vertical: Boolean = false // Boolean
  export let noDataText: Boolean | string = null // [Boolean, String]
  export let round: number = 0 // Number
  export let darkMode: Boolean // Boolean

  const dispatch = createEventDispatcher()

  let SQUARE_BORDER_SIZE = Heatmap.SQUARE_SIZE / 5,
    SQUARE_SIZE = Heatmap.SQUARE_SIZE + SQUARE_BORDER_SIZE,
    LEFT_SECTION_WIDTH = Math.ceil(Heatmap.SQUARE_SIZE * 2.5),
    RIGHT_SECTION_WIDTH = SQUARE_SIZE * 3,
    TOP_SECTION_HEIGHT = Heatmap.SQUARE_SIZE + Heatmap.SQUARE_SIZE / 2,
    BOTTOM_SECTION_HEIGHT = Heatmap.SQUARE_SIZE + Heatmap.SQUARE_SIZE / 2,
    yearWrapperTransform = `translate(${LEFT_SECTION_WIDTH}, ${TOP_SECTION_HEIGHT})`,
    now = new Date(),
    heatmap = new Heatmap(endDate as Date, values, max),
    width = 0,
    height = 0,
    viewbox = '0 0 0 0',
    legendViewbox = '0 0 0 0',
    daysLabelWrapperTransform = '',
    monthsLabelWrapperTransform = '',
    legendWrapperTransform = '',
    lo = <Locale>({} as any),
    tippyInstances = new Map<HTMLElement, Instance>()

  let tippySingleton: CreateSingletonInstance

  $: {
    rangeColor = rangeColor || (darkMode ? Heatmap.DEFAULT_RANGE_COLOR_DARK : Heatmap.DEFAULT_RANGE_COLOR_LIGHT)

    if (vertical) {
      width = LEFT_SECTION_WIDTH + SQUARE_SIZE * Heatmap.DAYS_IN_WEEK + RIGHT_SECTION_WIDTH
      height = TOP_SECTION_HEIGHT + SQUARE_SIZE * heatmap.weekCount + SQUARE_BORDER_SIZE
      daysLabelWrapperTransform = `translate(${LEFT_SECTION_WIDTH}, 0)`
      monthsLabelWrapperTransform = `translate(0, ${TOP_SECTION_HEIGHT})`
    } else {
      width = LEFT_SECTION_WIDTH + SQUARE_SIZE * heatmap.weekCount + SQUARE_BORDER_SIZE
      height = TOP_SECTION_HEIGHT + SQUARE_SIZE * Heatmap.DAYS_IN_WEEK
      daysLabelWrapperTransform = `translate(0, ${TOP_SECTION_HEIGHT})`
      monthsLabelWrapperTransform = `translate(${LEFT_SECTION_WIDTH}, 0)`
    }

    viewbox = ` 0 0 ${width} ${height}`
    legendWrapperTransform = vertical ? `translate(${LEFT_SECTION_WIDTH + SQUARE_SIZE * Heatmap.DAYS_IN_WEEK}, ${TOP_SECTION_HEIGHT})` : `translate(${width - SQUARE_SIZE * rangeColor.length - 30}, ${height - BOTTOM_SECTION_HEIGHT})`

    lo = locale ? { ...Heatmap.DEFAULT_LOCALE, ...locale } : Heatmap.DEFAULT_LOCALE
    legendViewbox = `0 0 ${Heatmap.SQUARE_SIZE * (rangeColor.length + 1)} ${Heatmap.SQUARE_SIZE}`

    heatmap = new Heatmap(endDate, values, max)
    tippyInstances.forEach((item) => item.destroy())
  }

  function initTippy() {
    tippyInstances.clear()
    if (tippySingleton) {
      tippySingleton.setInstances(Array.from(tippyInstances.values()))
    } else {
      tippySingleton = createSingleton(Array.from(tippyInstances.values()), {
        overrides: [],
        moveTransition: 'transform 0.1s ease-out',
        allowHTML: true
      })
    }
  }

  function tooltipOptions(day: CalendarItem) {
    if (tooltip) {
      if (day.count !== undefined) {
        if (tooltipFormatter) {
          return tooltipFormatter(day, tooltipUnit!)
        }
        return `<b>${day.count} ${tooltipUnit}</b> ${lo.on} ${lo.months[day.date.getMonth()]} ${day.date.getDate()}, ${day.date.getFullYear()}`
      } else if (noDataText) {
        return `${noDataText}`
      } else if (noDataText !== false) {
        return `<b>No ${tooltipUnit}</b> ${lo.on} ${lo.months[day.date.getMonth()]} ${day.date.getDate()}, ${day.date.getFullYear()}`
      }
    }
    return undefined
  }

  function getWeekPosition(index: number) {
    if (vertical) {
      return `translate(0, ${SQUARE_SIZE * heatmap.weekCount - (index + 1) * SQUARE_SIZE})`
    }
    return `translate(${index * SQUARE_SIZE}, 0)`
  }

  function getDayPosition(index: number) {
    if (vertical) {
      return `translate(${index * SQUARE_SIZE}, 0)`
    }
    return `translate(0, ${index * SQUARE_SIZE})`
  }

  function getMonthLabelPosition(month: Month) {
    if (vertical) {
      return {
        x: 3,
        y: SQUARE_SIZE * heatmap.weekCount - SQUARE_SIZE * month.index - SQUARE_SIZE / 4
      }
    }
    return { x: SQUARE_SIZE * month.index, y: SQUARE_SIZE - SQUARE_BORDER_SIZE }
  }

  function initTippyLazy(e: MouseEvent) {
    if (tippySingleton && e.target && (e.target as HTMLElement).classList.contains('vch__day__square') && (e.target as HTMLElement).dataset.weekIndex !== undefined && (e.target as HTMLElement).dataset.dayIndex !== undefined) {
      const weekIndex = Number((e.target as HTMLElement).dataset.weekIndex),
        dayIndex = Number((e.target as HTMLElement).dataset.dayIndex)

      if (!isNaN(weekIndex) && !isNaN(dayIndex)) {
        const tooltip = tooltipOptions(heatmap.calendar[weekIndex][dayIndex])
        if (tooltip) {
          const instance = tippyInstances.get(e.target as HTMLElement)

          if (instance) {
            instance.setContent(tooltip)
          } else if (!instance) {
            tippyInstances.set(e.target as HTMLElement, tippy(e.target as HTMLElement, { content: tooltip } as any))
            tippySingleton.setInstances(Array.from(tippyInstances.values()))
          }
        }
      }
    }
  }

  onMount(initTippy)
  onDestroy(() => {
    tippySingleton?.destroy()
    tippyInstances.forEach((item) => item.destroy())
  })
</script>

<div class={`vch__container ${darkMode ? 'dark-mode' : ''}`}>
  <svg class="vch__wrapper" viewBox={viewbox}>
    <g class="vch__months__labels__wrapper" transform={monthsLabelWrapperTransform}>
      {#each heatmap.firstFullWeekOfMonths as month, index}
        <text class="vch__month__label" x={getMonthLabelPosition(month).x} y={getMonthLabelPosition(month).y}>
          {lo.months[month.value]}
        </text>
      {/each}
    </g>
    <g class="vch__days__labels__wrapper" transform={daysLabelWrapperTransform}>
      <text class="vch__day__label" x={vertical ? SQUARE_SIZE : 0} y={vertical ? SQUARE_SIZE - SQUARE_BORDER_SIZE : 20}>
        {lo.days[1]}
      </text>
      <text class="vch__day__label" x={vertical ? SQUARE_SIZE * 3 : 0} y={vertical ? SQUARE_SIZE - SQUARE_BORDER_SIZE : 44}>
        {lo.days[3]}
      </text>
      <text class="vch__day__label" x={vertical ? SQUARE_SIZE * 5 : 0} y={vertical ? SQUARE_SIZE - SQUARE_BORDER_SIZE : 69}>
        {lo.days[5]}
      </text>
    </g>
    {#if vertical}
      <g class="vch__legend__wrapper" transform={legendWrapperTransform}>
        <text x={SQUARE_SIZE * 1.25} y="8">
          {lo.less}
        </text>
        {#each rangeColor as color, index}
          <rect rx={round} ry={round} style:fill={color} width={SQUARE_SIZE - SQUARE_BORDER_SIZE} height={SQUARE_SIZE - SQUARE_BORDER_SIZE} x={SQUARE_SIZE * 1.75} y={SQUARE_SIZE * (index + 1)} />
        {/each}
        <text x={SQUARE_SIZE * 1.25} y={SQUARE_SIZE * (rangeColor.length + 2) - SQUARE_BORDER_SIZE}>
          {lo.more}
        </text>
      </g>
    {/if}
    <g class="vch__year__wrapper" transform={yearWrapperTransform} on:focus on:mouseover={initTippyLazy}>
      {#each heatmap.calendar as week, weekIndex}
        <g class="vch__month__wrapper" transform={getWeekPosition(weekIndex)}>
          {#each week as day, dayIndex}
            {#if day.date < now}
              <rect class="vch__day__square" rx={round} ry={round} transform={getDayPosition(dayIndex)} width={SQUARE_SIZE - SQUARE_BORDER_SIZE} height={SQUARE_SIZE - SQUARE_BORDER_SIZE} style:fill={rangeColor[day.colorIndex]} data-week-index={weekIndex} data-day-index={dayIndex} on:keydown on:keyup on:keypress on:click={() => dispatch('day-click', day)} />
            {/if}
          {/each}
        </g>
      {/each}
    </g>
  </svg>
  <div class="vch__legend">
    <slot name="legend">
      <div class="vch__legend-left">
        <slot name="vch__legend-left" />
      </div>
      <div class="vch__legend-right">
        <slot name="legend-right">
          <div class="vch__legend">
            <div>
              {lo.less}
            </div>
            {#if !vertical}
              <svg class="vch__external-legend-wrapper" viewBox={legendViewbox} height={SQUARE_SIZE - SQUARE_BORDER_SIZE}>
                <g class="vch__legend__wrapper">
                  {#each rangeColor as color, index}
                    <rect rx={round} ry={round} style:fill={color} width={SQUARE_SIZE - SQUARE_BORDER_SIZE} height={SQUARE_SIZE - SQUARE_BORDER_SIZE} x={SQUARE_SIZE * index} />
                  {/each}
                </g>
              </svg>
            {/if}
            <div>
              {lo.more}
            </div>
          </div>
        </slot>
      </div>
    </slot>
  </div>
</div>

<style lang="postcss">
  :global(.vch__container) {
    .vch__legend {
      @apply flex items-center justify-between;
    }

    .vch__external-legend-wrapper {
      @apply my-2;
    }
  }
  :global(svg.vch__wrapper) {
    @apply w-full leading-[10px];
    .vch__months__labels__wrapper text.vch__month__label {
      @apply text-[8px];
    }

    .vch__days__labels__wrapper text.vch__day__label,
    .vch__legend__wrapper text {
      @apply text-[9px];
    }

    & text.vch__month__label,
    & text.vch__day__label,
    & .vch__legend__wrapper text {
      @apply fill-base-content;
    }

    & rect.vch__day__square:hover {
      /*TODO: match with tailwind style*/
      /* paint-order: stroke; */
      @apply stroke-mako-800 stroke-2;
    }

    & rect.vch__day__square:focus {
      /* outline: none; */
      @apply outline-none;
    }
  }
  :global(.dark-mode svg.vch__wrapper) {
    & text.vch__month__label,
    & text.vch__day__label,
    & .vch__legend__wrapper text {
      @apply fill-base-content;
    }
  }
</style>
