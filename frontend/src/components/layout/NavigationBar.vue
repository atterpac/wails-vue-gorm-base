<template>
    <section class="container" :style="{ width: isCollapse ? '56px' : '140px' }">
        <CubeTransparent @click='toggleCollapse' class="logo" hero_stroke="var(--primary-color)" hero_width='48px' hero_height='48px'/>
        <div class='nav-bar'>
            <div class='nav-item' v-for="item in navigationItems" :key="item.path">
                <button 
                    class='nav-button'
                    :key=item.path
                    @click='onClick(item.path)'>
                    <div class='icon'>
                        <Component :is='item.icon' hero_style="filled" hero_stroke="var(--primary-color)" hero_fill="var(--primary-color)" hero_width='24px' hero_height='24px' />
                    </div>
                    <div class='name' v-if='!isCollapse'>
                        {{item.name}}
                    </div>
                </button>
            </div>
        </div>
    </section>
</template>

<script lang="ts" setup>
import type { RouteLocationRaw } from 'vue-router';
import { router } from '../../router';
import CubeTransparent from '../../assets/hero_icons/CubeTransparent.vue'
import HomeIcon from '../../assets/hero_icons/HomeIcon.vue'
import UserIcon from '../../assets/hero_icons/UserIcon.vue'
import ColorSwatch from '../../assets/hero_icons/ColorSwatch.vue'
import type { Component } from 'vue';
import { ref } from 'vue';

type NavigationItem = {
    name: string;
    path: string;
    icon?: Component
}

const navigationItems: NavigationItem[] = [
    {
        name: 'Home',
        path: '/',
        icon: HomeIcon
    },
    {
        name: 'Auth',
        path: '/login',
        icon: UserIcon
    },
    {
        name: 'Colors',
        path: '/colors',
        icon: ColorSwatch
    },
    {
        name: 'Events',
        path: '/events',
        icon: ColorSwatch
    },
]

const isCollapse = ref(false);

const toggleCollapse = () => {
    isCollapse.value = !isCollapse.value;
    // Get .nav-button and change justify-content to center
    var root = document.querySelector(':root') as HTMLElement;
    root.style.setProperty('--nav-justify-content', isCollapse.value ? 'center' : 'flex-start');
}

const onClick = (path: RouteLocationRaw) => {
    router.push(path);
}
</script>

<style scoped>
:root {
    --nav-justify-content: flex-start;
}

.container {
    display: flex;
    flex-direction: column;
    background-color: var(--background-color-light);

}
.nav-bar {
    display: flex;
    flex-direction: column;
    align-items: center;
}

.nav-item {
    height: 36px;
    width: 100%;
    display: flex;
    flex-direction: row;
    align-items: center;
}

.nav-button {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: var(--nav-justify-content);
    height: 100%;
}

button {
    width: 100%;
    height: 36px;
    border: none;
    background-color: transparent;
    color: white;
    font-size: 1rem;
    text-align: center;
    cursor: pointer;
    border-radius: 5px;
}

.icon {
    margin-right: 10px;
    margin-left: 10px;
}

.nav-item:hover {
    background-color: var(--gray-2);
}

.logo {
    margin-top: 24px;
    margin-bottom: 10px;
    cursor: pointer;
}
</style>

