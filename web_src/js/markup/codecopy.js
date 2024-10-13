import { svg } from '../common/svg.js'

export async function renderCodeCopy() {
  const els = document.querySelectorAll('.markup .code-block code')
  if (!els.length) return

  const button = document.createElement('button')
  button.classList.add('code-copy', 'btn', 'btn-square', 'btn-sm')
  button.innerHTML = await svg('octicon-copy')

  for (const el of els) {
    const btn = button.cloneNode(true)
    btn.setAttribute('data-clipboard-text', el.textContent)
    el.after(btn)
  }
}
