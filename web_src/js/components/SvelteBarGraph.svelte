<script>
  import gsap from 'gsap'
  import { onMount } from 'svelte'
  export let title = '' // String
  export let points = () => [] // Array
  export let height = 100 // Number
  export let width = 300 // Number
  export let showYAxis = false // Boolean
  export let showXAxis = false // Boolean
  export let labelHeight = 12 // Number
  export let showTrendLine = false // Boolean
  export let trendLineColor = 'green' // String
  export let trendLineWidth = 2 // Number
  export let easeIn = true // Boolean
  export let showValues = false // Boolean
  export let maxYAxis = 0 // Number
  export let animationDuration = 0.5 // Number
  export let barColor = 'deepskyblue' // String
  export let textColor = 'black' // String
  export let textAltColor = 'black' // String
  export let textFont = '10px sans-serif' // String
  export let useCustomLabels = false // Boolean
  export let customLabels = () => [] // Array
  let dynamicPoints = []
  let staticPoints = []
  let extraTopHeightForYAxisLabel = 4
  let extraBottomHeightForYAxisLabel = 4
  let digitsUsedInYAxis = 0
  $: usingObjectsForDataPoints = points.every((x) => typeof x === 'object')
  $: dataPoints = usingObjectsForDataPoints ? points.map((item) => item.value) : points
  $: dataLabels = points.map((point, i) => {
    if (useCustomLabels) {
      return customLabels[i]
    }
    return usingObjectsForDataPoints ? point.label : i + 1
  })
  $: dataColors = points.map((item) => ({
    barColor: item && item.barColor ? item.barColor : barColor,
    textColor: item && item.textColor ? item.textColor : textColor,
    textAltColor: item && item.textAltColor ? item.textAltColor : textAltColor
  }))
  $: yAxisWidth = digitsUsedInYAxis * 5.8 + 5
  $: xAxisHeight = showYAxis ? labelHeight : labelHeight + extraBottomHeightForYAxisLabel + extraTopHeightForYAxisLabel
  $: fullSvgWidth = width
  $: fullSvgHeight = height
  $: innerChartWidth = showYAxis ? width - yAxisWidth : width
  $: innerChartHeight = (() => {
    let chartHeight = height
    if (showYAxis) {
      chartHeight -= extraTopHeightForYAxisLabel + extraBottomHeightForYAxisLabel
    }
    if (showXAxis) {
      chartHeight -= xAxisHeight
    }
    return chartHeight
  })()
  $: partitionWidth = innerChartWidth / dataPoints.length
  $: maxDomain = maxYAxis ? maxYAxis : Math.ceil(Math.max(...dataPoints))
  $: chartData = dynamicPoints.map((dynamicValue, index) => ({
    staticValue: staticPoints[index],
    index,
    label: dataLabels[index],
    width: partitionWidth - 2,
    midPoint: partitionWidth / 2,
    yLabel: innerChartHeight + 4,
    x: index * partitionWidth,
    xMidpoint: index * partitionWidth + partitionWidth / 2,
    yOffset: innerChartHeight - y(dynamicValue),
    height: y(dynamicValue),
    barColor: dataColors[index].barColor,
    textColor: dataColors[index].textColor,
    textAltColor: dataColors[index].textAltColor
  }))
  $: trendLine = (() => {
    const slopeValues = applySlope(dynamicPoints)
    return {
      x1: partitionWidth / 2,
      y1: roundTo(innerChartHeight - y(slopeValues[0]), 2),
      x2: innerChartWidth - partitionWidth / 2,
      y2: roundTo(innerChartHeight - y(slopeValues[slopeValues.length - 1]), 2)
    }
  })()

  // created:
  $: if (easeIn) {
    tween(dataPoints)
  } else {
    dynamicPoints = dataPoints
    staticPoints = dataPoints
  }

  function y(val) {
    return (val / maxDomain) * innerChartHeight
  }

  function roundTo(n, digits = 0) {
    let negative = false
    let number = n
    if (number < 0) {
      negative = true
      number *= -1
    }
    const multiplicator = 10 ** digits
    number = parseFloat((number * multiplicator).toFixed(11))
    number = (Math.round(number) / multiplicator).toFixed(2)
    if (negative) {
      number = (number * -1).toFixed(2)
    }
    return number
  }

  function tween(desiredDataArray) {
    const desiredData = {}
    const initialData = {}
    for (let i = 0; i < desiredDataArray.length; i += 1) {
      const key = i.toString()
      desiredData[key] = desiredDataArray[i]
      initialData[key] = dynamicPoints[i] || 0
    }
    const convertBackToArray = () => {
      const obj = Object.values(initialData)
      obj.pop()
      dynamicPoints = obj
    }
    gsap.to(initialData, {
      ...desiredData,
      onUpdate: convertBackToArray,
      duration: animationDuration
    })
    staticPoints = desiredDataArray
  }

  function getTicks() {
    for (let i = 6; i > 0; i -= 1) {
      if (maxDomain % i === 0) {
        const shouldForceDecimals = i < 3
        const numberOfTicks = shouldForceDecimals ? 3 : i
        digitsUsedInYAxis = maxDomain.toFixed(shouldForceDecimals ? 1 : 0).replace('.', '').length
        return [...new Array(numberOfTicks + 1)].map((item, key) => {
          const tickValue = (maxDomain / numberOfTicks) * (numberOfTicks - key)
          const yCoord = (innerChartHeight / numberOfTicks) * key
          return {
            key,
            text: shouldForceDecimals ? tickValue.toFixed(1) : tickValue,
            yText: yCoord < 10 ? 10 : yCoord + 4,
            x1: yAxisWidth - 4,
            y1: yCoord,
            x2: yAxisWidth - 1,
            y2: yCoord
          }
        })
      }
    }
    return []
  }

  function applySlope(values) {
    let xAvg = 0
    let yAvg = 0
    for (let x = 0; x < values.length; x += 1) {
      xAvg += x
      yAvg += values[x]
    }
    xAvg /= values.length
    yAvg /= values.length
    let v1 = 0
    let v2 = 0
    for (let x = 0; x < values.length; x += 1) {
      v1 += (x - xAvg) * (values[x] - yAvg)
      v2 += (x - xAvg) ** 2
    }
    const a = v1 / v2
    const b = yAvg - a * xAvg
    const result = []
    for (let index = 0; index < values.length; index += 1) {
      result.push(a * index + b)
    }
    return result
  }
