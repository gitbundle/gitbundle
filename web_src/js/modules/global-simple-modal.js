import GlobalSimpleModal from '../components/GlobalSimpleModal.svelte'

function _init() {
  const els = document.querySelectorAll('[svelte-simple-modal]')
  els.forEach((el) => {
    el.addEventListener('click', () => {
      const data = {
        dyFieldType: {
          _csrf: 'hidden'
        },
        dyFieldLabel: {},
        dyFieldRequired: {},

        csrfToken: window.config.csrfToken,
        asyncForm: true,
        checked: false,
        ...window.config.pageData.modalDeletion
      }
      // const dyFieldType = {
      //     _csrf: 'hidden'
      //   },
      //   dyFieldLabel = {},
      //   dyFieldRequired = {}

      let formData = new FormData()
      formData.append('_csrf', window.config.csrfToken)

      Array.from(el.attributes).forEach(({ name, value }) => {
        // console.log(name, value)
        const match = name.match(/data-form-key\[([0-9-a-zA-Z]+)\]/)
        if (match && match.length == 2) {
          // console.log(match, match[1])
          const fieldValue = el.getAttribute(`data-form-val[${match[1]}]`)
          if (!fieldValue) {
            console.log(
              `Unmatched value for '${match}' '${value}' '${fieldValue}'`
            )
          }
          formData.append(value, fieldValue ?? '')

          const fieldType = el.getAttribute(`data-form-type[${match[1]}]`)
          data.dyFieldType[value] = fieldType ?? 'hidden'

          const fieldLabel = el.getAttribute(`data-form-label[${match[1]}]`)
          data.dyFieldLabel[value] = fieldLabel ?? undefined

          const fieldRequired = el.getAttribute(
            `data-form-required[${match[1]}]`
          )
          data.dyFieldRequired[value] = fieldRequired ?? undefined
        }
      })
      data.method = el.getAttribute('data-method') ?? 'post'
      data.href = el.getAttribute('data-href')
      data.name = el.getAttribute('data-name')
      data.desc = el.getAttribute('data-desc') ?? undefined
      data.descClass = el.getAttribute('data-desc-class') ?? undefined
      data.title = el.getAttribute('data-title')
      data.titleSvgName = el.getAttribute('data-title-svg-name') ?? undefined
      data.approveText = el.getAttribute('data-action-approve-text')
      data.approveColor = el.getAttribute('data-action-approve-color')
      data.negativeText = el.getAttribute('data-action-negative-text')
      data.confirmForm = el.getAttribute('data-confirm-form') ?? undefined

      const asyncForm = el.getAttribute('data-async')
      if (asyncForm !== undefined) {
        data.async = asyncForm
      }

      if (el.tagName !== 'label') {
        data.checked = true
      }

      new GlobalSimpleModal({
        target: document.body,
        props: { data, formData }
      })
    })
  })
}

_init()
