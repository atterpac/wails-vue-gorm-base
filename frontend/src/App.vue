<script lang="ts" setup>
import NavBar from './components/layout/NavigationBar.vue'
import HeaderBar from './components/layout/HeaderBar.vue'
import FloatingDrawer from './components/layout/FloatingDrawer.vue'
import { RouterView } from 'vue-router'
</script>

<template>
<section class="app">
    <HeaderBar height='48px' />
    <div class="wrapper">
        <NavBar class='nav'/>
        <div class='view'>
            <router-view />
        </div>
        <FloatingDrawer :visible='true' />
    </div>
</section>
</template>

<style>
.app {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    position: absolute;
    top: 0;
    left: 0;
}

.wrapper {
    flex: 1;
    display: flex;
    justify-content: flex-start;
    overflow: hidden;
}

.nav {
    justify-self: flex-start;
}

.view {
    background-color: var(--background-color);
    border-radius: 8px;
    display: flex;
    flex: 1;
    justify-self: center;
    justify-content: center;
    align-items: center;
    min-width: 300px;
    overflow: hidden;
    margin-left: 20px;
    margin-right: 20px;
    margin-bottom: 20px;
    margin-top: 20px;
    /* box-shadow: 0px 0px 20px 0px hsla(var(--hue), 70%, 20%, 0.75); */
}
</style>
