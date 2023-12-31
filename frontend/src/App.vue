<script lang="ts" setup>
import NavBar from './components/layout/NavigationBar.vue'
import HeaderBar from './components/layout/HeaderBar.vue'
import FloatingDrawer from './components/layout/FloatingDrawer.vue'
import { RouterView } from 'vue-router'
</script>

<template>
    <section class="app">
        <HeaderBar height='48px' />
        <NavBar class='nav'/>
        <div class='view'>
            <router-view/>
        </div>
        <FloatingDrawer :visible='true' />
    </section>
</template>

<style>
.app {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100%;
}
.nav {
    position: fixed;
    top: 0;
    left: 0;
    height: 100%;
    width: 15%;
    z-index: 100;
}

.view {
    display: flex;
    margin-left: 10%;
    height: 95%;
    width: 87%;
    border-radius: 5px;
    justify-content: center;
    align-items: center;
    overflow: hidden;
}
</style>
