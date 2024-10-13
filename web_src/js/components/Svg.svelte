<script>
  // const svgs = (name) => import(`../../../public/img/svg/${name}.svg`)

  // const parser = new DOMParser()
  // const serializer = new XMLSerializer()

  // export const svg = async (name, size = 16, className = '') => {
  //   const module = await svgs(name)

  //   const document = parser.parseFromString(module.default, 'image/svg+xml')
  //   const svgNode = document.firstChild
  //   svgNode.setAttribute('width', String(size))
  //   svgNode.setAttribute('height', String(size))
  //   if (className) svgNode.classList.add(...className.split(/\s+/).filter(Boolean))
  //   return serializer.serializeToString(svgNode)
  // }

  import { svg } from '../common/svg.js'

  export let name
  export let size
  export let className
</script>

{#await svg(name, size, className) then module}
  {@html module}
{/await}
