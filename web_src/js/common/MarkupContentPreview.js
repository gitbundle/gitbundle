import $ from 'jquery'
import { initMarkupContent } from '../markup/content.js'

const { csrfToken } = window.config

export function initCompMarkupContentPreviewTab(target) {
  const $commentTabs = $(target)
  if (!$commentTabs.length) {
    console.error(`not found target: ${target}`)
    return
  }

  $commentTabs.each((i, el) => {
    $(el)
      .find(`[jq-tabs] > [data-tab]`)
      .off('click')
      .on('click', function () {
        const $this = $(this)
        const text = $this
          .closest(target)
          .find(
            `[jq-contents] [data-tab=${$this.parent().data('write')}] textarea`
          )
          .val()
        if (!text) {
          console.info('No content for preview')
          return
        }
        $this.addClass('tab-active').siblings().removeClass('tab-active')
        $this
          .closest(target)
          .find(`[jq-contents] [data-tab=${$this.data('tab')}]`)
          .show()
          .siblings()
          .hide()
        const url = $this.data('url')
        if (!url) return
        $.post(
          url,
          {
            _csrf: csrfToken,
            mode: 'comment',
            context: $this.data('context'),
            text
          },
          (data) => {
            const $previewPanel = $this
              .closest(target)
              .find(
                `[jq-contents] [data-tab=${$this.parent().data('preview')}]`
              )
            $previewPanel.html(data)
            initMarkupContent()
          }
        )
      })
  })
}
