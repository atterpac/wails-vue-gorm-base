import { createRouter, createWebHashHistory } from 'vue-router';
import HelloWorld from './components/HelloWorld.vue'
import LoginComponent from './components/auth/LoginComponent.vue'
import RegisterComponent from './components/auth/RegisterComponent.vue'
import PasswordReset from './components/auth/ResetPassword.vue'
import RunTime from './components/RunTime.vue'

const routes = [
    {path:"/", name:"Home" , component: HelloWorld },
    {path:"/login", name:"Login" , component: LoginComponent },
    {path:"/register", name:"Register" , component: RegisterComponent },
    {path:"/reset", name:"PasswordReset" , component: PasswordReset },
    {path:"/runtime", name:"RunTime" , component: RunTime }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes,
})


