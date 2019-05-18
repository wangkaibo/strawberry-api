import Vue from 'vue'
import ElementUI from 'element-ui'
import App from './App.vue'
import router from './router'
import store from './store'
import '../node_modules/element-ui/lib/theme-chalk/index.css'
import './assets/index.css'

Vue.config.productionTip = false

Vue.use(ElementUI);

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')
