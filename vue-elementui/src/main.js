import ElementUI from 'element-ui'

import 'element-ui/lib/theme-chalk/index.css'
import Vue from 'vue'
import App from './App'
// 新添加1
import router from './router'

Vue.use(ElementUI)
Vue.use(router)


Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  render:h=>h(App)
})
