import Vue from 'vue'
import Router from 'vue-router'
import login from '../Login'
import Main from '../views/Main.vue'
import UserSystem from '../System'
import UserProfile from '../views/user/Profile'
import UserList from '../views/user/List'


Vue.use(Router);


export default new Router({
  routes: [{
    path: '/main',
    component: Main
  },
    {
      path: '/login',
      component: login
    },
    {
      path: '/system',
      component: UserSystem,
      children:[
        {path:'/profile',component:UserProfile},
        {path:'/list',component:UserList}
      ]
    }
  ]
});