</script>

<svg width={fullSvgWidth} height={fullSvgHeight} aria-labelledby="title" role="img">
  {#if title}
    <title id="title">
      {title}
    </title>
  {/if}
  <g transform={`translate(0,${showYAxis ? extraTopHeightForYAxisLabel : 0})`}>
    <g transform={`translate(${showYAxis ? yAxisWidth : 0},0)`} width={innerChartWidth} height={innerChartHeight}>
      {#each chartData as bar}
        <g key={bar.index} transform={`translate(${bar.x},0)`}>
          <title>
            <slot name="title" {bar}>
              <tspan>
                {bar.staticValue}
              </tspan>
            </slot>
          </title>
          <rect width={bar.width} height={bar.height} x={2} y={bar.yOffset} style:fill={bar.barColor} />
          {#if showValues}
            <text x={bar.midPoint} y={bar.yOffset} dy={`${bar.height < 22 ? '-5px' : '15px'}`} text-anchor="middle" style:fill={bar.height < 22 ? bar.textColor : bar.textAltColor} style:font={textFont}>
              {bar.staticValue}
            </text>
          {/if}
          {#if showXAxis}
            <g>
              <slot name="label" {bar} text-style:fill={textColor} text-style:font={textFont}>
                <text x={bar.midPoint} y={`${bar.yLabel + 10}px`} text-anchor="middle" style:fill={textColor} style:font={textFont}>
                  {bar.label}
                </text>
              </slot>
              <line x1={bar.midPoint} x2={bar.midPoint} y1={innerChartHeight + 3} y2={innerChartHeight} stroke="#555555" stroke-width="1" />
            </g>
          {/if}
        </g>
      {/each}
      {#if showTrendLine}
        <line x1={trendLine.x1} y1={trendLine.y1} x2={trendLine.x2} y2={trendLine.y2} stroke-width={trendLineWidth} stroke={trendLineColor} />
      {/if}
    </g>
    {#if showXAxis}
      <g>
        <line x1={showYAxis ? yAxisWidth - 1 : 2} x2={innerChartWidth + yAxisWidth} y1={innerChartHeight} y2={innerChartHeight} stroke="#555555" stroke-width="1" />
      </g>
    {/if}
    {#if showYAxis}
      <g>
        <line x1={yAxisWidth - 1} x2={yAxisWidth - 1} y1={innerChartHeight} y2="0" stroke="#555555" stroke-width="1" />
        {#each getTicks() as tick}
          <g key={tick.key}>
            <line x1={tick.x1} y1={tick.y1} x2={tick.x2} y2={tick.y2} stroke="#555555" stroke-width="1" />
            <text x="0" y={tick.yText} alignment-baseline="central" style:fill={textColor} style:font={textFont}>
              {tick.text}
            </text>
          </g>
        {/each}
      </g>
    {/if}
  </g>
</svg>
