import { createDropzone } from '../common/dropzone.js'
import { svg } from '../common/svg.js'
import $ from '../features/jquery.js'

export function initGlobalDropzone() {
  const { csrfToken } = window.config

  // Dropzone
  for (const el of document.querySelectorAll('.dropzone')) {
    const $dropzone = $(el)
    const _promise = createDropzone(el, {
      url: $dropzone.data('upload-url'),
      headers: { 'X-Csrf-Token': csrfToken },
      maxFiles: $dropzone.data('max-file'),
      maxFilesize: $dropzone.data('max-size'),
      acceptedFiles: ['*/*', ''].includes($dropzone.data('accepts'))
        ? null
        : $dropzone.data('accepts'),
      addRemoveLinks: true,
      dictDefaultMessage: $dropzone.data('default-message'),
      dictInvalidFileType: $dropzone.data('invalid-input-type'),
      dictFileTooBig: $dropzone.data('file-too-big'),
      dictRemoveFile: $dropzone.data('remove-file'),
      timeout: 0,
      thumbnailMethod: 'contain',
      thumbnailWidth: 480,
      thumbnailHeight: 480,
      init() {
        this.on('success', async (file, data) => {
          file.uuid = data.uuid
          const input = $(
            `<input id="${data.uuid}" name="files" type="hidden">`
          ).val(data.uuid)
          $dropzone.find('.files').append(input)
          const copyIcon = await svg('octicon-copy', 14, 'copy link')
          // Create a "Copy Link" element, to conveniently copy the image
          // or file link as Markdown to the clipboard
          const copyLinkElement = document.createElement('div')
          copyLinkElement.className =
            'btn btn-xs btn-outline items-center gap-x-1'
          // The a element has a hardcoded cursor: pointer because the default is overridden by .dropzone
          copyLinkElement.innerHTML = `${copyIcon} Copy link`
          copyLinkElement.addEventListener('click', (e) => {
            e.preventDefault()
            let fileMarkdown = `[${file.name}](/attachments/${file.uuid})`
            if (file.type.startsWith('image/')) {
              fileMarkdown = `!${fileMarkdown}`
            }
            navigator.clipboard.writeText(fileMarkdown)
          })
          file.previewTemplate.appendChild(copyLinkElement)
        })
        this.on('removedfile', (file) => {
          $(`#${file.uuid}`).remove()
          if (!$dropzone.find('.dz-preview').length) {
            $dropzone.find('.dz-message').show()
          }
          if ($dropzone.data('remove-url')) {
            $.post($dropzone.data('remove-url'), {
              file: file.uuid,
              _csrf: csrfToken
            })
          }
        })
      }
    })
  }
}

initGlobalDropzone()
