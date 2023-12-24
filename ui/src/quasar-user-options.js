
import './styles/quasar.sass'
import '@quasar/extras/material-icons/material-icons.css'
import '@quasar/extras/fontawesome-v6/fontawesome-v6.css'

import { Notify, Dialog } from 'quasar'

// To be used on app.use(Quasar, { ... })
export default {
  plugins: {
    Notify,
    Dialog
  },
  extras: [
    'material-icons',
    'fontawesome-v6',
    'roboto-font'
  ],
  config: {
    notify: { /* look at QuasarConfOptions from the API card */ }
  }
}