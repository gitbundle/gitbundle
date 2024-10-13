import $ from '../jquery.js'
import './are-you-sure.js'
import './popup.js'
import './transition.js'

$.expr[':'].icontains = $.expr.createPseudo(function (arg) {
  return function (elem) {
    return $(elem).text().toUpperCase().indexOf(arg.toUpperCase()) >= 0
  }
})

export default $